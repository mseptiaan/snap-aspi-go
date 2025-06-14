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

// MPMService implements the MPM Transfer Credit business logic
// Based on ASPI Transfer Credit MPM API documentation
type MPMService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI MPM API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewMPMService creates a new MPM Transfer Credit service instance
func NewMPMService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*MPMService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &MPMService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI MPM Transfer Credit API endpoints
		endpoints: map[string]string{
			"transfer":        "/api/v1.0/transfer-kredit/mpm",
			"inquiry":         "/api/v1.0/transfer-kredit/mpm/inquiry",
			"status":          "/api/v1.0/transfer-kredit/mpm/status",
			"refund":          "/api/v1.0/transfer-kredit/mpm/refund",
			"balance-inquiry": "/api/v1.0/transfer-kredit/mpm/balance-inquiry",
			"account-inquiry": "/api/v1.0/transfer-kredit/mpm/account-inquiry",
			"history":         "/api/v1.0/transfer-kredit/mpm/history",
			"generate-qr":     "/api/v1.0/qr/qr-mpm-generate",
			"notify-qr":       "/api/v1.0/qr/qr-mpm-notify",
		},
		methods: map[string]string{
			"transfer":        "POST",
			"inquiry":         "POST",
			"status":          "POST",
			"refund":          "POST",
			"balance-inquiry": "POST",
			"account-inquiry": "POST",
			"history":         "POST",
			"generate-qr":     "POST",
			"notify-qr":       "POST",
		},
	}, nil
}

// Transfer executes an MPM transfer credit transaction
func (m *MPMService) Transfer(ctx context.Context, payload *types.MPMTransferPayload) (map[string]any, error) {
	return m.makeRequest(ctx, "transfer", payload)
}

// Inquiry performs MPM transfer inquiry
func (m *MPMService) Inquiry(ctx context.Context, payload *types.MPMInquiryPayload) (map[string]any, error) {
	return m.makeRequest(ctx, "inquiry", payload)
}

// Status checks MPM transfer status
func (m *MPMService) Status(ctx context.Context, payload *types.MPMStatusPayload) (map[string]any, error) {
	return m.makeRequest(ctx, "status", payload)
}

// Refund processes MPM transfer refund
func (m *MPMService) Refund(ctx context.Context, payload *types.MPMRefundPayload) (map[string]any, error) {
	return m.makeRequest(ctx, "refund", payload)
}

// BalanceInquiry performs MPM balance inquiry
func (m *MPMService) BalanceInquiry(
	ctx context.Context,
	payload *types.MPMBalanceInquiryPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "balance-inquiry", payload)
}

// AccountInquiry performs MPM account inquiry
func (m *MPMService) AccountInquiry(
	ctx context.Context,
	payload *types.MPMAccountInquiryPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "account-inquiry", payload)
}

// History retrieves MPM transaction history
func (m *MPMService) History(ctx context.Context, payload *types.MPMHistoryPayload) (map[string]any, error) {
	return m.makeRequest(ctx, "history", payload)
}

// GenerateQR generates QR code for MPM payment
func (m *MPMService) GenerateQR(
	ctx context.Context,
	payload *types.MPMGenerateQRPayload,
) (*types.MPMQRResponse, error) {
	result, err := m.makeRequest(ctx, "generate-qr", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR: %w", err)
	}

	// Convert result to MPMQRResponse
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal QR response: %w", err)
	}

	var qrResponse types.MPMQRResponse
	if err := json.Unmarshal(resultBytes, &qrResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal QR response: %w", err)
	}

	return &qrResponse, nil
}

// NotifyQR handles QR MPM payment notification
func (m *MPMService) NotifyQR(
	ctx context.Context,
	payload *types.MPMNotifyQRPayload,
) (*types.MPMNotifyResponse, error) {
	result, err := m.makeRequest(ctx, "notify-qr", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to process QR notification: %w", err)
	}

	// Convert result to MPMNotifyResponse
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal notify response: %w", err)
	}

	var notifyResponse types.MPMNotifyResponse
	if err := json.Unmarshal(resultBytes, &notifyResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal notify response: %w", err)
	}

	return &notifyResponse, nil
}

// makeRequest handles the HTTP request for ASPI MPM API calls
// Following the same pattern as VirtualAccountService
func (m *MPMService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
	// Validate endpoint
	path, exists := m.endpoints[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("endpoint '%s' not found", endpoint))
	}

	method, exists := m.methods[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("method for endpoint '%s' not found", endpoint))
	}

	// Generate timestamp for request
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Get B2B access token
	b2bToken, err := m.authManager.GetAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get B2B access token: %w", err)
	}

	// Get B2B2C access token (using hardcoded values from PHP for now)
	// TODO: Make this configurable or dynamic
	b2b2cTokenReq := auth.CustomerTokenRequest{
		AuthCode:       "a6975f82-d00a-4ddc-9633-087fefb6275e",
		RefreshToken:   "83a58570-6795-11ec-90d6-0242ac120003",
		AdditionalInfo: map[string]interface{}{},
	}

	b2b2cToken, err := m.authManager.GetCustomerAccessToken(ctx, b2b2cTokenReq)
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

	signatureValue, err := m.symSigner.Sign(sigPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Build request headers
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           m.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "localhost", // TODO: Make configurable
		"X-Partner-Id":           m.config.ASPI.ClientID,
		"X-External-Id":          "41807553358950093184162180797837", // TODO: Make configurable
		"X-IP-Address":           "127.0.0.1",                        // TODO: Get from request
		"X-Device-Id":            "09864ADCASA",                      // TODO: Make configurable
		"Channel-Id":             "95221",                            // TODO: Make configurable
	}

	// Log the request
	m.logger.WithFields(map[string]any{
		"service":   "MPMService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making MPM API request")

	// Make the HTTP request
	var response *contracts.Response

	url := m.config.ASPI.BaseURL + path

	switch method {
	case "POST":
		response, err = m.httpClient.Post(ctx, url, payload, requestHeaders)
	case "PUT":
		response, err = m.httpClient.Put(ctx, url, payload, requestHeaders)
	case "DELETE":
		response, err = m.httpClient.Delete(ctx, url, requestHeaders)
	case "GET":
		response, err = m.httpClient.Get(ctx, url, requestHeaders)
	default:
		return nil, errors.NewValidationError(fmt.Sprintf("unsupported HTTP method: %s", method))
	}

	if err != nil {
		m.logger.WithFields(map[string]any{
			"service":  "MPMService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("MPM API request failed")
		return nil, fmt.Errorf("MPM API request failed: %w", err)
	}

	// Log successful response
	m.logger.WithFields(map[string]any{
		"service":     "MPMService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("MPM API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse MPM API response: %w", err)
	}

	return result, nil
}
