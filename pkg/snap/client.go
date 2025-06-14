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
	config                 *Config
	logger                 logging.Logger
	httpClient             contracts.HTTPClient
	authManager            *auth.AccessTokenManager
	virtualAccountSvc      *services.VirtualAccountService
	mpmSvc                 *services.MPMService
	registrationSvc        *services.RegistrationService
	balanceInquirySvc      *services.BalanceInquiryService
	transactionHistorySvc  *services.TransactionHistoryService
	transferCreditSvc      *services.TransferCreditService
	transferDebitSvc       *services.TransferDebitService
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

	// Initialize Registration service with custom endpoints
	registrationService, err := services.NewRegistrationServiceWithEndpoints(
		httpClient,
		authManager,
		internalConfig,
		logger,
		buildRegistrationEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Registration service: %w", err)
	}

	// Initialize Balance Inquiry service with custom endpoints
	balanceInquiryService, err := services.NewBalanceInquiryServiceWithEndpoints(
		httpClient,
		authManager,
		internalConfig,
		logger,
		buildBalanceInquiryEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Balance Inquiry service: %w", err)
	}

	// Initialize Transaction History service with custom endpoints
	transactionHistoryService, err := services.NewTransactionHistoryServiceWithEndpoints(
		httpClient,
		authManager,
		internalConfig,
		logger,
		buildTransactionHistoryEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Transaction History service: %w", err)
	}

	// Initialize Transfer Credit service with custom endpoints
	transferCreditService, err := services.NewTransferCreditServiceWithEndpoints(
		httpClient,
		authManager,
		internalConfig,
		logger,
		buildTransferCreditEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Transfer Credit service: %w", err)
	}

	// Initialize Transfer Debit service with custom endpoints
	transferDebitService, err := services.NewTransferDebitServiceWithEndpoints(
		httpClient,
		authManager,
		internalConfig,
		logger,
		buildTransferDebitEndpoints(cfg.CustomEndpoints),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Transfer Debit service: %w", err)
	}

	return &Client{
		config:                &cfg,
		logger:                logger,
		httpClient:            httpClient,
		authManager:           authManager,
		virtualAccountSvc:     vaService,
		mpmSvc:                mpmService,
		registrationSvc:       registrationService,
		balanceInquirySvc:     balanceInquiryService,
		transactionHistorySvc: transactionHistoryService,
		transferCreditSvc:     transferCreditService,
		transferDebitSvc:      transferDebitService,
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

// Registration returns the registration service
func (c *Client) Registration() RegistrationService {
	return &registrationService{svc: c.registrationSvc}
}

// BalanceInquiry returns the balance inquiry service
func (c *Client) BalanceInquiry() BalanceInquiryService {
	return &balanceInquiryService{svc: c.balanceInquirySvc}
}

// TransactionHistory returns the transaction history service
func (c *Client) TransactionHistory() TransactionHistoryService {
	return &transactionHistoryService{svc: c.transactionHistorySvc}
}

// TransferCredit returns the transfer credit service
func (c *Client) TransferCredit() TransferCreditService {
	return &transferCreditService{svc: c.transferCreditSvc}
}

// TransferDebit returns the transfer debit service
func (c *Client) TransferDebit() TransferDebitService {
	return &transferDebitService{svc: c.transferDebitSvc}
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

// buildRegistrationEndpoints builds Registration endpoints map from custom configuration
func buildRegistrationEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default Registration endpoints
	endpoints := map[string]string{
		"register":      "/api/v1.0/registration/register",
		"register-card": "/api/v1.0/registration/card-registration",
		"bind-account":  "/api/v1.0/registration/account-binding",
		"verify-otp":    "/api/v1.0/registration/otp-verification",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.Registration != nil {
		reg := customEndpoints.Registration
		if reg.Register != "" {
			endpoints["register"] = reg.Register
		}
		if reg.RegisterCard != "" {
			endpoints["register-card"] = reg.RegisterCard
		}
		if reg.BindAccount != "" {
			endpoints["bind-account"] = reg.BindAccount
		}
		if reg.VerifyOTP != "" {
			endpoints["verify-otp"] = reg.VerifyOTP
		}
	}

	return endpoints
}

// buildBalanceInquiryEndpoints builds Balance Inquiry endpoints map from custom configuration
func buildBalanceInquiryEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default Balance Inquiry endpoints
	endpoints := map[string]string{
		"balance-inquiry": "/api/v1.0/balance-inquiry",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.BalanceInquiry != nil {
		bi := customEndpoints.BalanceInquiry
		if bi.BalanceInquiry != "" {
			endpoints["balance-inquiry"] = bi.BalanceInquiry
		}
	}

	return endpoints
}

// buildTransactionHistoryEndpoints builds Transaction History endpoints map from custom configuration
func buildTransactionHistoryEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default Transaction History endpoints
	endpoints := map[string]string{
		"transaction-history-list":   "/api/v1.0/transaction-history-list",
		"transaction-history-detail": "/api/v1.0/transaction-history-detail",
		"bank-statement":             "/api/v1.0/bank-statement",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.TransactionHistory != nil {
		th := customEndpoints.TransactionHistory
		if th.TransactionHistoryList != "" {
			endpoints["transaction-history-list"] = th.TransactionHistoryList
		}
		if th.TransactionHistoryDetail != "" {
			endpoints["transaction-history-detail"] = th.TransactionHistoryDetail
		}
		if th.BankStatement != "" {
			endpoints["bank-statement"] = th.BankStatement
		}
	}

	return endpoints
}

// buildTransferCreditEndpoints builds Transfer Credit endpoints map from custom configuration
func buildTransferCreditEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default Transfer Credit endpoints
	endpoints := map[string]string{
		"account-inquiry":        "/api/v1.0/account-inquiry-external",
		"trigger-transfer":       "/api/v1.0/trigger-transfer",
		"transfer-status":        "/api/v1.0/transfer/status",
		"customer-top-up":        "/api/v1.0/emoney/topup",
		"bulk-cashin":            "/api/v1.0/emoney/bulk-cashin-payment",
		"transfer-to-bank":       "/api/v1.0/emoney/transfer-bank",
		"transfer-to-otc":        "/api/v1.0/emoney/otc-cashout",
		"transfer-to-otc-status": "/api/v1.0/emoney/otc-status",
		"transfer-to-otc-cancel": "/api/v1.0/emoney/otc-cancel",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.TransferCredit != nil {
		tc := customEndpoints.TransferCredit
		if tc.AccountInquiry != "" {
			endpoints["account-inquiry"] = tc.AccountInquiry
		}
		if tc.TriggerTransfer != "" {
			endpoints["trigger-transfer"] = tc.TriggerTransfer
		}
		if tc.TransferStatus != "" {
			endpoints["transfer-status"] = tc.TransferStatus
		}
		if tc.CustomerTopUp != "" {
			endpoints["customer-top-up"] = tc.CustomerTopUp
		}
		if tc.BulkCashin != "" {
			endpoints["bulk-cashin"] = tc.BulkCashin
		}
		if tc.TransferToBank != "" {
			endpoints["transfer-to-bank"] = tc.TransferToBank
		}
		if tc.TransferToOTC != "" {
			endpoints["transfer-to-otc"] = tc.TransferToOTC
		}
		if tc.TransferToOTCStatus != "" {
			endpoints["transfer-to-otc-status"] = tc.TransferToOTCStatus
		}
		if tc.TransferToOTCCancel != "" {
			endpoints["transfer-to-otc-cancel"] = tc.TransferToOTCCancel
		}
	}

	return endpoints
}

// buildTransferDebitEndpoints builds Transfer Debit endpoints map from custom configuration
func buildTransferDebitEndpoints(customEndpoints *CustomEndpoints) map[string]string {
	// Default Transfer Debit endpoints
	endpoints := map[string]string{
		"direct-debit-payment":       "/api/v1.0/debit/payment-host-to-host",
		"direct-debit-status":        "/api/v1.0/debit/status",
		"direct-debit-cancel":        "/api/v1.0/debit/cancel",
		"direct-debit-refund":        "/api/v1.0/debit/refund",
		"cpm-generate-qr":            "/api/v1.0/qr/qr-cpm-generate",
		"cpm-payment":                "/api/v1.0/qr/qr-cpm-payment",
		"auth-payment":               "/api/v1.0/auth/payment",
		"auth-capture":               "/api/v1.0/auth/capture",
		"auth-void":                  "/api/v1.0/auth/void",
		"direct-debit-bifast":        "/api/v1.0/debit/fast-emandate",
		"direct-debit-bifast-payment": "/api/v1.0/debit/fast-payment",
	}

	// Override with custom endpoints if provided
	if customEndpoints != nil && customEndpoints.TransferDebit != nil {
		td := customEndpoints.TransferDebit
		if td.DirectDebitPayment != "" {
			endpoints["direct-debit-payment"] = td.DirectDebitPayment
		}
		if td.DirectDebitStatus != "" {
			endpoints["direct-debit-status"] = td.DirectDebitStatus
		}
		if td.DirectDebitCancel != "" {
			endpoints["direct-debit-cancel"] = td.DirectDebitCancel
		}
		if td.DirectDebitRefund != "" {
			endpoints["direct-debit-refund"] = td.DirectDebitRefund
		}
		if td.CPMGenerateQR != "" {
			endpoints["cpm-generate-qr"] = td.CPMGenerateQR
		}
		if td.CPMPayment != "" {
			endpoints["cpm-payment"] = td.CPMPayment
		}
		if td.AuthPayment != "" {
			endpoints["auth-payment"] = td.AuthPayment
		}
		if td.AuthCapture != "" {
			endpoints["auth-capture"] = td.AuthCapture
		}
		if td.AuthVoid != "" {
			endpoints["auth-void"] = td.AuthVoid
		}
		if td.DirectDebitBIFAST != "" {
			endpoints["direct-debit-bifast"] = td.DirectDebitBIFAST
		}
		if td.DirectDebitBIFASTPayment != "" {
			endpoints["direct-debit-bifast-payment"] = td.DirectDebitBIFASTPayment
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

	// Merge Registration endpoints
	if bankEndpoints.Registration != nil {
		if userEndpoints.Registration == nil {
			userEndpoints.Registration = bankEndpoints.Registration
		} else {
			mergeRegistrationEndpoints(userEndpoints.Registration, bankEndpoints.Registration)
		}
	}

	// Merge Balance Inquiry endpoints
	if bankEndpoints.BalanceInquiry != nil {
		if userEndpoints.BalanceInquiry == nil {
			userEndpoints.BalanceInquiry = bankEndpoints.BalanceInquiry
		} else {
			mergeBalanceInquiryEndpoints(userEndpoints.BalanceInquiry, bankEndpoints.BalanceInquiry)
		}
	}

	// Merge Transaction History endpoints
	if bankEndpoints.TransactionHistory != nil {
		if userEndpoints.TransactionHistory == nil {
			userEndpoints.TransactionHistory = bankEndpoints.TransactionHistory
		} else {
			mergeTransactionHistoryEndpoints(userEndpoints.TransactionHistory, bankEndpoints.TransactionHistory)
		}
	}

	// Merge Transfer Credit endpoints
	if bankEndpoints.TransferCredit != nil {
		if userEndpoints.TransferCredit == nil {
			userEndpoints.TransferCredit = bankEndpoints.TransferCredit
		} else {
			mergeTransferCreditEndpoints(userEndpoints.TransferCredit, bankEndpoints.TransferCredit)
		}
	}

	// Merge Transfer Debit endpoints
	if bankEndpoints.TransferDebit != nil {
		if userEndpoints.TransferDebit == nil {
			userEndpoints.TransferDebit = bankEndpoints.TransferDebit
		} else {
			mergeTransferDebitEndpoints(userEndpoints.TransferDebit, bankEndpoints.TransferDebit)
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

func mergeRegistrationEndpoints(user, bank *RegistrationEndpoints) {
	if bank.Register != "" && user.Register == "" {
		user.Register = bank.Register
	}
	if bank.RegisterCard != "" && user.RegisterCard == "" {
		user.RegisterCard = bank.RegisterCard
	}
	if bank.BindAccount != "" && user.BindAccount == "" {
		user.BindAccount = bank.BindAccount
	}
	if bank.VerifyOTP != "" && user.VerifyOTP == "" {
		user.VerifyOTP = bank.VerifyOTP
	}
}

func mergeBalanceInquiryEndpoints(user, bank *BalanceInquiryEndpoints) {
	if bank.BalanceInquiry != "" && user.BalanceInquiry == "" {
		user.BalanceInquiry = bank.BalanceInquiry
	}
}

func mergeTransactionHistoryEndpoints(user, bank *TransactionHistoryEndpoints) {
	if bank.TransactionHistoryList != "" && user.TransactionHistoryList == "" {
		user.TransactionHistoryList = bank.TransactionHistoryList
	}
	if bank.TransactionHistoryDetail != "" && user.TransactionHistoryDetail == "" {
		user.TransactionHistoryDetail = bank.TransactionHistoryDetail
	}
	if bank.BankStatement != "" && user.BankStatement == "" {
		user.BankStatement = bank.BankStatement
	}
}

func mergeTransferCreditEndpoints(user, bank *TransferCreditEndpoints) {
	if bank.AccountInquiry != "" && user.AccountInquiry == "" {
		user.AccountInquiry = bank.AccountInquiry
	}
	if bank.TriggerTransfer != "" && user.TriggerTransfer == "" {
		user.TriggerTransfer = bank.TriggerTransfer
	}
	if bank.TransferStatus != "" && user.TransferStatus == "" {
		user.TransferStatus = bank.TransferStatus
	}
	if bank.CustomerTopUp != "" && user.CustomerTopUp == "" {
		user.CustomerTopUp = bank.CustomerTopUp
	}
	if bank.BulkCashin != "" && user.BulkCashin == "" {
		user.BulkCashin = bank.BulkCashin
	}
	if bank.TransferToBank != "" && user.TransferToBank == "" {
		user.TransferToBank = bank.TransferToBank
	}
	if bank.TransferToOTC != "" && user.TransferToOTC == "" {
		user.TransferToOTC = bank.TransferToOTC
	}
	if bank.TransferToOTCStatus != "" && user.TransferToOTCStatus == "" {
		user.TransferToOTCStatus = bank.TransferToOTCStatus
	}
	if bank.TransferToOTCCancel != "" && user.TransferToOTCCancel == "" {
		user.TransferToOTCCancel = bank.TransferToOTCCancel
	}
}

func mergeTransferDebitEndpoints(user, bank *TransferDebitEndpoints) {
	if bank.DirectDebitPayment != "" && user.DirectDebitPayment == "" {
		user.DirectDebitPayment = bank.DirectDebitPayment
	}
	if bank.DirectDebitStatus != "" && user.DirectDebitStatus == "" {
		user.DirectDebitStatus = bank.DirectDebitStatus
	}
	if bank.DirectDebitCancel != "" && user.DirectDebitCancel == "" {
		user.DirectDebitCancel = bank.DirectDebitCancel
	}
	if bank.DirectDebitRefund != "" && user.DirectDebitRefund == "" {
		user.DirectDebitRefund = bank.DirectDebitRefund
	}
	if bank.CPMGenerateQR != "" && user.CPMGenerateQR == "" {
		user.CPMGenerateQR = bank.CPMGenerateQR
	}
	if bank.CPMPayment != "" && user.CPMPayment == "" {
		user.CPMPayment = bank.CPMPayment
	}
	if bank.AuthPayment != "" && user.AuthPayment == "" {
		user.AuthPayment = bank.AuthPayment
	}
	if bank.AuthCapture != "" && user.AuthCapture == "" {
		user.AuthCapture = bank.AuthCapture
	}
	if bank.AuthVoid != "" && user.AuthVoid == "" {
		user.AuthVoid = bank.AuthVoid
	}
	if bank.DirectDebitBIFAST != "" && user.DirectDebitBIFAST == "" {
		user.DirectDebitBIFAST = bank.DirectDebitBIFAST
	}
	if bank.DirectDebitBIFASTPayment != "" && user.DirectDebitBIFASTPayment == "" {
		user.DirectDebitBIFASTPayment = bank.DirectDebitBIFASTPayment
	}
}