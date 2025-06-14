package cache

import (
	"sync"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
)

// TokenCache provides thread-safe caching for access tokens
type TokenCache struct {
	mu     sync.RWMutex
	tokens map[string]*CachedToken
}

// CachedToken represents a cached token with expiration
type CachedToken struct {
	Token     *auth.TokenResponse
	ExpiresAt time.Time
}

// NewTokenCache creates a new token cache
func NewTokenCache() *TokenCache {
	return &TokenCache{
		tokens: make(map[string]*CachedToken),
	}
}

// Get retrieves a token from cache if valid
func (c *TokenCache) Get(key string) (*auth.TokenResponse, bool) {
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
func (c *TokenCache) Set(key string, token *auth.TokenResponse, ttl time.Duration) {
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