package mocks

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/stretchr/testify/mock"
)

// MockAccessTokenManager is a mock implementation of auth.AccessTokenManager
type MockAccessTokenManager struct {
	mock.Mock
}

// GetAccessToken returns a B2B access token
func (m *MockAccessTokenManager) GetAccessToken(ctx context.Context) (*auth.TokenResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.TokenResponse), args.Error(1)
}

// GetCustomerAccessToken returns a B2B2C access token for customer
func (m *MockAccessTokenManager) GetCustomerAccessToken(
	ctx context.Context,
	req auth.CustomerTokenRequest,
) (*auth.TokenResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.TokenResponse), args.Error(1)
}

// RefreshAccessToken refreshes an existing access token
func (m *MockAccessTokenManager) RefreshAccessToken(
	ctx context.Context,
	refreshToken string,
) (*auth.TokenResponse, error) {
	args := m.Called(ctx, refreshToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.TokenResponse), args.Error(1)
}

// IsTokenValid checks if a token is still valid
func (m *MockAccessTokenManager) IsTokenValid(token string) bool {
	args := m.Called(token)
	return args.Bool(0)
}
