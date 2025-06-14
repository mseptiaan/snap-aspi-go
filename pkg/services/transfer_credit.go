package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/signature"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// TransferCreditService implements the Transfer Credit business logic
type TransferCreditService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI Transfer Credit API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewTransferCreditService creates a new Transfer Credit service instance
func NewTransferCreditService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*TransferCreditService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &TransferCreditService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Transfer Credit API endpoints
		endpoints: map[string]string{
			"account-inquiry":        "/api/v1.0/account-inquiry-external",
			"trigger-transfer":       "/api/v1.0/trigger-transfer",
			"transfer-status":        "/api/v1.0/transfer/status",
			"customer-top-up":        "/api/v1.0/emoney/topup",
			"bulk-cashin":            "/api/v1.0/emoney/bulk-cashin-payment",
			"transfer-to-bank":       "/api/v1.0/emoney/transfer-bank",
			"transfer-to-otc":        "/api/v1.0/emoney/otc-cashout",
			"transfer-to-otc-status": "/api/v1.0/emoney/otc-status",
			"transfer-to-otc-cancel": "/api/v1.0/emoney/otc-cancel",
		},
		methods: map[string]string{
			"account-inquiry":        "POST",
			"trigger-transfer":       "POST",
			"transfer-status":        "POST",
			"customer-top-up":        "POST",
			"bulk-cashin":            "POST",
			"transfer-to-bank":       "POST",
			"transfer-to-otc":        "POST",
			"transfer-to-otc-status": "POST",
			"transfer-to-otc-cancel": "POST",
		},
	}, nil
}

// NewTransferCreditServiceWithEndpoints creates a new Transfer Credit service with custom endpoints
func NewTransferCreditServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*TransferCreditService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getTransferCreditEndpoints(customEndpoints)
	methods := getTransferCreditMethods()

	return &TransferCreditService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getTransferCreditEndpoints returns Transfer Credit endpoints (custom or default)
func getTransferCreditEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
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

	// If custom endpoints provided, merge them
	if customEndpoints != nil {
		for key, endpoint := range customEndpoints {
			if endpoint != "" {
				defaultEndpoints[key] = endpoint
			}
		}
	}

	return defaultEndpoints
}

// getTransferCreditMethods returns HTTP methods for Transfer Credit endpoints
func getTransferCreditMethods() map[string]string {
	return map[string]string{
		"account-inquiry":        "POST",
		"trigger-transfer":       "POST",
		"transfer-status":        "POST",
		"customer-top-up":        "POST",
		"bulk-cashin":            "POST",
		"transfer-to-bank":       "POST",
		"transfer-to-otc":        "POST",
		"transfer-to-otc-status": "POST",
		"transfer-to-otc-cancel": "POST",
	}
}

// AccountInquiry performs account inquiry
func (t *TransferCreditService) AccountInquiry(ctx context.Context, payload *types.AccountInquiryPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "account-inquiry", payload)
}

// TriggerTransfer performs transfer
func (t *TransferCreditService) TriggerTransfer(ctx context.Context, payload *types.TriggerTransferPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "trigger-transfer", payload)
}

// TransferStatusInquiry checks transfer status
func (t *TransferCreditService) TransferStatusInquiry(ctx context.Context, payload *types.TransferStatusInquiryPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "transfer-status", payload)
}

// CustomerTopUp performs customer top up
func (t *TransferCreditService) CustomerTopUp(ctx context.Context, payload *types.CustomerTopUpPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "customer-top-up", payload)
}

// BulkCashin performs bulk cashin
func (t *TransferCreditService) BulkCashin(ctx context.Context, payload *types.BulkCashinPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "bulk-cashin", payload)
}

// TransferToBank performs transfer to bank
func (t *TransferCreditService) TransferToBank(ctx context.Context, payload *types.TransferToBankPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "transfer-to-bank", payload)
}

// TransferToOTC performs transfer to OTC
func (t *TransferCreditService) TransferToOTC(ctx context.Context, payload *types.TransferToOTCPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "transfer-to-otc", payload)
}

// makeRequest handles the HTTP request for ASPI Transfer Credit API calls
func (t *TransferCreditService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
	// Validate endpoint
	path, exists := t.endpoints[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("endpoint '%s' not found", endpoint))
	}

	method, exists := t.methods[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("method for endpoint '%s' not found", endpoint))
	}

	// Generate timestamp for request
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Get B2B access token
	b2bToken, err := t.authManager.GetAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get B2B access token: %w", err)
	}

	// Get B2B2C access token
	// In a production environment, this should be obtained from the user's session
	// or passed in as a parameter rather than using hardcoded values
	b2b2cTokenReq := auth.CustomerTokenRequest{
		AuthCode:       "a6975f82-d00a-4ddc-9633-087fefb6275e",
		RefreshToken:   "83a58570-6795-11ec-90d6-0242ac120003",
		AdditionalInfo: map[string]interface{}{},
	}

	b2b2cToken, err := t.authManager.GetCustomerAccessToken(ctx, b2b2cTokenReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get B2B2C access token: %w", err)
	}

	// Validate payload can be marshaled to JSON
	if _, err := json.Marshal(payload); err != nil {
		return nil, errors.NewValidationError(fmt.Sprintf("failed to marshal payload: %v", err))
	}

	// Generate signature
	timestampTime := time.Now()

	sigPayload := signature.SymmetricPayload{
		HTTPMethod:  method,
		EndpointURL: path,
		AccessToken: b2bToken.AccessToken,
		RequestBody: payload,
		Timestamp:   timestampTime,
	}

	signatureValue, err := t.symSigner.Sign(sigPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Build request headers with proper values for production use
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           t.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "https://api.yourdomain.com", // Should be configured based on your domain
		"X-Partner-Id":           t.config.ASPI.ClientID,
		"X-External-Id":          generateExternalId(), // Should generate a unique ID for each request
		"X-IP-Address":           getClientIP(),        // Should get from client request
		"X-Device-Id":            getDeviceId(),        // Should get from client request
		"Channel-Id":             getChannelId(),       // Should be configured based on your channel
	}

	// Log the request
	t.logger.WithFields(map[string]any{
		"service":   "TransferCreditService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making Transfer Credit API request")

	// Make the HTTP request
	var response *contracts.Response

	url := t.config.ASPI.BaseURL + path

	switch method {
	case "POST":
		response, err = t.httpClient.Post(ctx, url, payload, requestHeaders)
	case "PUT":
		response, err = t.httpClient.Put(ctx, url, payload, requestHeaders)
	case "DELETE":
		response, err = t.httpClient.Delete(ctx, url, requestHeaders)
	case "GET":
		response, err = t.httpClient.Get(ctx, url, requestHeaders)
	default:
		return nil, errors.NewValidationError(fmt.Sprintf("unsupported HTTP method: %s", method))
	}

	if err != nil {
		t.logger.WithFields(map[string]any{
			"service":  "TransferCreditService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("Transfer Credit API request failed")
		return nil, fmt.Errorf("Transfer Credit API request failed: %w", err)
	}

	// Log successful response
	t.logger.WithFields(map[string]any{
		"service":     "TransferCreditService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("Transfer Credit API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Transfer Credit API response: %w", err)
	}

	return result, nil
}