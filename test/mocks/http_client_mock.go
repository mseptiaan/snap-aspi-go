package mocks

import (
	"context"
	"time"

	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock implementation of contracts.HTTPClient
type MockHTTPClient struct {
	mock.Mock
}

// Get performs a GET request
func (m *MockHTTPClient) Get(ctx context.Context, url string, headers map[string]string) (*contracts.Response, error) {
	args := m.Called(ctx, url, headers)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.Response), args.Error(1)
}

// Post performs a POST request with JSON body
func (m *MockHTTPClient) Post(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	args := m.Called(ctx, url, body, headers)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.Response), args.Error(1)
}

// Put performs a PUT request with JSON body
func (m *MockHTTPClient) Put(
	ctx context.Context,
	url string,
	body interface{},
	headers map[string]string,
) (*contracts.Response, error) {
	args := m.Called(ctx, url, body, headers)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.Response), args.Error(1)
}

// Delete performs a DELETE request
func (m *MockHTTPClient) Delete(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*contracts.Response, error) {
	args := m.Called(ctx, url, headers)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.Response), args.Error(1)
}

// WithTimeout creates a new client with different timeout
func (m *MockHTTPClient) WithTimeout(timeout time.Duration) contracts.HTTPClient {
	args := m.Called(timeout)
	return args.Get(0).(contracts.HTTPClient)
}

// WithRetry creates a new client with different retry configuration
func (m *MockHTTPClient) WithRetry(maxAttempts int, backoffDuration time.Duration) contracts.HTTPClient {
	args := m.Called(maxAttempts, backoffDuration)
	return args.Get(0).(contracts.HTTPClient)
}

// WithBaseURL creates a new client with different base URL
func (m *MockHTTPClient) WithBaseURL(baseURL string) contracts.HTTPClient {
	args := m.Called(baseURL)
	return args.Get(0).(contracts.HTTPClient)
}
