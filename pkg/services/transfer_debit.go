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

// TransferDebitService implements the Transfer Debit business logic
type TransferDebitService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI Transfer Debit API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewTransferDebitService creates a new Transfer Debit service instance
func NewTransferDebitService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*TransferDebitService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &TransferDebitService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Transfer Debit API endpoints
		endpoints: map[string]string{
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
		},
		methods: map[string]string{
			"direct-debit-payment":       "POST",
			"direct-debit-status":        "POST",
			"direct-debit-cancel":        "POST",
			"direct-debit-refund":        "POST",
			"cpm-generate-qr":            "POST",
			"cpm-payment":                "POST",
			"auth-payment":               "POST",
			"auth-capture":               "POST",
			"auth-void":                  "POST",
			"direct-debit-bifast":        "POST",
			"direct-debit-bifast-payment": "POST",
		},
	}, nil
}

// NewTransferDebitServiceWithEndpoints creates a new Transfer Debit service with custom endpoints
func NewTransferDebitServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*TransferDebitService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getTransferDebitEndpoints(customEndpoints)
	methods := getTransferDebitMethods()

	return &TransferDebitService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getTransferDebitEndpoints returns Transfer Debit endpoints (custom or default)
func getTransferDebitEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
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

// getTransferDebitMethods returns HTTP methods for Transfer Debit endpoints
func getTransferDebitMethods() map[string]string {
	return map[string]string{
		"direct-debit-payment":       "POST",
		"direct-debit-status":        "POST",
		"direct-debit-cancel":        "POST",
		"direct-debit-refund":        "POST",
		"cpm-generate-qr":            "POST",
		"cpm-payment":                "POST",
		"auth-payment":               "POST",
		"auth-capture":               "POST",
		"auth-void":                  "POST",
		"direct-debit-bifast":        "POST",
		"direct-debit-bifast-payment": "POST",
	}
}

// DirectDebitPayment performs direct debit payment
func (t *TransferDebitService) DirectDebitPayment(ctx context.Context, payload *types.DirectDebitPaymentPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-payment", payload)
}

// DirectDebitStatus checks direct debit status
func (t *TransferDebitService) DirectDebitStatus(ctx context.Context, payload *types.DirectDebitStatusPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-status", payload)
}

// DirectDebitCancel cancels direct debit
func (t *TransferDebitService) DirectDebitCancel(ctx context.Context, payload *types.DirectDebitCancelPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-cancel", payload)
}

// DirectDebitRefund refunds direct debit
func (t *TransferDebitService) DirectDebitRefund(ctx context.Context, payload *types.DirectDebitRefundPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-refund", payload)
}

// CPMGenerateQR generates QR for CPM
func (t *TransferDebitService) CPMGenerateQR(ctx context.Context, payload *types.CPMGenerateQRPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "cpm-generate-qr", payload)
}

// CPMPayment performs CPM payment
func (t *TransferDebitService) CPMPayment(ctx context.Context, payload *types.CPMPaymentPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "cpm-payment", payload)
}

// AuthPayment performs auth payment
func (t *TransferDebitService) AuthPayment(ctx context.Context, payload *types.AuthPaymentPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "auth-payment", payload)
}

// AuthCapture performs auth capture
func (t *TransferDebitService) AuthCapture(ctx context.Context, payload *types.AuthCapturePayload) (map[string]any, error) {
	return t.makeRequest(ctx, "auth-capture", payload)
}

// AuthVoid performs auth void
func (t *TransferDebitService) AuthVoid(ctx context.Context, payload *types.AuthVoidPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "auth-void", payload)
}

// DirectDebitBIFAST performs direct debit BI-FAST
func (t *TransferDebitService) DirectDebitBIFAST(ctx context.Context, payload *types.DirectDebitBIFASTPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-bifast", payload)
}

// DirectDebitBIFASTPayment performs direct debit BI-FAST payment
func (t *TransferDebitService) DirectDebitBIFASTPayment(ctx context.Context, payload *types.DirectDebitBIFASTPaymentPayload) (map[string]any, error) {
	return t.makeRequest(ctx, "direct-debit-bifast-payment", payload)
}

// makeRequest handles the HTTP request for ASPI Transfer Debit API calls
func (t *TransferDebitService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
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
		"service":   "TransferDebitService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making Transfer Debit API request")

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
			"service":  "TransferDebitService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("Transfer Debit API request failed")
		return nil, fmt.Errorf("Transfer Debit API request failed: %w", err)
	}

	// Log successful response
	t.logger.WithFields(map[string]any{
		"service":     "TransferDebitService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("Transfer Debit API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Transfer Debit API response: %w", err)
	}

	return result, nil
}