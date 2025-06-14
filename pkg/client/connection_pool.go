package client

import (
	"net"
	"net/http"
	"time"
)

// CreateOptimizedHTTPClient creates an HTTP client with connection pooling and optimizations
func CreateOptimizedHTTPClient(connectTimeout, requestTimeout time.Duration) *http.Client {
	// Create custom transport with connection pooling
	transport := &http.Transport{
		// Connection pooling settings
		MaxIdleConns:        100,              // Maximum idle connections across all hosts
		MaxIdleConnsPerHost: 10,               // Maximum idle connections per host
		MaxConnsPerHost:     50,               // Maximum connections per host
		IdleConnTimeout:     90 * time.Second, // How long idle connections are kept

		// Timeouts
		DialContext: (&net.Dialer{
			Timeout:   connectTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// Enable HTTP/2
		ForceAttemptHTTP2: true,

		// Disable compression for better performance with JSON
		DisableCompression: false,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   requestTimeout,
	}
}