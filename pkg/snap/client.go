package snap

import (
	"context"
	"fmt"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// Client represents the main SNAP ASPI SDK client
type Client struct {
	config              *config.Config
	logger              logging.Logger
	httpClient          contracts.HTTPClient
	authManager         *auth.AccessTokenManager
	virtualAccountSvc   *services.VirtualAccountService
	mpmSvc              *services.MPMService
}

// Config represents the SDK configuration
type Config struct {
	// ASPI API Configuration
	BaseURL        string `json:"baseUrl"`
	ClientID       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	PrivateKeyPath string `json:"privateKeyPath"`
	PublicKeyPath  string `json:"publicKeyPath"`
	
	// Optional: Environment (sandbox/production)
	Environment string `json:"environment,omitempty"`
	
	// Optional: Timeout configurations
	ConnectTimeout int `json:"connectTimeout,omitempty"` // seconds
	RequestTimeout int `json:"requestTimeout,omitempty"` // seconds
	
	// Optional: Retry configurations
	MaxRetries      int `json:"maxRetries,omitempty"`
	BackoffDuration int `json:"backoffDuration,omitempty"` // seconds
	
	// Optional: Logging
	LogLevel string `json:"logLevel,omitempty"` // debug, info, warn, error
}

// NewClient creates a new SNAP ASPI SDK client
func NewClient(cfg Config) (*Client, error) {
	// Validate required configuration
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Convert SDK config to internal config
	internalConfig := convertToInternalConfig(cfg)

	// Initialize logger
	logger := logging.NewLogger(&internalConfig.Logger)

	// Initialize HTTP client
	httpClient := client.NewOptimizedSnapClient(internalConfig, logger)

	// Initialize auth manager
	authManager, err := auth.NewAccessTokenManager(internalConfig, httpClient, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize auth manager: %w", err)
	}

	// Initialize Virtual Account service
	vaService, err := services.NewVirtualAccountService(httpClient, authManager, internalConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Virtual Account service: %w", err)
	}

	// Initialize MPM service
	mpmService, err := services.NewMPMService(httpClient, authManager, internalConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MPM service: %w", err)
	}

	return &Client{
		config:            internalConfig,
		logger:            logger,
		httpClient:        httpClient,
		authManager:       authManager,
		virtualAccountSvc: vaService,
		mpmSvc:            mpmService,
	}, nil
}

// VirtualAccount returns the Virtual Account service
func (c *Client) VirtualAccount() VirtualAccountService {
	return &virtualAccountService{svc: c.virtualAccountSvc}
}

// MPM returns the MPM (Merchant Payment Management) service
func (c *Client) MPM() MPMService {
	return &mpmService{svc: c.mpmSvc}
}

// Auth returns the authentication service
func (c *Client) Auth() AuthService {
	return &authService{manager: c.authManager}
}

// validateConfig validates the SDK configuration
func validateConfig(cfg Config) error {
	if cfg.BaseURL == "" {
		return fmt.Errorf("baseUrl is required")
	}
	if cfg.ClientID == "" {
		return fmt.Errorf("clientId is required")
	}
	if cfg.ClientSecret == "" {
		return fmt.Errorf("clientSecret is required")
	}
	if cfg.PrivateKeyPath == "" {
		return fmt.Errorf("privateKeyPath is required")
	}
	return nil
}

// convertToInternalConfig converts SDK config to internal config format
func convertToInternalConfig(cfg Config) *config.Config {
	// Set defaults
	environment := cfg.Environment
	if environment == "" {
		environment = "sandbox"
	}
	
	connectTimeout := cfg.ConnectTimeout
	if connectTimeout == 0 {
		connectTimeout = 10
	}
	
	requestTimeout := cfg.RequestTimeout
	if requestTimeout == 0 {
		requestTimeout = 30
	}
	
	maxRetries := cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 3
	}
	
	backoffDuration := cfg.BackoffDuration
	if backoffDuration == 0 {
		backoffDuration = 1
	}
	
	logLevel := cfg.LogLevel
	if logLevel == "" {
		logLevel = "info"
	}

	return &config.Config{
		App: config.AppConfig{
			Name:        "snap-aspi-sdk",
			Version:     "1.0.0",
			Environment: environment,
		},
		ASPI: config.ASPIConfig{
			BaseURL:        cfg.BaseURL,
			ClientID:       cfg.ClientID,
			ClientSecret:   cfg.ClientSecret,
			PrivateKeyPath: cfg.PrivateKeyPath,
			PublicKeyPath:  cfg.PublicKeyPath,
			Endpoints: config.EndpointsConfig{
				B2BToken:   "/api/v1.0/access-token/b2b",
				B2B2CToken: "/api/v1.0/access-token/b2b2c",
			},
			Timeouts: config.TimeoutsConfig{
				ConnectTimeout: fmt.Sprintf("%ds", connectTimeout),
				RequestTimeout: fmt.Sprintf("%ds", requestTimeout),
			},
			Retry: config.RetryConfig{
				MaxAttempts:     maxRetries,
				BackoffDuration: fmt.Sprintf("%ds", backoffDuration),
			},
		},
		Logger: config.LoggerConfig{
			Level:  logLevel,
			Format: "json",
			Output: "stdout",
		},
	}
}