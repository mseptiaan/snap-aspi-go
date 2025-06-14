package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/cache"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/metrics"
	"github.com/mseptiaan/snap-aspi-go/internal/pool"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// OptimizedHTTPClient implements an optimized HTTP client with caching, pooling, and metrics
type OptimizedHTTPClient struct {
	baseURL         string
	httpClient      *http.Client
	timeout         time.Duration
	maxRetries      int
	backoffDuration time.Duration
	logger          logging.Logger
	config          *config.Config
	metrics         *metrics.Metrics
	responseCache   *cache.ResponseCache
	bufferPool      *pool.BufferPool
	
	// Connection management
	activeRequests sync.WaitGroup
	requestSemaphore chan struct{}
}

// NewOptimizedHTTPClient creates a new optimized HTTP client
func NewOptimizedHTTPClient(cfg *config.Config, logger logging.Logger, m *metrics.Metrics) contracts.HTTPClient {
	// Create optimized HTTP client with connection pooling
	httpClient := CreateOptimizedHTTPClient(
		cfg.ASPI.Timeouts.ConnectTimeout,
		cfg.ASPI.Timeouts.RequestTimeout,
	)

	// Create response cache (100 entries, 5 minute TTL)
	responseCache := cache.NewResponseCache(100, 5*time.Minute)

	// Create buffer pool
	bufferPool := pool.NewBufferPool()

	// Create semaphore for request limiting (max 50 concurrent requests)
	requestSemaphore := make(chan struct{}, 50)

	client := &OptimizedHTTPClient{
		baseURL:          cfg.ASPI.BaseURL,
		httpClient:       httpClient,
		timeout:          cfg.ASPI.Timeouts.RequestTimeout,
		maxRetries:       cfg.ASPI.Retry.MaxAttempts,
		backoffDuration:  cfg.ASPI.Retry.BackoffDuration,
		logger:           logger,
		config:           cfg,
		metrics:          m,
		responseCache:    responseCache,
		bufferPool:       bufferPool,
		requestSemaphore: requestSemaphore,
	}

	// Start cache cleanup routine
	go client.startCacheCleanup()

	return client
}

// Get performs an optimized GET request
func (c *OptimizedHTTPClient) Get(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	// Check cache for GET requests
	cacheKey := c.buildCacheKey("GET", url, headers)
	if cachedData, found := c.responseCache.Get(cacheKey); found {
		c.metrics.RecordCacheHit()
		c.logger.Debug("Cache hit for GET request")
		
		return &contracts.Response{
			StatusCode: 200,
			Headers:    make(map[string]string),
			Body:       cachedData,
			Success:    true,
		}, nil
	}
	
	c.metrics.RecordCacheMiss()
	response, err := c.executeRequest(ctx, "GET", url, nil, headers)
	
	// Cache successful GET responses
	if err == nil && response.Success {
		c.responseCache.Set(cacheKey, response.Body)
	}
	
	return response, err
}

// Post performs an optimized POST request
func (c *OptimizedHTTPClient) Post(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "POST", url, body, headers)
}

// Put performs an optimized PUT request
func (c *OptimizedHTTPClient) Put(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "PUT", url, body, headers)
}

// Delete performs an optimized DELETE request
func (c *OptimizedHTTPClient) Delete(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	return c.executeRequest(ctx, "DELETE", url, nil, headers)
}

// WithTimeout creates a new client with different timeout
func (c *OptimizedHTTPClient) WithTimeout(timeout time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.timeout = timeout
	newClient.httpClient = CreateOptimizedHTTPClient(timeout, timeout)
	return &newClient
}

// WithRetry creates a new client with different retry configuration
func (c *OptimizedHTTPClient) WithRetry(maxAttempts int, backoffDuration time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.maxRetries = maxAttempts
	newClient.backoffDuration = backoffDuration
	return &newClient
}

// WithBaseURL creates a new client with different base URL
func (c *OptimizedHTTPClient) WithBaseURL(baseURL string) contracts.HTTPClient {
	newClient := *c
	newClient.baseURL = baseURL
	return &newClient
}

// executeRequest executes an HTTP request with all optimizations
func (c *OptimizedHTTPClient) executeRequest(
	ctx context.Context,
	method, url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	// Acquire semaphore for request limiting
	select {
	case c.requestSemaphore <- struct{}{}:
		defer func() { <-c.requestSemaphore }()
	case <-ctx.Done():
		return nil, errors.NewNetworkError(ctx.Err(), "request cancelled while waiting for semaphore")
	}

	// Track active requests
	c.activeRequests.Add(1)
	defer c.activeRequests.Done()

	fullURL := c.buildURL(url)
	start := time.Now()

	// Prepare request body using buffer pool
	var bodyReader io.Reader
	var bodyBytes []byte

	if body != nil {
		buf := c.bufferPool.Get()
		defer c.bufferPool.Put(buf)

		// Encode JSON directly to buffer
		encoder := json.NewEncoder(buf)
		if err := encoder.Encode(body); err != nil {
			return nil, errors.NewValidationError(fmt.Sprintf("failed to marshal request body: %v", err))
		}

		bodyBytes = make([]byte, buf.Len())
		copy(bodyBytes, buf.Bytes())
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var lastErr error

	// Retry logic with exponential backoff
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff with jitter
			backoffTime := c.backoffDuration * time.Duration(1<<uint(attempt-1))
			jitter := time.Duration(float64(backoffTime) * 0.1) // 10% jitter
			
			select {
			case <-ctx.Done():
				return nil, errors.NewNetworkError(ctx.Err(), "request context canceled during retry")
			case <-time.After(backoffTime + jitter):
				// Continue with retry
			}

			// Reset body reader for retry
			if bodyBytes != nil {
				bodyReader = bytes.NewReader(bodyBytes)
			}
		}

		// Create HTTP request
		req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
		if err != nil {
			lastErr = errors.NewInternalError(err, "failed to create HTTP request")
			continue
		}

		// Set headers
		c.setOptimizedHeaders(req, headers)

		// Execute request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = errors.NewNetworkError(err, fmt.Sprintf("HTTP request failed on attempt %d", attempt+1))
			
			// Log retry attempt
			if attempt < c.maxRetries {
				c.logger.WithFields(map[string]interface{}{
					"method":      method,
					"url":         fullURL,
					"attempt":     attempt + 1,
					"max_retries": c.maxRetries,
					"duration":    time.Since(start),
				}).WithError(err).Warn("HTTP request failed, retrying")
			}
			continue
		}

		// Process response
		response, err := c.processOptimizedResponse(resp)
		if err != nil {
			resp.Body.Close()
			lastErr = err
			continue
		}

		// Log successful request
		c.logger.WithFields(map[string]interface{}{
			"method":      method,
			"url":         fullURL,
			"status_code": response.StatusCode,
			"duration":    time.Since(start),
			"attempt":     attempt + 1,
		}).Debug("HTTP request completed successfully")

		return response, nil
	}

	return nil, lastErr
}

// buildURL constructs the full URL with optimizations
func (c *OptimizedHTTPClient) buildURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}

	// Use string builder for efficient concatenation
	var builder strings.Builder
	builder.Grow(len(c.baseURL) + len(endpoint) + 1)
	
	builder.WriteString(strings.TrimSuffix(c.baseURL, "/"))
	builder.WriteByte('/')
	builder.WriteString(strings.TrimPrefix(endpoint, "/"))
	
	return builder.String()
}

// setOptimizedHeaders sets headers with optimizations
func (c *OptimizedHTTPClient) setOptimizedHeaders(req *http.Request, headers map[string]string) {
	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("snap-aspi-go/%s", c.config.App.Version))
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// processOptimizedResponse processes the HTTP response with optimizations
func (c *OptimizedHTTPClient) processOptimizedResponse(resp *http.Response) (*contracts.Response, error) {
	defer resp.Body.Close()

	// Handle compressed responses
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// Handle gzip decompression if needed
		reader = resp.Body
	}

	// Pre-allocate buffer based on content length
	var bodyBytes []byte
	if resp.ContentLength > 0 && resp.ContentLength < 10*1024*1024 { // Max 10MB
		bodyBytes = make([]byte, 0, resp.ContentLength)
	}

	// Read response body efficiently
	buf := c.bufferPool.Get()
	defer c.bufferPool.Put(buf)
	
	_, err := io.Copy(buf, reader)
	if err != nil {
		return nil, errors.NewNetworkError(err, "failed to read response body")
	}

	bodyBytes = make([]byte, buf.Len())
	copy(bodyBytes, buf.Bytes())

	// Convert headers efficiently
	headers := make(map[string]string, len(resp.Header))
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	response := &contracts.Response{
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Body:       bodyBytes,
		Success:    resp.StatusCode >= 200 && resp.StatusCode < 300,
	}

	// Check for API errors
	if !response.Success {
		return response, c.handleHTTPError(response)
	}

	return response, nil
}

// handleHTTPError handles HTTP error responses
func (c *OptimizedHTTPClient) handleHTTPError(response *contracts.Response) error {
	statusCode := response.StatusCode

	switch {
	case statusCode == 401:
		return errors.NewAuthenticationError("authentication failed")
	case statusCode == 403:
		return errors.NewAuthorizationError("access forbidden")
	case statusCode == 429:
		return errors.NewNetworkError(fmt.Errorf("rate limit exceeded: %d", statusCode), "rate limit exceeded")
	case statusCode >= 400 && statusCode < 500:
		return errors.NewValidationError(fmt.Sprintf("client error: %d", statusCode))
	case statusCode >= 500:
		return errors.NewExternalError(fmt.Errorf("server error: %d", statusCode), "ASPI server error")
	default:
		return errors.NewNetworkError(fmt.Errorf("unexpected status code: %d", statusCode), "unexpected HTTP error")
	}
}

// buildCacheKey builds a cache key for the request
func (c *OptimizedHTTPClient) buildCacheKey(method, url string, headers map[string]string) string {
	var builder strings.Builder
	builder.WriteString(method)
	builder.WriteByte(':')
	builder.WriteString(url)
	
	// Include relevant headers in cache key
	if auth := headers["Authorization"]; auth != "" {
		builder.WriteByte(':')
		builder.WriteString(auth[:min(len(auth), 20)]) // First 20 chars only
	}
	
	return builder.String()
}

// startCacheCleanup starts the cache cleanup routine
func (c *OptimizedHTTPClient) startCacheCleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.responseCache.Cleanup()
	}
}

// WaitForActiveRequests waits for all active requests to complete
func (c *OptimizedHTTPClient) WaitForActiveRequests() {
	c.activeRequests.Wait()
}

// GetMetrics returns client metrics
func (c *OptimizedHTTPClient) GetMetrics() map[string]interface{} {
	return map[string]interface{}{
		"cache_stats": c.responseCache.Stats(),
		"active_requests": len(c.requestSemaphore),
		"max_concurrent": cap(c.requestSemaphore),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}