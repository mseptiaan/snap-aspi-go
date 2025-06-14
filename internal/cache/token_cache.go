package cache

import (
	"sync"
	"time"
)

// TokenResponse represents the response from access token endpoints
type TokenResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccessToken     string `json:"accessToken"`
	TokenType       string `json:"tokenType"`
	ExpiresIn       string `json:"expiresIn"`
}

// TokenCache provides thread-safe caching for access tokens
type TokenCache struct {
	mu     sync.RWMutex
	tokens map[string]*CachedToken
}

// CachedToken represents a cached token with expiration
type CachedToken struct {
	Token     *TokenResponse
	ExpiresAt time.Time
}

// NewTokenCache creates a new token cache
func NewTokenCache() *TokenCache {
	return &TokenCache{
		tokens: make(map[string]*CachedToken),
	}
}

// Get retrieves a token from cache if valid
func (c *TokenCache) Get(key string) (*TokenResponse, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cached, exists := c.tokens[key]
	if !exists {
		return nil, false
	}

	// Check if token is expired (with 5 minute buffer)
	if time.Now().Add(5 * time.Minute).After(cached.ExpiresAt) {
		return nil, false
	}

	return cached.Token, true
}

// Set stores a token in cache
func (c *TokenCache) Set(key string, token *TokenResponse, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.tokens[key] = &CachedToken{
		Token:     token,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// Delete removes a token from cache
func (c *TokenCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.tokens, key)
}

// Cleanup removes expired tokens
func (c *TokenCache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, cached := range c.tokens {
		if now.After(cached.ExpiresAt) {
			delete(c.tokens, key)
		}
	}
}
