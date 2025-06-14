package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// SnapClient implements the HTTPClient interface for ASPI API calls
type SnapClient struct {
	baseURL         string
	httpClient      *http.Client
	timeout         time.Duration
	maxRetries      int
	backoffDuration time.Duration
	logger          logging.Logger
	config          *config.Config
}

// NewSnapClient creates a new ASPI HTTP client
func NewSnapClient(cfg *config.Config, logger logging.Logger) contracts.HTTPClient {
	httpClient := &http.Client{
		Timeout: cfg.ASPI.Timeouts.ConnectTimeout,
	}

	return &SnapClient{
		baseURL:         cfg.ASPI.BaseURL,
		httpClient:      httpClient,
		timeout:         cfg.ASPI.Timeouts.RequestTimeout,
		maxRetries:      cfg.ASPI.Retry.MaxAttempts,
		backoffDuration: cfg.ASPI.Retry.BackoffDuration,
		logger:          logger,
		config:          cfg,
	}
}

// Get performs a GET request
func (c *SnapClient) Get(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	return c.executeRequest(ctx, "GET", url, nil, headers)
}

// Post performs a POST request with JSON body
func (c *SnapClient) Post(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "POST", url, body, headers)
}

// Put performs a PUT request with JSON body
func (c *SnapClient) Put(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	return c.executeRequest(ctx, "PUT", url, body, headers)
}

// Delete performs a DELETE request
func (c *SnapClient) Delete(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	return c.executeRequest(ctx, "DELETE", url, nil, headers)
}

// WithTimeout creates a new client with different timeout
func (c *SnapClient) WithTimeout(timeout time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.timeout = timeout
	newClient.httpClient = &http.Client{
		Timeout: timeout,
	}
	return &newClient
}

// WithRetry creates a new client with different retry configuration
func (c *SnapClient) WithRetry(maxAttempts int, backoffDuration time.Duration) contracts.HTTPClient {
	newClient := *c
	newClient.maxRetries = maxAttempts
	newClient.backoffDuration = backoffDuration
	return &newClient
}

// WithBaseURL creates a new client with different base URL
func (c *SnapClient) WithBaseURL(baseURL string) contracts.HTTPClient {
	newClient := *c
	newClient.baseURL = baseURL
	return &newClient
}

// executeRequest executes an HTTP request with retries and proper error handling
func (c *SnapClient) executeRequest(
	ctx context.Context,
	method, url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	fullURL := c.buildURL(url)

	// Prepare request body
	var bodyReader io.Reader
	var bodyBytes []byte

	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, errors.NewValidationError(fmt.Sprintf("failed to marshal request body: %v", err))
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			// Wait before retry
			select {
			case <-ctx.Done():
				return nil, errors.NewNetworkError(ctx.Err(), "request context canceled during retry")
			case <-time.After(c.backoffDuration * time.Duration(attempt)):
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

		// Log request
		c.logRequest(req, bodyBytes)

		// Execute request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = errors.NewNetworkError(err, fmt.Sprintf("HTTP request failed on attempt %d", attempt+1))
			c.logger.WithFields(map[string]interface{}{
				"method":      method,
				"url":         fullURL,
				"attempt":     attempt + 1,
				"max_retries": c.maxRetries,
			}).WithError(err).Warn("HTTP request failed, retrying")
			continue
		}

		// Read response
		response, err := c.processResponse(resp)
		if err != nil {
			resp.Body.Close()
			lastErr = err
			continue
		}

		// Log response
		c.logResponse(response)

		return response, nil
	}

	return nil, lastErr
}

// buildURL constructs the full URL
func (c *SnapClient) buildURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}

	baseURL := strings.TrimSuffix(c.baseURL, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")

	return fmt.Sprintf("%s/%s", baseURL, endpoint)
}

// setHeaders sets default and custom headers
func (c *SnapClient) setHeaders(req *http.Request, headers map[string]string) {
	// Set default headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("github.com/mseptiaan/snap-aspi-go/%s", c.config.App.Version))

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// processResponse processes the HTTP response
func (c *SnapClient) processResponse(resp *http.Response) (*contracts.Response, error) {
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewNetworkError(err, "failed to read response body")
	}

	// Convert headers
	headers := make(map[string]string)
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
func (c *SnapClient) handleHTTPError(response *contracts.Response) error {
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

// logRequest logs the outgoing request
func (c *SnapClient) logRequest(req *http.Request, body []byte) {
	c.logger.WithFields(map[string]interface{}{
		"method":       req.Method,
		"url":          req.URL.String(),
		"headers":      req.Header,
		"body_size":    len(body),
		"content_type": req.Header.Get("Content-Type"),
	}).Info("Making HTTP request")
}

// logResponse logs the incoming response
func (c *SnapClient) logResponse(response *contracts.Response) {
	logLevel := "info"
	if !response.Success {
		logLevel = "warn"
	}

	fields := map[string]interface{}{
		"status_code": response.StatusCode,
		"body_size":   len(response.Body),
		"success":     response.Success,
	}

	if logLevel == "info" {
		c.logger.WithFields(fields).Info("HTTP response received")
	} else {
		c.logger.WithFields(fields).Warn("HTTP response with error status")
	}
}
