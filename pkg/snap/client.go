package snap

import (
	"fmt"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
)

// Client represents the main SNAP ASPI SDK client
type Client struct {
	config              *Config
	logger              logging.Logger
	httpClient          contracts.HTTPClient
	authManager         *auth.AccessTokenManager
	virtualAccountSvc   *services.VirtualAccountService
	mpmSvc              *services.MPMService
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

	// Initialize Virtual Account service with custom endpoints
	vaService, err := services.NewVirtualAccountServiceWithEndpoints(
		httpClient, 
		authManager, 
		internalConfig, 
		logger,
		buildVAEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Virtual Account service: %w", err)
	}

	// Initialize MPM service with custom endpoints
	mpmService, err := services.NewMPMServiceWithEndpoints(
		httpClient, 
		authManager, 
		internalConfig, 
		logger,
		buildMPMEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MPM service: %w", err)
	}

	return &Client{
		config:            &cfg,
		logger:            logger,
		httpClient:        httpClient,
		authManager:       authManager,
		virtualAccountSvc: vaService,
		mpmSvc:            mpmService,
	}, nil
}

// NewClientForBank creates a new client with predefined bank configuration
func NewClientForBank(cfg Config, bankCode string) (*Client, error) {
	// Get bank-specific endpoints
	bankPresets := &BankPresets{}
	bankEndpoints := bankPresets.GetBankConfig(bankCode)
	
	if bankEndpoints != nil {
		// Merge bank endpoints with user config
		if cfg.CustomEndpoints == nil {
			cfg.CustomEndpoints = bankEndpoints
		} else {
			// Merge configurations (user config takes precedence)
			mergeEndpoints(cfg.CustomEndpoints, bankEndpoints)
		}
	}

	return NewClient(cfg)
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

// GetConfig returns the current client configuration
func (c *Client) GetConfig() *Config {
	return c.config
}

// Helper functions

// buildVAEndpoints builds VA endpoints map from custom configuration
func buildVAEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default VA endpoints
	endpoints := map[string]string{
		"inquiry":       "/api/v1.0/transfer-va/inquiry",
		"inquiry-va":    "/api/v1.0/transfer-va/inquiry-va",
		"create-va":     "/api/v1.0/transfer-va/create-va",
		"update-va":     "/api/v1.0/transfer-va/update-va",
		"delete-va":     "/api/v1.0/transfer-va/delete-va",
		"payment":       "/api/v1.0/transfer-va/payment",
		"status":        "/api/v1.0/transfer-va/status",
		"report":        "/api/v1.0/transfer-va/report",
		"update-status": "/api/v1.0/transfer-va/update-status",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.VirtualAccount != nil {
		va := customEndpoints.VirtualAccount
		if va.Inquiry != "" {
			endpoints["inquiry"] = va.Inquiry
		}
		if va.InquiryVA != "" {
			endpoints["inquiry-va"] = va.InquiryVA
		}
		if va.CreateVA != "" {
			endpoints["create-va"] = va.CreateVA
		}
		if va.UpdateVA != "" {
			endpoints["update-va"] = va.UpdateVA
		}
		if va.DeleteVA != "" {
			endpoints["delete-va"] = va.DeleteVA
		}
		if va.Payment != "" {
			endpoints["payment"] = va.Payment
		}
		if va.Status != "" {
			endpoints["status"] = va.Status
		}
		if va.Report != "" {
			endpoints["report"] = va.Report
		}
		if va.UpdateStatus != "" {
			endpoints["update-status"] = va.UpdateStatus
		}
	}

	return endpoints
}

// buildMPMEndpoints builds MPM endpoints map from custom configuration
func buildMPMEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default MPM endpoints
	endpoints := map[string]string{
		"transfer":        "/api/v1.0/transfer-kredit/mpm",
		"inquiry":         "/api/v1.0/transfer-kredit/mpm/inquiry",
		"status":          "/api/v1.0/transfer-kredit/mpm/status",
		"refund":          "/api/v1.0/transfer-kredit/mpm/refund",
		"balance-inquiry": "/api/v1.0/transfer-kredit/mpm/balance-inquiry",
		"account-inquiry": "/api/v1.0/transfer-kredit/mpm/account-inquiry",
		"history":         "/api/v1.0/transfer-kredit/mpm/history",
		"generate-qr":     "/api/v1.0/qr/qr-mpm-generate",
		"notify-qr":       "/api/v1.0/qr/qr-mpm-notify",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.MPM != nil {
		mpm := customEndpoints.MPM
		if mpm.Transfer != "" {
			endpoints["transfer"] = mpm.Transfer
		}
		if mpm.Inquiry != "" {
			endpoints["inquiry"] = mpm.Inquiry
		}
		if mpm.Status != "" {
			endpoints["status"] = mpm.Status
		}
		if mpm.Refund != "" {
			endpoints["refund"] = mpm.Refund
		}
		if mpm.BalanceInquiry != "" {
			endpoints["balance-inquiry"] = mpm.BalanceInquiry
		}
		if mpm.AccountInquiry != "" {
			endpoints["account-inquiry"] = mpm.AccountInquiry
		}
		if mpm.History != "" {
			endpoints["history"] = mpm.History
		}
		if mpm.GenerateQR != "" {
			endpoints["generate-qr"] = mpm.GenerateQR
		}
		if mpm.NotifyQR != "" {
			endpoints["notify-qr"] = mpm.NotifyQR
		}
	}

	return endpoints
}

// mergeEndpoints merges bank endpoints with user endpoints (user takes precedence)
func mergeEndpoints(userEndpoints, bankEndpoints *CustomEndpoints) {
	if bankEndpoints.B2BToken != "" && userEndpoints.B2BToken == "" {
		userEndpoints.B2BToken = bankEndpoints.B2BToken
	}
	if bankEndpoints.B2B2CToken != "" && userEndpoints.B2B2CToken == "" {
		userEndpoints.B2B2CToken = bankEndpoints.B2B2CToken
	}

	// Merge VA endpoints
	if bankEndpoints.VirtualAccount != nil {
		if userEndpoints.VirtualAccount == nil {
			userEndpoints.VirtualAccount = bankEndpoints.VirtualAccount
		} else {
			mergeVAEndpoints(userEndpoints.VirtualAccount, bankEndpoints.VirtualAccount)
		}
	}

	// Merge MPM endpoints
	if bankEndpoints.MPM != nil {
		if userEndpoints.MPM == nil {
			userEndpoints.MPM = bankEndpoints.MPM
		} else {
			mergeMPMEndpoints(userEndpoints.MPM, bankEndpoints.MPM)
		}
	}
}

func mergeVAEndpoints(user, bank *VirtualAccountEndpoints) {
	if bank.CreateVA != "" && user.CreateVA == "" {
		user.CreateVA = bank.CreateVA
	}
	if bank.UpdateVA != "" && user.UpdateVA == "" {
		user.UpdateVA = bank.UpdateVA
	}
	if bank.DeleteVA != "" && user.DeleteVA == "" {
		user.DeleteVA = bank.DeleteVA
	}
	if bank.InquiryVA != "" && user.InquiryVA == "" {
		user.InquiryVA = bank.InquiryVA
	}
	if bank.Inquiry != "" && user.Inquiry == "" {
		user.Inquiry = bank.Inquiry
	}
	if bank.Payment != "" && user.Payment == "" {
		user.Payment = bank.Payment
	}
	if bank.Status != "" && user.Status == "" {
		user.Status = bank.Status
	}
	if bank.Report != "" && user.Report == "" {
		user.Report = bank.Report
	}
	if bank.UpdateStatus != "" && user.UpdateStatus == "" {
		user.UpdateStatus = bank.UpdateStatus
	}
}

func mergeMPMEndpoints(user, bank *MPMEndpoints) {
	if bank.Transfer != "" && user.Transfer == "" {
		user.Transfer = bank.Transfer
	}
	if bank.Inquiry != "" && user.Inquiry == "" {
		user.Inquiry = bank.Inquiry
	}
	if bank.Status != "" && user.Status == "" {
		user.Status = bank.Status
	}
	if bank.Refund != "" && user.Refund == "" {
		user.Refund = bank.Refund
	}
	if bank.BalanceInquiry != "" && user.BalanceInquiry == "" {
		user.BalanceInquiry = bank.BalanceInquiry
	}
	if bank.AccountInquiry != "" && user.AccountInquiry == "" {
		user.AccountInquiry = bank.AccountInquiry
	}
	if bank.History != "" && user.History == "" {
		user.History = bank.History
	}
	if bank.GenerateQR != "" && user.GenerateQR == "" {
		user.GenerateQR = bank.GenerateQR
	}
	if bank.NotifyQR != "" && user.NotifyQR == "" {
		user.NotifyQR = bank.NotifyQR
	}
}