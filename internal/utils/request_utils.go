package utils

import (
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/logging"
)

// RequestUtils provides utility functions for API requests
type RequestUtils struct {
	logger logging.Logger
}

// NewRequestUtils creates a new RequestUtils instance
func NewRequestUtils(logger logging.Logger) *RequestUtils {
	return &RequestUtils{
		logger: logger,
	}
}

// GenerateExternalId generates a unique external ID for each request
// Format: timestamp-randomhex
func GenerateExternalId() string {
	// Generate 8 random bytes
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Fallback to timestamp only if random generation fails
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	
	// Combine timestamp and random bytes for uniqueness
	return fmt.Sprintf("%d-%x", time.Now().UnixNano(), randomBytes)
}

// GetClientIP attempts to get the client IP address
// In production, this should be obtained from the client request
func GetClientIP() string {
	// Try to get the outbound IP address
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		// Fallback to localhost if we can't determine the IP
		return "127.0.0.1"
	}
	defer conn.Close()
	
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetDeviceId gets or generates a device ID
// In production, this should be obtained from the client request
func GetDeviceId() string {
	// Try to get hostname as a device identifier
	hostname, err := os.Hostname()
	if err != nil {
		// Fallback to a generated ID if hostname is unavailable
		return fmt.Sprintf("DEVICE-%d", time.Now().Unix())
	}
	
	return fmt.Sprintf("SDK-%s-%d", hostname, os.Getpid())
}

// GetChannelId returns the channel ID based on configuration
// In production, this should be configured based on your channel
func GetChannelId() string {
	// Check for environment variable first
	channelId := os.Getenv("ASPI_CHANNEL_ID")
	if channelId != "" {
		return channelId
	}
	
	// Default channel ID
	return "95221"
}

// GetOriginDomain returns the origin domain for API requests
// In production, this should be configured based on your domain
func GetOriginDomain() string {
	// Check for environment variable first
	origin := os.Getenv("ASPI_ORIGIN_DOMAIN")
	if origin != "" {
		return origin
	}
	
	// Default origin
	return "https://api.yourdomain.com"
}

// GetUserAgent returns the user agent string for API requests
func GetUserAgent(version string) string {
	return fmt.Sprintf("snap-aspi-go/%s", version)
}