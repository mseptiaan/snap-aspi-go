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

	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// OptimizedSnapClient implements the HTTPClient interface with optimizations
type OptimizedSnapClient struct {
	baseURL         string
	httpClient      *http.Client
	timeout         time.Duration
	maxRetries      int
	backoffDuration time.Duration
	logger          logging.Logger
	config          *config.Config

	// Request body pool for reuse
	bodyPool sync.Pool
}

// NewOptimizedSnapClient creates a new optimized ASPI HTTP client
func NewOptimizedSnapClient(cfg *config.Config, logger logging.Logger) contracts.HTTPClient {
	httpClient := CreateOptimizedHTTPClient(
		cfg.ASPI.Timeouts.ConnectTimeout,
		cfg.ASPI.Timeouts.RequestTimeout,
	)

	client := &OptimizedSnapClient{
		baseURL:         cfg.ASPI.BaseURL,
		httpClient:      httpClient,
		timeout:         cfg.ASPI.Timeouts.RequestTimeout,
		maxRetries:      cfg.ASPI.Retry.MaxAttempts,
		backoffDuration: cfg.ASPI.Retry.BackoffDuration,
		logger:          logger,
		config:          cfg,
	}

	// Initialize body pool
	client.bodyPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	return client
}

// Get performs a GET request
func (c *OptimizedSnapClient) Get(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	return c.executeRequest(ctx, "GET", url, nil, headers)
}

// Post performs a POST request with JSON body
func (c *OptimizedSnapClient) Post(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "POST", url, body, headers)
}

// Put performs a PUT request with JSON body
func (c *OptimizedSnapClient) Put(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "PUT", url, body, headers)
}

// Delete performs a DELETE request
func (c *OptimizedSnapClient) Delete(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	return c.executeRequest(ctx, "DELETE", url, nil, headers)
}

// WithTimeout creates a new client with different timeout
func (c *OptimizedSnapClient) WithTimeout(timeout time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.timeout = timeout
	newClient.httpClient = CreateOptimizedHTTPClient(timeout, timeout)
	return &newClient
}

// WithRetry creates a new client with different retry configuration
func (c *OptimizedSnapClient) WithRetry(maxAttempts int, backoffDuration time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.maxRetries = maxAttempts
	newClient.backoffDuration = backoffDuration
	return &newClient
}

// WithBaseURL creates a new client with different base URL
func (c *OptimizedSnapClient) WithBaseURL(baseURL string) contracts.HTTPClient {
	newClient := *c
	newClient.baseURL = baseURL
	return &newClient
}

// executeRequest executes an HTTP request with optimizations
func (c *OptimizedSnapClient) executeRequest(
	ctx context.Context,
	method, url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	fullURL := c.buildURL(url)

	// Prepare request body using pool
	var bodyReader io.Reader
	var bodyBytes []byte

	if body != nil {
		// Get buffer from pool
		buf := c.bodyPool.Get().(*bytes.Buffer)
		defer func() {
			buf.Reset()
			c.bodyPool.Put(buf)
		}()

		// Encode JSON directly to buffer
		encoder := json.NewEncoder(buf)
		if err := encoder.Encode(body); err != nil {
			return nil, errors.NewValidationError(fmt.Sprintf("failed to marshal request body: %v", err))
		}

		bodyBytes = buf.Bytes()
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var lastErr error

	// Retry logic with exponential backoff
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff
			backoffTime := c.backoffDuration * time.Duration(1<<uint(attempt-1))
			select {
			case <-ctx.Done():
				return nil, errors.NewNetworkError(ctx.Err(), "request context canceled during retry")
			case <-time.After(backoffTime):
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
		c.setHeaders(req, headers)

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
				}).WithError(err).Warn("HTTP request failed, retrying")
			}
			continue
		}

		// Process response
		response, err := c.processResponse(resp)
		if err != nil {
			resp.Body.Close()
			lastErr = err
			continue
		}

		return response, nil
	}

	return nil, lastErr
}

// buildURL constructs the full URL
func (c *OptimizedSnapClient) buildURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}

	baseURL := strings.TrimSuffix(c.baseURL, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")

	return fmt.Sprintf("%s/%s", baseURL, endpoint)
}

// setHeaders sets default and custom headers
func (c *OptimizedSnapClient) setHeaders(req *http.Request, headers map[string]string) {
	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("snap-aspi-go/%s", c.config.App.Version))
	req.Header.Set("Connection", "keep-alive")

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// processResponse processes the HTTP response with optimizations
func (c *OptimizedSnapClient) processResponse(resp *http.Response) (*contracts.Response, error) {
	defer resp.Body.Close()

	// Pre-allocate buffer based on content length
	var bodyBytes []byte
	if resp.ContentLength > 0 {
		bodyBytes = make([]byte, 0, resp.ContentLength)
	}

	// Read response body
	buf := bytes.NewBuffer(bodyBytes)
	_, err := io.Copy(buf, resp.Body)
	if err != nil {
		return nil, errors.NewNetworkError(err, "failed to read response body")
	}

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
		Body:       buf.Bytes(),
		Success:    resp.StatusCode >= 200 && resp.StatusCode < 300,
	}

	// Check for API errors
	if !response.Success {
		return response, c.handleHTTPError(response)
	}

	return response, nil
}

// handleHTTPError handles HTTP error responses
func (c *OptimizedSnapClient) handleHTTPError(response *contracts.Response) error {
	statusCode := response.StatusCode

	switch {
	case statusCode == 401:
		return errors.NewAuthenticationError("authentication failed")
	case statusCode == 403:
		return errors.NewAuthorizationError("access forbidden")
	case statusCode >= 400 && statusCode < 500:
		return errors.NewValidationError(fmt.Sprintf("client error: %d", statusCode))
	case statusCode >= 500:
		return errors.NewExternalError(fmt.Errorf("server error: %d", statusCode), "ASPI server error")
	default:
		return errors.NewNetworkError(fmt.Errorf("unexpected status code: %d", statusCode), "unexpected HTTP error")
	}
}