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
	
	// Registration endpoints
	Registration *RegistrationEndpoints `json:"registration,omitempty"`
	
	// Balance Inquiry endpoints
	BalanceInquiry *BalanceInquiryEndpoints `json:"balanceInquiry,omitempty"`
	
	// Transaction History endpoints
	TransactionHistory *TransactionHistoryEndpoints `json:"transactionHistory,omitempty"`
	
	// Transfer Credit endpoints
	TransferCredit *TransferCreditEndpoints `json:"transferCredit,omitempty"`
	
	// Transfer Debit endpoints
	TransferDebit *TransferDebitEndpoints `json:"transferDebit,omitempty"`
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

// RegistrationEndpoints defines custom Registration endpoints
type RegistrationEndpoints struct {
	Register     string `json:"register,omitempty"`
	RegisterCard string `json:"registerCard,omitempty"`
	BindAccount  string `json:"bindAccount,omitempty"`
	VerifyOTP    string `json:"verifyOTP,omitempty"`
}

// BalanceInquiryEndpoints defines custom Balance Inquiry endpoints
type BalanceInquiryEndpoints struct {
	BalanceInquiry string `json:"balanceInquiry,omitempty"`
}

// TransactionHistoryEndpoints defines custom Transaction History endpoints
type TransactionHistoryEndpoints struct {
	TransactionHistoryList   string `json:"transactionHistoryList,omitempty"`
	TransactionHistoryDetail string `json:"transactionHistoryDetail,omitempty"`
	BankStatement            string `json:"bankStatement,omitempty"`
}

// TransferCreditEndpoints defines custom Transfer Credit endpoints
type TransferCreditEndpoints struct {
	AccountInquiry        string `json:"accountInquiry,omitempty"`
	TriggerTransfer       string `json:"triggerTransfer,omitempty"`
	TransferStatus        string `json:"transferStatus,omitempty"`
	CustomerTopUp         string `json:"customerTopUp,omitempty"`
	BulkCashin            string `json:"bulkCashin,omitempty"`
	TransferToBank        string `json:"transferToBank,omitempty"`
	TransferToOTC         string `json:"transferToOTC,omitempty"`
	TransferToOTCStatus   string `json:"transferToOTCStatus,omitempty"`
	TransferToOTCCancel   string `json:"transferToOTCCancel,omitempty"`
}

// TransferDebitEndpoints defines custom Transfer Debit endpoints
type TransferDebitEndpoints struct {
	DirectDebitPayment       string `json:"directDebitPayment,omitempty"`
	DirectDebitStatus        string `json:"directDebitStatus,omitempty"`
	DirectDebitCancel        string `json:"directDebitCancel,omitempty"`
	DirectDebitRefund        string `json:"directDebitRefund,omitempty"`
	CPMGenerateQR            string `json:"cpmGenerateQR,omitempty"`
	CPMPayment               string `json:"cpmPayment,omitempty"`
	AuthPayment              string `json:"authPayment,omitempty"`
	AuthCapture              string `json:"authCapture,omitempty"`
	AuthVoid                 string `json:"authVoid,omitempty"`
	DirectDebitBIFAST        string `json:"directDebitBIFAST,omitempty"`
	DirectDebitBIFASTPayment string `json:"directDebitBIFASTPayment,omitempty"`
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bca/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bca/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bca/transaction-history-detail",
			BankStatement:            "/api/v1.0/bca/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bca/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bca/trigger-transfer",
			TransferStatus:  "/api/v1.0/bca/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bca/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bca/debit/status",
			DirectDebitCancel:  "/api/v1.0/bca/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bca/debit/refund",
			AuthPayment:        "/api/v1.0/bca/auth/payment",
			AuthCapture:        "/api/v1.0/bca/auth/capture",
			AuthVoid:           "/api/v1.0/bca/auth/void",
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bni/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bni/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bni/transaction-history-detail",
			BankStatement:            "/api/v1.0/bni/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bni/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bni/trigger-transfer",
			TransferStatus:  "/api/v1.0/bni/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bni/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bni/debit/status",
			DirectDebitCancel:  "/api/v1.0/bni/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bni/debit/refund",
			AuthPayment:        "/api/v1.0/bni/auth/payment",
			AuthCapture:        "/api/v1.0/bni/auth/capture",
			AuthVoid:           "/api/v1.0/bni/auth/void",
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bri/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bri/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bri/transaction-history-detail",
			BankStatement:            "/api/v1.0/bri/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bri/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bri/trigger-transfer",
			TransferStatus:  "/api/v1.0/bri/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bri/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bri/debit/status",
			DirectDebitCancel:  "/api/v1.0/bri/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bri/debit/refund",
			AuthPayment:        "/api/v1.0/bri/auth/payment",
			AuthCapture:        "/api/v1.0/bri/auth/capture",
			AuthVoid:           "/api/v1.0/bri/auth/void",
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/mandiri/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/mandiri/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/mandiri/transaction-history-detail",
			BankStatement:            "/api/v1.0/mandiri/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/mandiri/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/mandiri/trigger-transfer",
			TransferStatus:  "/api/v1.0/mandiri/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/mandiri/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/mandiri/debit/status",
			DirectDebitCancel:  "/api/v1.0/mandiri/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/mandiri/debit/refund",
			AuthPayment:        "/api/v1.0/mandiri/auth/payment",
			AuthCapture:        "/api/v1.0/mandiri/auth/capture",
			AuthVoid:           "/api/v1.0/mandiri/auth/void",
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/cimb/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/cimb/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/cimb/transaction-history-detail",
			BankStatement:            "/api/v1.0/cimb/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/cimb/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/cimb/trigger-transfer",
			TransferStatus:  "/api/v1.0/cimb/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/cimb/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/cimb/debit/status",
			DirectDebitCancel:  "/api/v1.0/cimb/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/cimb/debit/refund",
			AuthPayment:        "/api/v1.0/cimb/auth/payment",
			AuthCapture:        "/api/v1.0/cimb/auth/capture",
			AuthVoid:           "/api/v1.0/cimb/auth/void",
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
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/permata/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/permata/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/permata/transaction-history-detail",
			BankStatement:            "/api/v1.0/permata/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/permata/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/permata/trigger-transfer",
			TransferStatus:  "/api/v1.0/permata/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/permata/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/permata/debit/status",
			DirectDebitCancel:  "/api/v1.0/permata/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/permata/debit/refund",
			AuthPayment:        "/api/v1.0/permata/auth/payment",
			AuthCapture:        "/api/v1.0/permata/auth/capture",
			AuthVoid:           "/api/v1.0/permata/auth/void",
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