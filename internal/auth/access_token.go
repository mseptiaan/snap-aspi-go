package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/signature"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// TokenResponse represents the response from access token endpoints
type TokenResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccessToken     string `json:"accessToken"`
	TokenType       string `json:"tokenType"`
	ExpiresIn       string `json:"expiresIn"`
}

// CustomerTokenRequest represents the request for B2B2C token
type CustomerTokenRequest struct {
	AuthCode       string                 `json:"authCode"`
	RefreshToken   string                 `json:"refreshToken"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
}

// AccessTokenManager handles access token operations
type AccessTokenManager struct {
	config           *config.Config
	httpClient       contracts.HTTPClient
	asymmetricSigner *signature.AsymmetricSigner
	logger           logging.Logger

	// Endpoints
	b2bEndpoint   string
	b2b2cEndpoint string
}

// NewAccessTokenManager creates a new access token manager
func NewAccessTokenManager(
	cfg *config.Config,
	httpClient contracts.HTTPClient,
	logger logging.Logger,
) (*AccessTokenManager, error) {
	// Load private key for asymmetric signing
	privateKeyPEM, err := loadPrivateKey(cfg.ASPI.PrivateKeyPath)
	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to load private key: %v", err))
	}

	// Create asymmetric signer
	asymmetricSigner, err := signature.NewAsymmetricSignerFromPrivateKey(privateKeyPEM)
	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to create asymmetric signer: %v", err))
	}

	return &AccessTokenManager{
		config:           cfg,
		httpClient:       httpClient,
		asymmetricSigner: asymmetricSigner,
		logger:           logger,
		b2bEndpoint:      cfg.ASPI.Endpoints.B2BToken,
		b2b2cEndpoint:    cfg.ASPI.Endpoints.B2B2CToken,
	}, nil
}

// GetAccessToken retrieves a B2B access token
func (m *AccessTokenManager) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	// Create headers with signature
	headers, err := m.createAuthHeaders()
	if err != nil {
		return nil, errors.NewInternalError(err, "failed to create auth headers")
	}

	// Create request body
	requestBody := map[string]string{
		"grantType": "client_credentials",
	}

	// Log the request
	m.logger.WithFields(map[string]interface{}{
		"endpoint":   m.b2bEndpoint,
		"grant_type": "client_credentials",
	}).Info("Requesting B2B access token")

	// Make the request
	response, err := m.httpClient.Post(ctx, m.b2bEndpoint, requestBody, headers)
	if err != nil {
		return nil, errors.NewExternalError(err, "failed to request access token")
	}

	// Check if request was successful
	if !response.Success {
		return nil, errors.NewExternalError(
			fmt.Errorf("API request failed with status %d", response.StatusCode),
			"access token request failed",
		)
	}

	// Parse response
	var tokenResponse TokenResponse
	if err := json.Unmarshal(response.Body, &tokenResponse); err != nil {
		return nil, errors.NewInternalError(err, "failed to parse token response")
	}

	// Log successful token acquisition
	m.logger.WithFields(map[string]interface{}{
		"response_code": tokenResponse.ResponseCode,
		"token_type":    tokenResponse.TokenType,
		"expires_in":    tokenResponse.ExpiresIn,
	}).Info("Successfully obtained B2B access token")

	return &tokenResponse, nil
}

// GetCustomerAccessToken retrieves a B2B2C access token
func (m *AccessTokenManager) GetCustomerAccessToken(
	ctx context.Context,
	request CustomerTokenRequest,
) (*TokenResponse, error) {
	// Validate request
	if err := m.validateCustomerTokenRequest(request); err != nil {
		return nil, err
	}

	// Create headers with signature
	headers, err := m.createAuthHeaders()
	if err != nil {
		return nil, errors.NewInternalError(err, "failed to create auth headers")
	}

	// Create request body
	requestBody := map[string]interface{}{
		"grantType":      "authorization_code",
		"authCode":       request.AuthCode,
		"refreshToken":   request.RefreshToken,
		"additionalInfo": request.AdditionalInfo,
	}

	// Log the request
	m.logger.WithFields(map[string]interface{}{
		"endpoint":   m.b2b2cEndpoint,
		"grant_type": "authorization_code",
	}).Info("Requesting B2B2C access token")

	// Make the request
	response, err := m.httpClient.Post(ctx, m.b2b2cEndpoint, requestBody, headers)
	if err != nil {
		return nil, errors.NewExternalError(err, "failed to request customer access token")
	}

	// Check if request was successful
	if !response.Success {
		return nil, errors.NewExternalError(
			fmt.Errorf("API request failed with status %d", response.StatusCode),
			"customer access token request failed",
		)
	}

	// Parse response
	var tokenResponse TokenResponse
	if err := json.Unmarshal(response.Body, &tokenResponse); err != nil {
		return nil, errors.NewInternalError(err, "failed to parse customer token response")
	}

	// Log successful token acquisition
	m.logger.WithFields(map[string]interface{}{
		"response_code": tokenResponse.ResponseCode,
		"token_type":    tokenResponse.TokenType,
		"expires_in":    tokenResponse.ExpiresIn,
	}).Info("Successfully obtained B2B2C access token")

	return &tokenResponse, nil
}

// createAuthHeaders creates headers with asymmetric signature for authentication
func (m *AccessTokenManager) createAuthHeaders() (map[string]string, error) {
	// Create timestamp
	timestamp := time.Now()

	// Create asymmetric payload
	payload := signature.AsymmetricPayload{
		ClientID:  m.config.ASPI.ClientID,
		Timestamp: timestamp,
	}

	// Sign the payload
	signatureValue, err := m.asymmetricSigner.Sign(payload)
	if err != nil {
		return nil, errors.NewInternalError(err, "failed to create signature")
	}

	// Create headers
	headers := map[string]string{
		"X-TIMESTAMP":  timestamp.Format(time.RFC3339),
		"X-CLIENT-KEY": m.config.ASPI.ClientID,
		"X-SIGNATURE":  signatureValue,
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	return headers, nil
}

// validateCustomerTokenRequest validates the customer token request
func (m *AccessTokenManager) validateCustomerTokenRequest(request CustomerTokenRequest) error {
	if request.AuthCode == "" {
		return errors.NewValidationError("authCode is required")
	}
	if request.RefreshToken == "" {
		return errors.NewValidationError("refreshToken is required")
	}
	if request.AdditionalInfo == nil {
		return errors.NewValidationError("additionalInfo is required")
	}
	return nil
}

// loadPrivateKey loads the private key from file
func loadPrivateKey(keyPath string) ([]byte, error) {
	if keyPath == "" {
		return nil, fmt.Errorf("private key path not configured")
	}

	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %v", err)
	}

	return keyData, nil
}
