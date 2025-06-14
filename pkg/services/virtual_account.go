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

// VirtualAccountService implements the VirtualAccount business logic
// Converted from PHP VirtualAccount service class
type VirtualAccountService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewVirtualAccountService creates a new VirtualAccount service instance
func NewVirtualAccountService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*VirtualAccountService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &VirtualAccountService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Virtual Account API endpoints
		endpoints: map[string]string{
			"inquiry":       "/api/v1.0/transfer-va/inquiry",
			"inquiry-va":    "/api/v1.0/transfer-va/inquiry-va",
			"create-va":     "/api/v1.0/transfer-va/create-va",
			"update-va":     "/api/v1.0/transfer-va/update-va",
			"delete-va":     "/api/v1.0/transfer-va/delete-va",
			"payment":       "/api/v1.0/transfer-va/payment",
			"status":        "/api/v1.0/transfer-va/status",
			"report":        "/api/v1.0/transfer-va/report",
			"update-status": "/api/v1.0/transfer-va/update-status",
		},
		methods: map[string]string{
			"inquiry":       "POST",
			"inquiry-va":    "POST",
			"create-va":     "POST",
			"update-va":     "PUT",
			"delete-va":     "DELETE",
			"payment":       "POST",
			"status":        "POST",
			"report":        "POST",
			"update-status": "PUT",
		},
	}, nil
}

// CreateVA creates a new Virtual Account
func (v *VirtualAccountService) CreateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "create-va", payload)
}

// UpdateVA updates an existing Virtual Account
func (v *VirtualAccountService) UpdateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "update-va", payload)
}

// DeleteVA deletes a Virtual Account
func (v *VirtualAccountService) DeleteVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "delete-va", payload)
}

// InquiryVA performs Virtual Account inquiry
func (v *VirtualAccountService) InquiryVA(
	ctx context.Context,
	payload *types.InquiryVAPayload,
) (map[string]any, error) {
	return v.makeRequest(ctx, "inquiry-va", payload)
}

// Inquiry performs general inquiry
func (v *VirtualAccountService) Inquiry(ctx context.Context, payload *types.InquiryPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "inquiry", payload)
}

// Payment processes a Virtual Account payment
func (v *VirtualAccountService) Payment(ctx context.Context, payload *types.PaymentPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "payment", payload)
}

// Status checks transaction status
func (v *VirtualAccountService) Status(ctx context.Context, payload *types.StatusPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "status", payload)
}

// Report generates transaction reports
func (v *VirtualAccountService) Report(ctx context.Context, payload *types.ReportPayload) (map[string]any, error) {
	return v.makeRequest(ctx, "report", payload)
}

// UpdateStatus updates transaction status
func (v *VirtualAccountService) UpdateStatus(
	ctx context.Context,
	payload *types.StatusPayload,
) (map[string]any, error) {
	return v.makeRequest(ctx, "update-status", payload)
}

// makeRequest handles the HTTP request for ASPI API calls
// Equivalent to PHP VirtualAccount::send() method
func (v *VirtualAccountService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
	// Validate endpoint
	path, exists := v.endpoints[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("endpoint '%s' not found", endpoint))
	}

	method, exists := v.methods[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("method for endpoint '%s' not found", endpoint))
	}

	// Generate timestamp for request
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Get B2B access token
	b2bToken, err := v.authManager.GetAccessToken(ctx)
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

	b2b2cToken, err := v.authManager.GetCustomerAccessToken(ctx, b2b2cTokenReq)
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

	signatureValue, err := v.symSigner.Sign(sigPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Build request headers (equivalent to PHP headers() method)
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           v.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "localhost", // TODO: Make configurable
		"X-Partner-Id":           v.config.ASPI.ClientID,
		"X-External-Id":          "41807553358950093184162180797837", // TODO: Make configurable
		"X-IP-Address":           "127.0.0.1",                        // TODO: Get from request
		"X-Device-Id":            "09864ADCASA",                      // TODO: Make configurable
		"Channel-Id":             "95221",                            // TODO: Make configurable
	}

	// Log the request
	v.logger.WithFields(map[string]any{
		"endpoint": endpoint,
		"path":     path,
		"method":   method,
	}).Info("Making ASPI API request")

	// Make the HTTP request using the appropriate method
	var response *contracts.Response
	switch method {
	case "GET":
		response, err = v.httpClient.Get(ctx, v.config.ASPI.BaseURL+path, requestHeaders)
	case "POST":
		response, err = v.httpClient.Post(ctx, v.config.ASPI.BaseURL+path, payload, requestHeaders)
	case "PUT":
		response, err = v.httpClient.Put(ctx, v.config.ASPI.BaseURL+path, payload, requestHeaders)
	case "DELETE":
		response, err = v.httpClient.Delete(ctx, v.config.ASPI.BaseURL+path, requestHeaders)
	default:
		return nil, errors.NewValidationError(fmt.Sprintf("unsupported HTTP method: %s", method))
	}

	if err != nil {
		v.logger.WithError(err).Error("ASPI API request failed")
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	// Handle response based on status code
	if response.IsSuccessful() {
		var result map[string]any
		if err := json.Unmarshal(response.Body, &result); err != nil {
			return nil, errors.NewInternalError(err, "Failed to parse response")
		}

		v.logger.WithFields(map[string]any{
			"endpoint":    endpoint,
			"status_code": response.StatusCode,
		}).Info("ASPI API request successful")

		return result, nil
	}

	// Handle error responses
	v.logger.WithFields(map[string]any{
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
		"response":    string(response.Body),
	}).Error("ASPI API request failed")

	return nil, errors.NewExternalError(
		fmt.Errorf("ASPI API error (status %d): %s", response.StatusCode, string(response.Body)),
		"ASPI API request failed",
	)
}
