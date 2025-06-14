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

// BCAService implements BCA-specific business logic
type BCAService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// BCA API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewBCAService creates a new BCA service instance
func NewBCAService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*BCAService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &BCAService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// BCA API endpoints
		endpoints: map[string]string{
			"create-va":        "/api/v1.0/bca/transfer-va/create-va",
			"inquiry-va":       "/api/v1.0/bca/transfer-va/inquiry-va",
			"update-va":        "/api/v1.0/bca/transfer-va/update-va",
			"payment":          "/api/v1.0/bca/transfer-va/payment",
			"balance-inquiry":  "/api/v1.0/bca/balance-inquiry",
			"account-inquiry":  "/api/v1.0/bca/account-inquiry-external",
			"trigger-transfer": "/api/v1.0/bca/trigger-transfer",
			"transfer-status":  "/api/v1.0/bca/transfer/status",
			"generate-qr":      "/api/v1.0/bca/qr/qr-mpm-generate",
			"decode-qr":        "/api/v1.0/bca/qr/qr-mpm-decode",
		},
		methods: map[string]string{
			"create-va":        "POST",
			"inquiry-va":       "POST",
			"update-va":        "PUT",
			"payment":          "POST",
			"balance-inquiry":  "POST",
			"account-inquiry":  "POST",
			"trigger-transfer": "POST",
			"transfer-status":  "POST",
			"generate-qr":      "POST",
			"decode-qr":        "POST",
		},
	}, nil
}

// NewBCAServiceWithEndpoints creates a new BCA service with custom endpoints
func NewBCAServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*BCAService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getBCAEndpoints(customEndpoints)
	methods := getBCAMethods()

	return &BCAService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getBCAEndpoints returns BCA endpoints (custom or default)
func getBCAEndpoints(customEndpoints map[string]string) map[string]string {
	// Default BCA endpoints
	defaultEndpoints := map[string]string{
		"create-va":        "/api/v1.0/bca/transfer-va/create-va",
		"inquiry-va":       "/api/v1.0/bca/transfer-va/inquiry-va",
		"update-va":        "/api/v1.0/bca/transfer-va/update-va",
		"payment":          "/api/v1.0/bca/transfer-va/payment",
		"balance-inquiry":  "/api/v1.0/bca/balance-inquiry",
		"account-inquiry":  "/api/v1.0/bca/account-inquiry-external",
		"trigger-transfer": "/api/v1.0/bca/trigger-transfer",
		"transfer-status":  "/api/v1.0/bca/transfer/status",
		"generate-qr":      "/api/v1.0/bca/qr/qr-mpm-generate",
		"decode-qr":        "/api/v1.0/bca/qr/qr-mpm-decode",
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

// getBCAMethods returns HTTP methods for BCA endpoints
func getBCAMethods() map[string]string {
	return map[string]string{
		"create-va":        "POST",
		"inquiry-va":       "POST",
		"update-va":        "PUT",
		"payment":          "POST",
		"balance-inquiry":  "POST",
		"account-inquiry":  "POST",
		"trigger-transfer": "POST",
		"transfer-status":  "POST",
		"generate-qr":      "POST",
		"decode-qr":        "POST",
	}
}

// CreateBCAVirtualAccount creates a BCA virtual account
func (b *BCAService) CreateBCAVirtualAccount(ctx context.Context, payload *types.BCACreateVAPayload) (map[string]any, error) {
	return b.makeRequest(ctx, "create-va", payload)
}

// InquiryBCAVirtualAccount inquires a BCA virtual account
func (b *BCAService) InquiryBCAVirtualAccount(ctx context.Context, payload *types.BCAInquiryVAPayload) (map[string]any, error) {
	return b.makeRequest(ctx, "inquiry-va", payload)
}

// BCABalanceInquiry performs a BCA balance inquiry
func (b *BCAService) BCABalanceInquiry(ctx context.Context, payload *types.BCABalanceInquiryPayload) (map[string]any, error) {
	return b.makeRequest(ctx, "balance-inquiry", payload)
}

// GenerateBCAQR generates a BCA QR code
func (b *BCAService) GenerateBCAQR(ctx context.Context, payload *types.BCAGenerateQRPayload) (map[string]any, error) {
	return b.makeRequest(ctx, "generate-qr", payload)
}

// BCATransfer performs a BCA transfer
func (b *BCAService) BCATransfer(ctx context.Context, payload *types.BCATransferPayload) (map[string]any, error) {
	return b.makeRequest(ctx, "trigger-transfer", payload)
}

// makeRequest handles the HTTP request for BCA API calls
func (b *BCAService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
	// Validate endpoint
	path, exists := b.endpoints[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("endpoint '%s' not found", endpoint))
	}

	method, exists := b.methods[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("method for endpoint '%s' not found", endpoint))
	}

	// Generate timestamp for request
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Get B2B access token
	b2bToken, err := b.authManager.GetAccessToken(ctx)
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

	b2b2cToken, err := b.authManager.GetCustomerAccessToken(ctx, b2b2cTokenReq)
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

	signatureValue, err := b.symSigner.Sign(sigPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Build request headers with proper values for production use
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           b.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "https://api.yourdomain.com", // Should be configured based on your domain
		"X-Partner-Id":           b.config.ASPI.ClientID,
		"X-External-Id":          generateExternalId(), // Should generate a unique ID for each request
		"X-IP-Address":           getClientIP(),        // Should get from client request
		"X-Device-Id":            getDeviceId(),        // Should get from client request
		"Channel-Id":             getChannelId(),       // Should be configured based on your channel
		// BCA-specific headers
		"X-BCA-Key":              b.config.ASPI.ClientID,
		"X-BCA-Timestamp":        timestamp,
	}

	// Log the request
	b.logger.WithFields(map[string]any{
		"service":   "BCAService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making BCA API request")

	// Make the HTTP request
	var response *contracts.Response

	url := b.config.ASPI.BaseURL + path

	switch method {
	case "POST":
		response, err = b.httpClient.Post(ctx, url, payload, requestHeaders)
	case "PUT":
		response, err = b.httpClient.Put(ctx, url, payload, requestHeaders)
	case "DELETE":
		response, err = b.httpClient.Delete(ctx, url, requestHeaders)
	case "GET":
		response, err = b.httpClient.Get(ctx, url, requestHeaders)
	default:
		return nil, errors.NewValidationError(fmt.Sprintf("unsupported HTTP method: %s", method))
	}

	if err != nil {
		b.logger.WithFields(map[string]any{
			"service":  "BCAService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("BCA API request failed")
		return nil, fmt.Errorf("BCA API request failed: %w", err)
	}

	// Log successful response
	b.logger.WithFields(map[string]any{
		"service":     "BCAService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("BCA API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse BCA API response: %w", err)
	}

	return result, nil
}