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

// RegistrationService implements the Registration business logic
type RegistrationService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI Registration API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewRegistrationService creates a new Registration service instance
func NewRegistrationService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*RegistrationService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &RegistrationService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Registration API endpoints
		endpoints: map[string]string{
			"register":      "/api/v1.0/registration/register",
			"register-card": "/api/v1.0/registration/card-registration",
			"bind-account":  "/api/v1.0/registration/account-binding",
			"verify-otp":    "/api/v1.0/registration/otp-verification",
		},
		methods: map[string]string{
			"register":      "POST",
			"register-card": "POST",
			"bind-account":  "POST",
			"verify-otp":    "POST",
		},
	}, nil
}

// NewRegistrationServiceWithEndpoints creates a new Registration service with custom endpoints
func NewRegistrationServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*RegistrationService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getRegistrationEndpoints(customEndpoints)
	methods := getRegistrationMethods()

	return &RegistrationService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getRegistrationEndpoints returns Registration endpoints (custom or default)
func getRegistrationEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
		"register":      "/api/v1.0/registration/register",
		"register-card": "/api/v1.0/registration/card-registration",
		"bind-account":  "/api/v1.0/registration/account-binding",
		"verify-otp":    "/api/v1.0/registration/otp-verification",
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

// getRegistrationMethods returns HTTP methods for Registration endpoints
func getRegistrationMethods() map[string]string {
	return map[string]string{
		"register":      "POST",
		"register-card": "POST",
		"bind-account":  "POST",
		"verify-otp":    "POST",
	}
}

// Register performs user registration
func (r *RegistrationService) Register(ctx context.Context, payload *types.RegistrationPayload) (map[string]any, error) {
	return r.makeRequest(ctx, "register", payload)
}

// RegisterCard performs card registration
func (r *RegistrationService) RegisterCard(ctx context.Context, payload *types.CardRegistrationPayload) (map[string]any, error) {
	return r.makeRequest(ctx, "register-card", payload)
}

// BindAccount performs account binding
func (r *RegistrationService) BindAccount(ctx context.Context, payload *types.AccountBindingPayload) (map[string]any, error) {
	return r.makeRequest(ctx, "bind-account", payload)
}

// VerifyOTP performs OTP verification
func (r *RegistrationService) VerifyOTP(ctx context.Context, payload *types.OTPVerificationPayload) (map[string]any, error) {
	return r.makeRequest(ctx, "verify-otp", payload)
}

// makeRequest handles the HTTP request for ASPI Registration API calls
func (r *RegistrationService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
	// Validate endpoint
	path, exists := r.endpoints[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("endpoint '%s' not found", endpoint))
	}

	method, exists := r.methods[endpoint]
	if !exists {
		return nil, errors.NewValidationError(fmt.Sprintf("method for endpoint '%s' not found", endpoint))
	}

	// Generate timestamp for request
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Get B2B access token
	b2bToken, err := r.authManager.GetAccessToken(ctx)
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

	b2b2cToken, err := r.authManager.GetCustomerAccessToken(ctx, b2b2cTokenReq)
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

	signatureValue, err := r.symSigner.Sign(sigPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signature: %w", err)
	}

	// Build request headers with proper values for production use
	requestHeaders := map[string]string{
		"Content-Type":           "application/json",
		"X-Client-Key":           r.config.ASPI.ClientID,
		"X-Timestamp":            timestamp,
		"Authorization":          "Bearer " + b2bToken.AccessToken,
		"Authorization-Customer": "Bearer " + b2b2cToken.AccessToken,
		"X-Signature":            signatureValue,
		"X-Origin":               "https://api.yourdomain.com", // Should be configured based on your domain
		"X-Partner-Id":           r.config.ASPI.ClientID,
		"X-External-Id":          generateExternalId(), // Should generate a unique ID for each request
		"X-IP-Address":           getClientIP(),        // Should get from client request
		"X-Device-Id":            getDeviceId(),        // Should get from client request
		"Channel-Id":             getChannelId(),       // Should be configured based on your channel
	}

	// Log the request
	r.logger.WithFields(map[string]any{
		"service":   "RegistrationService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making Registration API request")

	// Make the HTTP request
	var response *contracts.Response

	url := r.config.ASPI.BaseURL + path

	switch method {
	case "POST":
		response, err = r.httpClient.Post(ctx, url, payload, requestHeaders)
	case "PUT":
		response, err = r.httpClient.Put(ctx, url, payload, requestHeaders)
	case "DELETE":
		response, err = r.httpClient.Delete(ctx, url, requestHeaders)
	case "GET":
		response, err = r.httpClient.Get(ctx, url, requestHeaders)
	default:
		return nil, errors.NewValidationError(fmt.Sprintf("unsupported HTTP method: %s", method))
	}

	if err != nil {
		r.logger.WithFields(map[string]any{
			"service":  "RegistrationService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("Registration API request failed")
		return nil, fmt.Errorf("Registration API request failed: %w", err)
	}

	// Log successful response
	r.logger.WithFields(map[string]any{
		"service":     "RegistrationService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("Registration API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Registration API response: %w", err)
	}

	return result, nil
}