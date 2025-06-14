package utils

import (
	"fmt"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
)

// HeaderUtils provides utility functions for working with HTTP headers
type HeaderUtils struct {
	config *config.Config
}

// NewHeaderUtils creates a new HeaderUtils instance
func NewHeaderUtils(cfg *config.Config) *HeaderUtils {
	return &HeaderUtils{
		config: cfg,
	}
}

// CreateStandardHeaders creates standard headers for ASPI API requests
func (h *HeaderUtils) CreateStandardHeaders(
	accessToken string,
	customerAccessToken string,
	signature string,
	timestamp string,
) map[string]string {
	headers := map[string]string{
		"Content-Type":           "application/json",
		"Accept":                 "application/json",
		"X-Client-Key":           h.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          fmt.Sprintf("Bearer %s", accessToken),
		"X-Signature":            signature,
		"X-Origin":               GetOriginDomain(),
		"X-Partner-Id":           h.config.ASPI.ClientID,
		"X-External-Id":          GenerateExternalId(),
		"X-IP-Address":           GetClientIP(),
		"X-Device-Id":            GetDeviceId(),
		"Channel-Id":             GetChannelId(),
		"User-Agent":             GetUserAgent(h.config.App.Version),
	}

	// Add customer token if provided
	if customerAccessToken != "" {
		headers["Authorization-Customer"] = fmt.Sprintf("Bearer %s", customerAccessToken)
	}

	// Add latitude and longitude if available
	if h.config.App.Environment == "production" {
		headers["X-Latitude"] = "-6.1617169"  // Default to Jakarta coordinates
		headers["X-Longitude"] = "106.6643946" // Default to Jakarta coordinates
	}

	return headers
}

// CreateAuthHeaders creates headers for authentication requests
func (h *HeaderUtils) CreateAuthHeaders(signature string) map[string]string {
	timestamp := time.Now().Format(time.RFC3339)
	
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"X-CLIENT-KEY": h.config.ASPI.ClientID,
		"X-TIMESTAMP":  timestamp,
		"X-SIGNATURE":  signature,
	}
}

// CreateCallbackHeaders creates headers for callback responses
func (h *HeaderUtils) CreateCallbackHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"X-TIMESTAMP":  time.Now().Format(time.RFC3339),
	}
}