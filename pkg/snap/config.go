package snap

import (
	"fmt"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
)

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

	// Optional: Custom Endpoints for different banks
	CustomEndpoints *CustomEndpoints `json:"customEndpoints,omitempty"`
}

// CustomEndpoints allows overriding default endpoints for different banks
type CustomEndpoints struct {
	// Authentication endpoints
	B2BToken   string `json:"b2bToken,omitempty"`
	B2B2CToken string `json:"b2b2cToken,omitempty"`

	// Virtual Account endpoints
	VirtualAccount *VirtualAccountEndpoints `json:"virtualAccount,omitempty"`

	// MPM endpoints
	MPM *MPMEndpoints `json:"mpm,omitempty"`
}

// VirtualAccountEndpoints defines custom VA endpoints
type VirtualAccountEndpoints struct {
	CreateVA     string `json:"createVA,omitempty"`
	UpdateVA     string `json:"updateVA,omitempty"`
	DeleteVA     string `json:"deleteVA,omitempty"`
	InquiryVA    string `json:"inquiryVA,omitempty"`
	Inquiry      string `json:"inquiry,omitempty"`
	Payment      string `json:"payment,omitempty"`
	Status       string `json:"status,omitempty"`
	Report       string `json:"report,omitempty"`
	UpdateStatus string `json:"updateStatus,omitempty"`
}

// MPMEndpoints defines custom MPM endpoints
type MPMEndpoints struct {
	Transfer       string `json:"transfer,omitempty"`
	Inquiry        string `json:"inquiry,omitempty"`
	Status         string `json:"status,omitempty"`
	Refund         string `json:"refund,omitempty"`
	BalanceInquiry string `json:"balanceInquiry,omitempty"`
	AccountInquiry string `json:"accountInquiry,omitempty"`
	History        string `json:"history,omitempty"`
	GenerateQR     string `json:"generateQR,omitempty"`
	NotifyQR       string `json:"notifyQR,omitempty"`
}

// BankPresets provides predefined configurations for different banks
type BankPresets struct{}

// GetBankConfig returns predefined configuration for specific banks
func (bp *BankPresets) GetBankConfig(bankCode string) *CustomEndpoints {
	switch bankCode {
	case "BCA":
		return bp.getBCAEndpoints()
	case "BNI":
		return bp.getBNIEndpoints()
	case "BRI":
		return bp.getBRIEndpoints()
	case "MANDIRI":
		return bp.getMandiriEndpoints()
	case "CIMB":
		return bp.getCIMBEndpoints()
	case "PERMATA":
		return bp.getPermataEndpoints()
	default:
		return nil // Use default endpoints
	}
}

// Bank-specific endpoint configurations

func (bp *BankPresets) getBCAEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bca/transfer-va/create-va",
			UpdateVA:     "/api/v1.0/bca/transfer-va/update-va",
			DeleteVA:     "/api/v1.0/bca/transfer-va/delete-va",
			InquiryVA:    "/api/v1.0/bca/transfer-va/inquiry-va",
			Inquiry:      "/api/v1.0/bca/transfer-va/inquiry",
			Payment:      "/api/v1.0/bca/transfer-va/payment",
			Status:       "/api/v1.0/bca/transfer-va/status",
			Report:       "/api/v1.0/bca/transfer-va/report",
			UpdateStatus: "/api/v1.0/bca/transfer-va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bca/transfer-kredit/mpm",
			Inquiry:        "/api/v1.0/bca/transfer-kredit/mpm/inquiry",
			Status:         "/api/v1.0/bca/transfer-kredit/mpm/status",
			Refund:         "/api/v1.0/bca/transfer-kredit/mpm/refund",
			BalanceInquiry: "/api/v1.0/bca/transfer-kredit/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/bca/transfer-kredit/mpm/account-inquiry",
			History:        "/api/v1.0/bca/transfer-kredit/mpm/history",
			GenerateQR:     "/api/v1.0/bca/qr/qr-mpm-generate",
			NotifyQR:       "/api/v1.0/bca/qr/qr-mpm-notify",
		},
	}
}

func (bp *BankPresets) getBNIEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bni/virtual-account/create",
			UpdateVA:     "/api/v1.0/bni/virtual-account/update",
			DeleteVA:     "/api/v1.0/bni/virtual-account/delete",
			InquiryVA:    "/api/v1.0/bni/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/bni/virtual-account/inquiry",
			Payment:      "/api/v1.0/bni/virtual-account/payment",
			Status:       "/api/v1.0/bni/virtual-account/status",
			Report:       "/api/v1.0/bni/virtual-account/report",
			UpdateStatus: "/api/v1.0/bni/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bni/mpm/transfer",
			Inquiry:        "/api/v1.0/bni/mpm/inquiry",
			Status:         "/api/v1.0/bni/mpm/status",
			Refund:         "/api/v1.0/bni/mpm/refund",
			BalanceInquiry: "/api/v1.0/bni/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/bni/mpm/account-inquiry",
			History:        "/api/v1.0/bni/mpm/history",
			GenerateQR:     "/api/v1.0/bni/qr/generate",
			NotifyQR:       "/api/v1.0/bni/qr/notify",
		},
	}
}

func (bp *BankPresets) getBRIEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bri/va/create",
			UpdateVA:     "/api/v1.0/bri/va/update",
			DeleteVA:     "/api/v1.0/bri/va/delete",
			InquiryVA:    "/api/v1.0/bri/va/inquiry-va",
			Inquiry:      "/api/v1.0/bri/va/inquiry",
			Payment:      "/api/v1.0/bri/va/payment",
			Status:       "/api/v1.0/bri/va/status",
			Report:       "/api/v1.0/bri/va/report",
			UpdateStatus: "/api/v1.0/bri/va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bri/mpm-transfer/transfer",
			Inquiry:        "/api/v1.0/bri/mpm-transfer/inquiry",
			Status:         "/api/v1.0/bri/mpm-transfer/status",
			Refund:         "/api/v1.0/bri/mpm-transfer/refund",
			BalanceInquiry: "/api/v1.0/bri/mpm-transfer/balance",
			AccountInquiry: "/api/v1.0/bri/mpm-transfer/account",
			History:        "/api/v1.0/bri/mpm-transfer/history",
			GenerateQR:     "/api/v1.0/bri/qr-mpm/generate",
			NotifyQR:       "/api/v1.0/bri/qr-mpm/notify",
		},
	}
}

func (bp *BankPresets) getMandiriEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/mandiri/virtual-account/create-va",
			UpdateVA:     "/api/v1.0/mandiri/virtual-account/update-va",
			DeleteVA:     "/api/v1.0/mandiri/virtual-account/delete-va",
			InquiryVA:    "/api/v1.0/mandiri/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/mandiri/virtual-account/inquiry",
			Payment:      "/api/v1.0/mandiri/virtual-account/payment",
			Status:       "/api/v1.0/mandiri/virtual-account/status",
			Report:       "/api/v1.0/mandiri/virtual-account/report",
			UpdateStatus: "/api/v1.0/mandiri/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/mandiri/transfer-kredit/mpm",
			Inquiry:        "/api/v1.0/mandiri/transfer-kredit/mpm/inquiry",
			Status:         "/api/v1.0/mandiri/transfer-kredit/mpm/status",
			Refund:         "/api/v1.0/mandiri/transfer-kredit/mpm/refund",
			BalanceInquiry: "/api/v1.0/mandiri/transfer-kredit/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/mandiri/transfer-kredit/mpm/account-inquiry",
			History:        "/api/v1.0/mandiri/transfer-kredit/mpm/history",
			GenerateQR:     "/api/v1.0/mandiri/qr/qr-mpm-generate",
			NotifyQR:       "/api/v1.0/mandiri/qr/qr-mpm-notify",
		},
	}
}

func (bp *BankPresets) getCIMBEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/cimb/va/create",
			UpdateVA:     "/api/v1.0/cimb/va/update",
			DeleteVA:     "/api/v1.0/cimb/va/delete",
			InquiryVA:    "/api/v1.0/cimb/va/inquiry-va",
			Inquiry:      "/api/v1.0/cimb/va/inquiry",
			Payment:      "/api/v1.0/cimb/va/payment",
			Status:       "/api/v1.0/cimb/va/status",
			Report:       "/api/v1.0/cimb/va/report",
			UpdateStatus: "/api/v1.0/cimb/va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/cimb/mpm/transfer",
			Inquiry:        "/api/v1.0/cimb/mpm/inquiry",
			Status:         "/api/v1.0/cimb/mpm/status",
			Refund:         "/api/v1.0/cimb/mpm/refund",
			BalanceInquiry: "/api/v1.0/cimb/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/cimb/mpm/account-inquiry",
			History:        "/api/v1.0/cimb/mpm/history",
			GenerateQR:     "/api/v1.0/cimb/qr/generate",
			NotifyQR:       "/api/v1.0/cimb/qr/notify",
		},
	}
}

func (bp *BankPresets) getPermataEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/permata/virtual-account/create",
			UpdateVA:     "/api/v1.0/permata/virtual-account/update",
			DeleteVA:     "/api/v1.0/permata/virtual-account/delete",
			InquiryVA:    "/api/v1.0/permata/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/permata/virtual-account/inquiry",
			Payment:      "/api/v1.0/permata/virtual-account/payment",
			Status:       "/api/v1.0/permata/virtual-account/status",
			Report:       "/api/v1.0/permata/virtual-account/report",
			UpdateStatus: "/api/v1.0/permata/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/permata/mpm/transfer",
			Inquiry:        "/api/v1.0/permata/mpm/inquiry",
			Status:         "/api/v1.0/permata/mpm/status",
			Refund:         "/api/v1.0/permata/mpm/refund",
			BalanceInquiry: "/api/v1.0/permata/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/permata/mpm/account-inquiry",
			History:        "/api/v1.0/permata/mpm/history",
			GenerateQR:     "/api/v1.0/permata/qr/generate",
			NotifyQR:       "/api/v1.0/permata/qr/notify",
		},
	}
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

	// Build endpoints configuration
	endpoints := buildEndpointsConfig(cfg.CustomEndpoints)

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
			Endpoints:      endpoints,
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
			Format: "json",
			Output: "stdout",
		},
	}
}

// buildEndpointsConfig builds the endpoints configuration
func buildEndpointsConfig(customEndpoints *CustomEndpoints) config.EndpointsConfig {
	// Default endpoints
	endpoints := config.EndpointsConfig{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil {
		if customEndpoints.B2BToken != "" {
			endpoints.B2BToken = customEndpoints.B2BToken
		}
		if customEndpoints.B2B2CToken != "" {
			endpoints.B2B2CToken = customEndpoints.B2B2CToken
		}
	}

	return endpoints
}
