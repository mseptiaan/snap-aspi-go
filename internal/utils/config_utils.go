package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
)

// ConfigUtils provides utility functions for configuration management
type ConfigUtils struct{}

// NewConfigUtils creates a new ConfigUtils instance
func NewConfigUtils() *ConfigUtils {
	return &ConfigUtils{}
}

// LoadConfigFromEnv loads configuration from environment variables
func (c *ConfigUtils) LoadConfigFromEnv() (*config.Config, error) {
	// Load environment variables with defaults
	baseURL := getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id")
	clientID := getEnvOrDefault("ASPI_CLIENT_ID", "")
	clientSecret := getEnvOrDefault("ASPI_CLIENT_SECRET", "")
	privateKeyPath := getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "")
	publicKeyPath := getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "")
	environment := getEnvOrDefault("ASPI_ENVIRONMENT", "sandbox")
	
	// Validate required configuration
	if clientID == "" {
		return nil, errors.NewConfigurationError("ASPI_CLIENT_ID environment variable is required")
	}
	if clientSecret == "" {
		return nil, errors.NewConfigurationError("ASPI_CLIENT_SECRET environment variable is required")
	}
	if privateKeyPath == "" {
		return nil, errors.NewConfigurationError("ASPI_PRIVATE_KEY_PATH environment variable is required")
	}
	
	// Load timeout configurations
	connectTimeout := getEnvIntOrDefault("ASPI_CONNECT_TIMEOUT", 10)
	requestTimeout := getEnvIntOrDefault("ASPI_REQUEST_TIMEOUT", 30)
	
	// Load retry configurations
	maxRetries := getEnvIntOrDefault("ASPI_MAX_RETRIES", 3)
	backoffDuration := getEnvIntOrDefault("ASPI_BACKOFF_DURATION", 1)
	
	// Load logging configurations
	logLevel := getEnvOrDefault("ASPI_LOG_LEVEL", "info")
	logFormat := getEnvOrDefault("ASPI_LOG_FORMAT", "json")
	logOutput := getEnvOrDefault("ASPI_LOG_OUTPUT", "stdout")
	
	// Create configuration
	cfg := &config.Config{
		App: config.AppConfig{
			Name:        "snap-aspi-sdk",
			Version:     "1.1.0",
			Environment: environment,
		},
		ASPI: config.ASPIConfig{
			BaseURL:        baseURL,
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			PrivateKeyPath: privateKeyPath,
			PublicKeyPath:  publicKeyPath,
			Endpoints: config.EndpointsConfig{
				B2BToken:   "/api/v1.0/access-token/b2b",
				B2B2CToken: "/api/v1.0/access-token/b2b2c",
			},
			Timeouts: config.TimeoutsConfig{
				ConnectTimeout: time.Duration(connectTimeout) * time.Second,
				RequestTimeout: time.Duration(requestTimeout) * time.Second,
			},
			Retry: config.RetryConfig{
				MaxAttempts:     maxRetries,
				BackoffDuration: time.Duration(backoffDuration) * time.Second,
			},
		},
		Logger: config.LoggerConfig{
			Level:  logLevel,
			Format: logFormat,
			Output: logOutput,
		},
	}
	
	return cfg, nil
}

// ValidateConfig validates the configuration
func (c *ConfigUtils) ValidateConfig(cfg *config.Config) error {
	if cfg.ASPI.ClientID == "" {
		return errors.NewConfigurationError("client ID is required")
	}
	if cfg.ASPI.ClientSecret == "" {
		return errors.NewConfigurationError("client secret is required")
	}
	if cfg.ASPI.BaseURL == "" {
		return errors.NewConfigurationError("base URL is required")
	}
	if cfg.ASPI.PrivateKeyPath == "" {
		return errors.NewConfigurationError("private key path is required")
	}
	
	// Validate private key file exists
	if _, err := os.Stat(cfg.ASPI.PrivateKeyPath); os.IsNotExist(err) {
		return errors.NewConfigurationError(fmt.Sprintf("private key file not found: %s", cfg.ASPI.PrivateKeyPath))
	}
	
	// Validate public key file exists if specified
	if cfg.ASPI.PublicKeyPath != "" {
		if _, err := os.Stat(cfg.ASPI.PublicKeyPath); os.IsNotExist(err) {
			return errors.NewConfigurationError(fmt.Sprintf("public key file not found: %s", cfg.ASPI.PublicKeyPath))
		}
	}
	
	return nil
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvIntOrDefault gets an integer environment variable or returns a default value
func getEnvIntOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	
	return intValue
}

// GetBankEndpoints returns bank-specific endpoints based on bank code
func (c *ConfigUtils) GetBankEndpoints(bankCode string) map[string]string {
	// Convert bank code to uppercase
	bankCode = strings.ToUpper(bankCode)
	
	// Define bank-specific endpoints
	switch bankCode {
	case "BCA":
		return map[string]string{
			"b2b_token":        "/api/v1.0/bca/access-token/b2b",
			"b2b2c_token":      "/api/v1.0/bca/access-token/b2b2c",
			"create_va":        "/api/v1.0/bca/transfer-va/create-va",
			"account_inquiry":  "/api/v1.0/bca/account-inquiry-external",
			"balance_inquiry":  "/api/v1.0/bca/balance-inquiry",
			"generate_qr":      "/api/v1.0/bca/qr/qr-mpm-generate",
		}
	case "BNI":
		return map[string]string{
			"b2b_token":        "/api/v1.0/bni/access-token/b2b",
			"b2b2c_token":      "/api/v1.0/bni/access-token/b2b2c",
			"create_va":        "/api/v1.0/bni/virtual-account/create",
			"account_inquiry":  "/api/v1.0/bni/account-inquiry-external",
			"balance_inquiry":  "/api/v1.0/bni/balance-inquiry",
			"generate_qr":      "/api/v1.0/bni/qr/qr-mpm-generate",
		}
	case "BRI":
		return map[string]string{
			"b2b_token":        "/api/v1.0/bri/access-token/b2b",
			"b2b2c_token":      "/api/v1.0/bri/access-token/b2b2c",
			"create_va":        "/api/v1.0/bri/va/create",
			"account_inquiry":  "/api/v1.0/bri/account-inquiry-external",
			"balance_inquiry":  "/api/v1.0/bri/balance-inquiry",
			"generate_qr":      "/api/v1.0/bri/qr/qr-mpm-generate",
		}
	case "MANDIRI":
		return map[string]string{
			"b2b_token":        "/api/v1.0/mandiri/access-token/b2b",
			"b2b2c_token":      "/api/v1.0/mandiri/access-token/b2b2c",
			"create_va":        "/api/v1.0/mandiri/virtual-account/create-va",
			"account_inquiry":  "/api/v1.0/mandiri/account-inquiry-external",
			"balance_inquiry":  "/api/v1.0/mandiri/balance-inquiry",
			"generate_qr":      "/api/v1.0/mandiri/qr/qr-mpm-generate",
		}
	default:
		// Return default endpoints
		return map[string]string{
			"b2b_token":        "/api/v1.0/access-token/b2b",
			"b2b2c_token":      "/api/v1.0/access-token/b2b2c",
			"create_va":        "/api/v1.0/transfer-va/create-va",
			"account_inquiry":  "/api/v1.0/account-inquiry-external",
			"balance_inquiry":  "/api/v1.0/balance-inquiry",
			"generate_qr":      "/api/v1.0/qr/qr-mpm-generate",
		}
	}
}