package services

import (
	"fmt"
	"time"
)

// generateExternalId generates a unique external ID for each request
// In production, this should be a UUID or other unique identifier
func generateExternalId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// getClientIP gets the client IP address
// In production, this should be obtained from the client request
func getClientIP() string {
	return "127.0.0.1"
}

// getDeviceId gets the client device ID
// In production, this should be obtained from the client request
func getDeviceId() string {
	return "DEVICE-" + fmt.Sprintf("%d", time.Now().Unix())
}

// getChannelId gets the channel ID
// In production, this should be configured based on your channel
func getChannelId() string {
	return "95221"
}
