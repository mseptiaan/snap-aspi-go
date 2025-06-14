package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	// GetB2BToken retrieves a B2B access token
	GetB2BToken(ctx context.Context) (*auth.TokenResponse, error)
	
	// GetB2B2CToken retrieves a B2B2C access token for customer
	GetB2B2CToken(ctx context.Context, req auth.CustomerTokenRequest) (*auth.TokenResponse, error)
}

// authService implements AuthService
type authService struct {
	manager *auth.AccessTokenManager
}

// GetB2BToken retrieves a B2B access token
func (a *authService) GetB2BToken(ctx context.Context) (*auth.TokenResponse, error) {
	return a.manager.GetAccessToken(ctx)
}

// GetB2B2CToken retrieves a B2B2C access token for customer
func (a *authService) GetB2B2CToken(ctx context.Context, req auth.CustomerTokenRequest) (*auth.TokenResponse, error) {
	return a.manager.GetCustomerAccessToken(ctx, req)
}