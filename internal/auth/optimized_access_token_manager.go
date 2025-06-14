package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/cache"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/metrics"
	"github.com/mseptiaan/snap-aspi-go/internal/signature"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// OptimizedAccessTokenManager handles access token operations with advanced optimizations
type OptimizedAccessTokenManager struct {
	config           *config.Config
	httpClient       contracts.HTTPClient
	asymmetricSigner *signature.AsymmetricSigner
	logger           logging.Logger
	tokenCache       *cache.TokenCache
	metrics          *metrics.Metrics

	// Endpoints
	b2bEndpoint   string
	b2b2cEndpoint string

	// Advanced concurrency control
	b2bMutex     sync.Mutex
	b2b2cMutex   sync.Mutex
	requestGroup sync.Map // For deduplicating concurrent requests

	// Token refresh management
	refreshThreshold time.Duration
	refreshMutex     sync.RWMutex
	refreshTimers    map[string]*time.Timer
}

// NewOptimizedAccessTokenManager creates a new optimized access token manager
func NewOptimizedAccessTokenManager(
	cfg *config.Config,
	httpClient contracts.HTTPClient,
	logger logging.Logger,
	m *metrics.Metrics,
) (*OptimizedAccessTokenManager, error) {
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

	manager := &OptimizedAccessTokenManager{
		config:           cfg,
		httpClient:       httpClient,
		asymmetricSigner: asymmetricSigner,
		logger:           logger,
		tokenCache:       cache.NewTokenCache(),
		metrics:          m,
		b2bEndpoint:      cfg.ASPI.Endpoints.B2BToken,
		b2b2cEndpoint:    cfg.ASPI.Endpoints.B2B2CToken,
		refreshThreshold: 5 * time.Minute, // Refresh tokens 5 minutes before expiry
		refreshTimers:    make(map[string]*time.Timer),
	}

	// Start background token refresh routine
	go manager.startTokenRefreshRoutine()

	return manager, nil
}

// GetAccessToken retrieves a B2B access token with advanced optimizations
func (m *OptimizedAccessTokenManager) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	cacheKey := "b2b_token"

	// Try to get from cache first
	if token, found := m.tokenCache.Get(cacheKey); found {
		m.logger.Debug("B2B token retrieved from cache")
		m.metrics.RecordTokenRequest(true)
		return token, nil
	}

	m.metrics.RecordTokenRequest(false)

	// Use request deduplication to prevent multiple concurrent requests for the same token
	if result, loaded := m.requestGroup.LoadOrStore(cacheKey, make(chan *TokenResponse, 1)); loaded {
		// Another goroutine is already fetching this token, wait for it
		resultChan := result.(chan *TokenResponse)
		select {
		case token := <-resultChan:
			if token != nil {
				m.logger.Debug("B2B token retrieved from concurrent request")
				return token, nil
			}
		case <-ctx.Done():
			return nil, errors.NewNetworkError(ctx.Err(), "request context canceled while waiting for token")
		case <-time.After(30 * time.Second):
			// Timeout waiting for concurrent request
			m.requestGroup.Delete(cacheKey)
		}
	}

	// Ensure we clean up the request group entry
	defer m.requestGroup.Delete(cacheKey)

	// Prevent concurrent requests for the same token
	m.b2bMutex.Lock()
	defer m.b2bMutex.Unlock()

	// Double-check cache after acquiring lock
	if token, found := m.tokenCache.Get(cacheKey); found {
		m.logger.Debug("B2B token retrieved from cache after lock")
		return token, nil
	}

	// Fetch new token
	token, err := m.fetchB2BToken(ctx)
	if err != nil {
		return nil, err
	}

	// Cache the token with automatic refresh
	if expiresIn, err := strconv.Atoi(token.ExpiresIn); err == nil {
		ttl := time.Duration(expiresIn) * time.Second
		m.tokenCache.Set(cacheKey, token, ttl)
		
		// Schedule automatic refresh
		m.scheduleTokenRefresh(cacheKey, ttl-m.refreshThreshold)
		
		m.logger.WithField("ttl", ttl).Debug("B2B token cached with auto-refresh")
	}

	// Notify any waiting goroutines
	if resultChan, exists := m.requestGroup.Load(cacheKey); exists {
		select {
		case resultChan.(chan *TokenResponse) <- token:
		default:
		}
	}

	return token, nil
}

// GetCustomerAccessToken retrieves a B2B2C access token with optimizations
func (m *OptimizedAccessTokenManager) GetCustomerAccessToken(
	ctx context.Context,
	request CustomerTokenRequest,
) (*TokenResponse, error) {
	// Validate request
	if err := m.validateCustomerTokenRequest(request); err != nil {
		return nil, err
	}

	// Create cache key based on auth code
	cacheKey := fmt.Sprintf("b2b2c_token_%s", request.AuthCode)

	// Try to get from cache first
	if token, found := m.tokenCache.Get(cacheKey); found {
		m.logger.Debug("B2B2C token retrieved from cache")
		m.metrics.RecordTokenRequest(true)
		return token, nil
	}

	m.metrics.RecordTokenRequest(false)

	// Use request deduplication
	if result, loaded := m.requestGroup.LoadOrStore(cacheKey, make(chan *TokenResponse, 1)); loaded {
		resultChan := result.(chan *TokenResponse)
		select {
		case token := <-resultChan:
			if token != nil {
				return token, nil
			}
		case <-ctx.Done():
			return nil, errors.NewNetworkError(ctx.Err(), "request context canceled")
		case <-time.After(30 * time.Second):
			m.requestGroup.Delete(cacheKey)
		}
	}

	defer m.requestGroup.Delete(cacheKey)

	// Prevent concurrent requests for the same token
	m.b2b2cMutex.Lock()
	defer m.b2b2cMutex.Unlock()

	// Double-check cache after acquiring lock
	if token, found := m.tokenCache.Get(cacheKey); found {
		m.logger.Debug("B2B2C token retrieved from cache after lock")
		return token, nil
	}

	// Fetch new token
	token, err := m.fetchB2B2CToken(ctx, request)
	if err != nil {
		return nil, err
	}

	// Cache the token
	if expiresIn, err := strconv.Atoi(token.ExpiresIn); err == nil {
		ttl := time.Duration(expiresIn) * time.Second
		m.tokenCache.Set(cacheKey, token, ttl)
		
		// Schedule automatic refresh
		m.scheduleTokenRefresh(cacheKey, ttl-m.refreshThreshold)
		
		m.logger.WithField("ttl", ttl).Debug("B2B2C token cached with auto-refresh")
	}

	// Notify any waiting goroutines
	if resultChan, exists := m.requestGroup.Load(cacheKey); exists {
		select {
		case resultChan.(chan *TokenResponse) <- token:
		default:
		}
	}

	return token, nil
}

// fetchB2BToken fetches a new B2B token from the API
func (m *OptimizedAccessTokenManager) fetchB2BToken(ctx context.Context) (*TokenResponse, error) {
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

// fetchB2B2CToken fetches a new B2B2C token from the API
func (m *OptimizedAccessTokenManager) fetchB2B2CToken(ctx context.Context, request CustomerTokenRequest) (*TokenResponse, error) {
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

// scheduleTokenRefresh schedules automatic token refresh
func (m *OptimizedAccessTokenManager) scheduleTokenRefresh(cacheKey string, refreshIn time.Duration) {
	m.refreshMutex.Lock()
	defer m.refreshMutex.Unlock()

	// Cancel existing timer if any
	if timer, exists := m.refreshTimers[cacheKey]; exists {
		timer.Stop()
	}

	// Schedule new refresh
	timer := time.AfterFunc(refreshIn, func() {
		m.refreshToken(cacheKey)
	})
	
	m.refreshTimers[cacheKey] = timer
}

// refreshToken refreshes a token in the background
func (m *OptimizedAccessTokenManager) refreshToken(cacheKey string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	m.logger.WithField("cache_key", cacheKey).Info("Refreshing token in background")

	if cacheKey == "b2b_token" {
		_, err := m.GetAccessToken(ctx)
		if err != nil {
			m.logger.WithError(err).Error("Failed to refresh B2B token")
		}
	} else if strings.HasPrefix(cacheKey, "b2b2c_token_") {
		// For B2B2C tokens, we would need to store the original request
		// This is a simplified implementation
		m.logger.WithField("cache_key", cacheKey).Warn("B2B2C token refresh not implemented")
	}

	// Clean up timer
	m.refreshMutex.Lock()
	delete(m.refreshTimers, cacheKey)
	m.refreshMutex.Unlock()
}

// startTokenRefreshRoutine starts the background token refresh routine
func (m *OptimizedAccessTokenManager) startTokenRefreshRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// Check for tokens that need refresh
		// This is a placeholder for more sophisticated refresh logic
		m.logger.Debug("Token refresh routine running")
	}
}

// createAuthHeaders creates headers with asymmetric signature for authentication
func (m *OptimizedAccessTokenManager) createAuthHeaders() (map[string]string, error) {
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
func (m *OptimizedAccessTokenManager) validateCustomerTokenRequest(request CustomerTokenRequest) error {
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

// CleanupCache removes expired tokens from cache
func (m *OptimizedAccessTokenManager) CleanupCache() {
	m.tokenCache.Cleanup()
}

// GetCacheStats returns cache statistics
func (m *OptimizedAccessTokenManager) GetCacheStats() map[string]interface{} {
	m.refreshMutex.RLock()
	activeRefreshTimers := len(m.refreshTimers)
	m.refreshMutex.RUnlock()

	return map[string]interface{}{
		"active_refresh_timers": activeRefreshTimers,
		"refresh_threshold":     m.refreshThreshold.String(),
	}
}

// Shutdown gracefully shuts down the token manager
func (m *OptimizedAccessTokenManager) Shutdown() {
	m.refreshMutex.Lock()
	defer m.refreshMutex.Unlock()

	// Cancel all refresh timers
	for _, timer := range m.refreshTimers {
		timer.Stop()
	}
	m.refreshTimers = make(map[string]*time.Timer)

	m.logger.Info("Access token manager shutdown complete")
}