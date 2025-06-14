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

// BalanceInquiryService implements the Balance Inquiry business logic
type BalanceInquiryService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI Balance Inquiry API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewBalanceInquiryService creates a new Balance Inquiry service instance
func NewBalanceInquiryService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*BalanceInquiryService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &BalanceInquiryService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Balance Inquiry API endpoints
		endpoints: map[string]string{
			"balance-inquiry": "/api/v1.0/balance-inquiry",
		},
		methods: map[string]string{
			"balance-inquiry": "POST",
		},
	}, nil
}

// NewBalanceInquiryServiceWithEndpoints creates a new Balance Inquiry service with custom endpoints
func NewBalanceInquiryServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*BalanceInquiryService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getBalanceInquiryEndpoints(customEndpoints)
	methods := getBalanceInquiryMethods()

	return &BalanceInquiryService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getBalanceInquiryEndpoints returns Balance Inquiry endpoints (custom or default)
func getBalanceInquiryEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
		"balance-inquiry": "/api/v1.0/balance-inquiry",
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

// getBalanceInquiryMethods returns HTTP methods for Balance Inquiry endpoints
func getBalanceInquiryMethods() map[string]string {
	return map[string]string{
		"balance-inquiry": "POST",
	}
}

// BalanceInquiry performs balance inquiry
func (b *BalanceInquiryService) BalanceInquiry(
	ctx context.Context,
	payload *types.BalanceInquiryPayload,
) (*types.BalanceInquiryResponse, error) {
	result, err := b.makeRequest(ctx, "balance-inquiry", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to perform balance inquiry: %w", err)
	}

	// Convert result to BalanceInquiryResponse
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal balance inquiry response: %w", err)
	}

	var balanceResponse types.BalanceInquiryResponse
	if err := json.Unmarshal(resultBytes, &balanceResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance inquiry response: %w", err)
	}

	return &balanceResponse, nil
}

// makeRequest handles the HTTP request for ASPI Balance Inquiry API calls
func (b *BalanceInquiryService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
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

	// Get B2B2C access token (using hardcoded values from PHP for now)
	// TODO: Make this configurable or dynamic
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

	// Build request headers
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           b.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "localhost", // TODO: Make configurable
		"X-Partner-Id":           b.config.ASPI.ClientID,
		"X-External-Id":          "41807553358950093184162180797837", // TODO: Make configurable
		"X-IP-Address":           "127.0.0.1",                        // TODO: Get from request
		"X-Device-Id":            "09864ADCASA",                      // TODO: Make configurable
		"Channel-Id":             "95221",                            // TODO: Make configurable
	}

	// Log the request
	b.logger.WithFields(map[string]any{
		"service":   "BalanceInquiryService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making Balance Inquiry API request")

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
			"service":  "BalanceInquiryService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("Balance Inquiry API request failed")
		return nil, fmt.Errorf("Balance Inquiry API request failed: %w", err)
	}

	// Log successful response
	b.logger.WithFields(map[string]any{
		"service":     "BalanceInquiryService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("Balance Inquiry API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Balance Inquiry API response: %w", err)
	}

	return result, nil
}