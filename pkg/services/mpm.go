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
		// ASPI MPM API endpoints - Updated to match ASPI documentation
		endpoints: map[string]string{
			"generate-qr":     "/api/v1.0/qr/qr-mpm-generate",
			"decode-qr":       "/api/v1.0/qr/qr-mpm-decode",
			"apply-ott":       "/api/v1.0/qr/apply-ott",
			"payment":         "/api/v1.0/qr/qr-mpm-payment",
			"query-payment":   "/api/v1.0/qr/qr-mpm-query",
			"notify-qr":       "/api/v1.0/qr/qr-mpm-notify",
			"cancel-payment":  "/api/v1.0/qr/qr-mpm-cancel",
			"refund":          "/api/v1.0/qr/qr-mpm-refund",
			"transfer":        "/api/v1.0/transfer-kredit/mpm",
			"inquiry":         "/api/v1.0/transfer-kredit/mpm/inquiry",
			"status":          "/api/v1.0/qr/qr-mpm-status",
			"balance-inquiry": "/api/v1.0/transfer-kredit/mpm/balance-inquiry",
			"account-inquiry": "/api/v1.0/transfer-kredit/mpm/account-inquiry",
			"history":         "/api/v1.0/transfer-kredit/mpm/history",
		},
		methods: map[string]string{
			"generate-qr":     "POST",
			"decode-qr":       "POST",
			"apply-ott":       "POST",
			"payment":         "POST",
			"query-payment":   "POST",
			"notify-qr":       "POST",
			"cancel-payment":  "POST",
			"refund":          "POST",
			"transfer":        "POST",
			"inquiry":         "POST",
			"status":          "POST",
			"balance-inquiry": "POST",
			"account-inquiry": "POST",
			"history":         "POST",
		},
	}, nil
}

// NewMPMServiceWithEndpoints creates a new MPM service with custom endpoints
func NewMPMServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*MPMService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getMPMEndpoints(customEndpoints)
	methods := getMPMMethods()

	return &MPMService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getMPMEndpoints returns MPM endpoints (custom or default)
func getMPMEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints - Updated to match ASPI documentation
	defaultEndpoints := map[string]string{
		"generate-qr":     "/api/v1.0/qr/qr-mpm-generate",
		"decode-qr":       "/api/v1.0/qr/qr-mpm-decode",
		"apply-ott":       "/api/v1.0/qr/apply-ott",
		"payment":         "/api/v1.0/qr/qr-mpm-payment",
		"query-payment":   "/api/v1.0/qr/qr-mpm-query",
		"notify-qr":       "/api/v1.0/qr/qr-mpm-notify",
		"cancel-payment":  "/api/v1.0/qr/qr-mpm-cancel",
		"refund":          "/api/v1.0/qr/qr-mpm-refund",
		"transfer":        "/api/v1.0/transfer-kredit/mpm",
		"inquiry":         "/api/v1.0/transfer-kredit/mpm/inquiry",
		"status":          "/api/v1.0/qr/qr-mpm-status",
		"balance-inquiry": "/api/v1.0/transfer-kredit/mpm/balance-inquiry",
		"account-inquiry": "/api/v1.0/transfer-kredit/mpm/account-inquiry",
		"history":         "/api/v1.0/transfer-kredit/mpm/history",
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

// getMPMMethods returns HTTP methods for MPM endpoints
func getMPMMethods() map[string]string {
	return map[string]string{
		"generate-qr":     "POST",
		"decode-qr":       "POST",
		"apply-ott":       "POST",
		"payment":         "POST",
		"query-payment":   "POST",
		"notify-qr":       "POST",
		"cancel-payment":  "POST",
		"refund":          "POST",
		"transfer":        "POST",
		"inquiry":         "POST",
		"status":          "POST",
		"balance-inquiry": "POST",
		"account-inquiry": "POST",
		"history":         "POST",
	}
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

// DecodeQR decodes a QR code for MPM payment
func (m *MPMService) DecodeQR(
	ctx context.Context,
	payload *types.MPMDecodeQRPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "decode-qr", payload)
}

// ApplyOTT applies a one-time token for payment redirect
func (m *MPMService) ApplyOTT(
	ctx context.Context,
	payload *types.MPMApplyOTTPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "apply-ott", payload)
}

// Payment processes a payment for MPM
func (m *MPMService) Payment(
	ctx context.Context,
	payload *types.MPMPaymentPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "payment", payload)
}

// QueryPayment queries payment status for MPM
func (m *MPMService) QueryPayment(
	ctx context.Context,
	payload *types.MPMQueryPaymentPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "query-payment", payload)
}

// CancelPayment cancels a payment for MPM
func (m *MPMService) CancelPayment(
	ctx context.Context,
	payload *types.MPMCancelPaymentPayload,
) (map[string]any, error) {
	return m.makeRequest(ctx, "cancel-payment", payload)
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

	// Get B2B2C access token
	// In a production environment, this should be obtained from the user's session
	// or passed in as a parameter rather than using hardcoded values
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

	// Build request headers with proper values for production use
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           m.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "https://api.yourdomain.com", // Should be configured based on your domain
		"X-Partner-Id":           m.config.ASPI.ClientID,
		"X-External-Id":          generateExternalId(), // Should generate a unique ID for each request
		"X-IP-Address":           getClientIP(),        // Should get from client request
		"X-Device-Id":            getDeviceId(),        // Should get from client request
		"Channel-Id":             getChannelId(),       // Should be configured based on your channel
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