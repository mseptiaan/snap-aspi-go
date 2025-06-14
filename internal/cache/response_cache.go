package cache

import (
	"encoding/json"
	"sync"
	"time"
)

// ResponseCache provides thread-safe caching for API responses
type ResponseCache struct {
	mu       sync.RWMutex
	cache    map[string]*CachedResponse
	maxSize  int
	ttl      time.Duration
	hitCount int64
	missCount int64
}

// CachedResponse represents a cached API response
type CachedResponse struct {
	Data      []byte
	ExpiresAt time.Time
	CreatedAt time.Time
}

// NewResponseCache creates a new response cache
func NewResponseCache(maxSize int, ttl time.Duration) *ResponseCache {
	return &ResponseCache{
		cache:   make(map[string]*CachedResponse),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

// Get retrieves a response from cache
func (c *ResponseCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cached, exists := c.cache[key]
	if !exists {
		c.missCount++
		return nil, false
	}

	// Check if expired
	if time.Now().After(cached.ExpiresAt) {
		c.missCount++
		return nil, false
	}

	c.hitCount++
	return cached.Data, true
}

// Set stores a response in cache
func (c *ResponseCache) Set(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if we need to evict old entries
	if len(c.cache) >= c.maxSize {
		c.evictOldest()
	}

	c.cache[key] = &CachedResponse{
		Data:      data,
		ExpiresAt: time.Now().Add(c.ttl),
		CreatedAt: time.Now(),
	}
}

// evictOldest removes the oldest entry from cache
func (c *ResponseCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, cached := range c.cache {
		if oldestKey == "" || cached.CreatedAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = cached.CreatedAt
		}
	}

	if oldestKey != "" {
		delete(c.cache, oldestKey)
	}
}

// Cleanup removes expired entries
func (c *ResponseCache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, cached := range c.cache {
		if now.After(cached.ExpiresAt) {
			delete(c.cache, key)
		}
	}
}

// Stats returns cache statistics
func (c *ResponseCache) Stats() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	total := c.hitCount + c.missCount
	hitRatio := float64(0)
	if total > 0 {
		hitRatio = float64(c.hitCount) / float64(total)
	}

	return map[string]interface{}{
		"size":      len(c.cache),
		"max_size":  c.maxSize,
		"hits":      c.hitCount,
		"misses":    c.missCount,
		"hit_ratio": hitRatio,
	}
}