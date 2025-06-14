package contracts

import (
	"context"
	"time"
)

// HTTPClient represents the interface for HTTP client operations
type HTTPClient interface {
	// Get performs a GET request
	Get(ctx context.Context, url string, headers map[string]string) (*Response, error)

	// Post performs a POST request with JSON body
	Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*Response, error)

	// Put performs a PUT request with JSON body
	Put(ctx context.Context, url string, body interface{}, headers map[string]string) (*Response, error)

	// Delete performs a DELETE request
	Delete(ctx context.Context, url string, headers map[string]string) (*Response, error)

	// WithTimeout sets request timeout
	WithTimeout(timeout time.Duration) HTTPClient

	// WithRetry sets retry configuration
	WithRetry(maxAttempts int, backoffDuration time.Duration) HTTPClient

	// WithBaseURL sets base URL for all requests
	WithBaseURL(baseURL string) HTTPClient
}

// Response represents an HTTP response
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
	Success    bool
}

// JSON unmarshals response body to the provided interface
func (r *Response) JSON(v interface{}) error {
	// Implementation will be in the concrete response type
	return nil
}

// String returns response body as string
func (r *Response) String() string {
	return string(r.Body)
}

// IsSuccessful returns true if status code indicates success (2xx)
func (r *Response) IsSuccessful() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// TokenStorage represents the interface for token storage operations
type TokenStorage interface {
	// Store saves a token with the given key
	Store(key string, token *AccessToken) error

	// Retrieve gets a token by key
	Retrieve(key string) (*AccessToken, error)

	// Delete removes a token by key
	Delete(key string) error

	// IsValid checks if a token is still valid
	IsValid(key string) (bool, error)

	// Cleanup removes expired tokens
	Cleanup() error
}

// AccessToken represents an access token with metadata
type AccessToken struct {
	Token     string    `json:"accessToken"`
	Type      string    `json:"tokenType"`
	ExpiresIn int       `json:"expiresIn"`
	ExpiresAt time.Time `json:"expiresAt"`
	Scope     string    `json:"scope,omitempty"`
}

// IsExpired checks if the token has expired
func (t *AccessToken) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// Signer represents the interface for cryptographic signing operations
type Signer interface {
	// SignAuth creates authentication signature
	SignAuth(clientKey string, timestamp time.Time) (string, error)

	// SignPayload creates payload signature for API requests
	SignPayload(method, url, body, timestamp string) (string, error)

	// VerifySignature verifies a signature
	VerifySignature(data, signature string) error

	// GetPublicKey returns the public key for signature verification
	GetPublicKey() (interface{}, error)
}

// AuthService represents the interface for authentication operations
type AuthService interface {
	// GetAccessToken retrieves B2B access token
	GetAccessToken(ctx context.Context) (*AccessToken, error)

	// GetCustomerAccessToken retrieves B2B2C access token
	GetCustomerAccessToken(
		ctx context.Context,
		authCode, refreshToken string,
		additionalInfo map[string]interface{},
	) (*AccessToken, error)

	// RefreshToken refreshes an existing token
	RefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)

	// ValidateToken validates token format and expiration
	ValidateToken(token string) error
}

// PaymentService represents the interface for payment operations
type PaymentService interface {
	// TransferCredit performs credit transfer operation
	TransferCredit(ctx context.Context, request *TransferCreditRequest) (*TransferCreditResponse, error)

	// GetTransferStatus gets the status of a transfer
	GetTransferStatus(ctx context.Context, transferID string) (*TransferStatusResponse, error)

	// GetAccountBalance gets account balance
	GetAccountBalance(ctx context.Context, accountID string) (*BalanceResponse, error)
}

// TransferCreditRequest represents a credit transfer request
type TransferCreditRequest struct {
	Amount             float64                `json:"amount"`
	Currency           string                 `json:"currency"`
	SourceAccount      string                 `json:"sourceAccount"`
	DestinationAccount string                 `json:"destinationAccount"`
	Reference          string                 `json:"reference"`
	Description        string                 `json:"description,omitempty"`
	AdditionalInfo     map[string]interface{} `json:"additionalInfo,omitempty"`
}

// TransferCreditResponse represents a credit transfer response
type TransferCreditResponse struct {
	ResponseCode    string  `json:"responseCode"`
	ResponseMessage string  `json:"responseMessage"`
	TransferID      string  `json:"transferId"`
	Status          string  `json:"status"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Reference       string  `json:"reference"`
}

// TransferStatusResponse represents transfer status response
type TransferStatusResponse struct {
	ResponseCode    string    `json:"responseCode"`
	ResponseMessage string    `json:"responseMessage"`
	TransferID      string    `json:"transferId"`
	Status          string    `json:"status"`
	Amount          float64   `json:"amount"`
	Currency        string    `json:"currency"`
	ProcessedAt     time.Time `json:"processedAt"`
}

// BalanceResponse represents account balance response
type BalanceResponse struct {
	ResponseCode    string    `json:"responseCode"`
	ResponseMessage string    `json:"responseMessage"`
	AccountID       string    `json:"accountId"`
	Balance         float64   `json:"balance"`
	Currency        string    `json:"currency"`
	LastUpdated     time.Time `json:"lastUpdated"`
}
