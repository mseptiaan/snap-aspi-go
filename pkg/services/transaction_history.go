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

// TransactionHistoryService implements the Transaction History business logic
type TransactionHistoryService struct {
	httpClient  contracts.HTTPClient
	authManager *auth.AccessTokenManager
	symSigner   *signature.SymmetricSigner
	config      *config.Config
	logger      logging.Logger

	// ASPI Transaction History API endpoints
	endpoints map[string]string
	methods   map[string]string
}

// NewTransactionHistoryService creates a new Transaction History service instance
func NewTransactionHistoryService(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) (*TransactionHistoryService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	return &TransactionHistoryService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		// ASPI Transaction History API endpoints
		endpoints: map[string]string{
			"transaction-history-list":   "/api/v1.0/transaction-history-list",
			"transaction-history-detail": "/api/v1.0/transaction-history-detail",
			"bank-statement":             "/api/v1.0/bank-statement",
		},
		methods: map[string]string{
			"transaction-history-list":   "POST",
			"transaction-history-detail": "POST",
			"bank-statement":             "POST",
		},
	}, nil
}

// NewTransactionHistoryServiceWithEndpoints creates a new Transaction History service with custom endpoints
func NewTransactionHistoryServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*TransactionHistoryService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getTransactionHistoryEndpoints(customEndpoints)
	methods := getTransactionHistoryMethods()

	return &TransactionHistoryService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getTransactionHistoryEndpoints returns Transaction History endpoints (custom or default)
func getTransactionHistoryEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
		"transaction-history-list":   "/api/v1.0/transaction-history-list",
		"transaction-history-detail": "/api/v1.0/transaction-history-detail",
		"bank-statement":             "/api/v1.0/bank-statement",
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

// getTransactionHistoryMethods returns HTTP methods for Transaction History endpoints
func getTransactionHistoryMethods() map[string]string {
	return map[string]string{
		"transaction-history-list":   "POST",
		"transaction-history-detail": "POST",
		"bank-statement":             "POST",
	}
}

// TransactionHistoryList retrieves transaction history list
func (t *TransactionHistoryService) TransactionHistoryList(
	ctx context.Context,
	payload *types.TransactionHistoryListPayload,
) (*types.TransactionHistoryResponse, error) {
	result, err := t.makeRequest(ctx, "transaction-history-list", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction history list: %w", err)
	}

	// Convert result to TransactionHistoryResponse
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transaction history response: %w", err)
	}

	var historyResponse types.TransactionHistoryResponse
	if err := json.Unmarshal(resultBytes, &historyResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transaction history response: %w", err)
	}

	return &historyResponse, nil
}

// TransactionHistoryDetail retrieves transaction history detail
func (t *TransactionHistoryService) TransactionHistoryDetail(
	ctx context.Context,
	payload *types.TransactionHistoryDetailPayload,
) (map[string]any, error) {
	return t.makeRequest(ctx, "transaction-history-detail", payload)
}

// BankStatement retrieves bank statement
func (t *TransactionHistoryService) BankStatement(
	ctx context.Context,
	payload *types.BankStatementPayload,
) (*types.BankStatementResponse, error) {
	result, err := t.makeRequest(ctx, "bank-statement", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve bank statement: %w", err)
	}

	// Convert result to BankStatementResponse
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bank statement response: %w", err)
	}

	var statementResponse types.BankStatementResponse
	if err := json.Unmarshal(resultBytes, &statementResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bank statement response: %w", err)
	}

	return &statementResponse, nil
}

// makeRequest handles the HTTP request for ASPI Transaction History API calls
func (t *TransactionHistoryService) makeRequest(ctx context.Context, endpoint string, payload any) (map[string]any, error) {
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
		"service":   "TransactionHistoryService",
		"endpoint":  endpoint,
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
	}).Info("Making Transaction History API request")

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
			"service":  "TransactionHistoryService",
			"endpoint": endpoint,
			"error":    err.Error(),
		}).Error("Transaction History API request failed")
		return nil, fmt.Errorf("Transaction History API request failed: %w", err)
	}

	// Log successful response
	t.logger.WithFields(map[string]any{
		"service":     "TransactionHistoryService",
		"endpoint":    endpoint,
		"status_code": response.StatusCode,
	}).Info("Transaction History API request completed")

	// Parse response
	var result map[string]any
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse Transaction History API response: %w", err)
	}

	return result, nil
}