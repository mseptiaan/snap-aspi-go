package metrics

import (
	"sync"
	"sync/atomic"
	"time"
)

// Metrics collects application performance metrics
type Metrics struct {
	mu sync.RWMutex

	// Request metrics
	TotalRequests     int64
	SuccessfulRequests int64
	FailedRequests    int64
	
	// Response time metrics
	TotalResponseTime time.Duration
	MinResponseTime   time.Duration
	MaxResponseTime   time.Duration
	
	// Cache metrics
	CacheHits   int64
	CacheMisses int64
	
	// Token metrics
	TokenRequests int64
	TokenCacheHits int64
	
	// Error metrics
	ErrorsByType map[string]int64
	
	// Connection metrics
	ActiveConnections int64
	TotalConnections  int64
}

// NewMetrics creates a new metrics collector
func NewMetrics() *Metrics {
	return &Metrics{
		ErrorsByType: make(map[string]int64),
		MinResponseTime: time.Hour, // Initialize to high value
	}
}

// RecordRequest records a request metric
func (m *Metrics) RecordRequest(duration time.Duration, success bool) {
	atomic.AddInt64(&m.TotalRequests, 1)
	
	if success {
		atomic.AddInt64(&m.SuccessfulRequests, 1)
	} else {
		atomic.AddInt64(&m.FailedRequests, 1)
	}
	
	// Update response time metrics
	m.mu.Lock()
	m.TotalResponseTime += duration
	if duration < m.MinResponseTime {
		m.MinResponseTime = duration
	}
	if duration > m.MaxResponseTime {
		m.MaxResponseTime = duration
	}
	m.mu.Unlock()
}

// RecordCacheHit records a cache hit
func (m *Metrics) RecordCacheHit() {
	atomic.AddInt64(&m.CacheHits, 1)
}

// RecordCacheMiss records a cache miss
func (m *Metrics) RecordCacheMiss() {
	atomic.AddInt64(&m.CacheMisses, 1)
}

// RecordTokenRequest records a token request
func (m *Metrics) RecordTokenRequest(fromCache bool) {
	atomic.AddInt64(&m.TokenRequests, 1)
	if fromCache {
		atomic.AddInt64(&m.TokenCacheHits, 1)
	}
}

// RecordError records an error by type
func (m *Metrics) RecordError(errorType string) {
	m.mu.Lock()
	m.ErrorsByType[errorType]++
	m.mu.Unlock()
}

// RecordConnection records connection metrics
func (m *Metrics) RecordConnection(active bool) {
	if active {
		atomic.AddInt64(&m.ActiveConnections, 1)
		atomic.AddInt64(&m.TotalConnections, 1)
	} else {
		atomic.AddInt64(&m.ActiveConnections, -1)
	}
}

// GetStats returns current metrics as a map
func (m *Metrics) GetStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	totalRequests := atomic.LoadInt64(&m.TotalRequests)
	successRate := float64(0)
	if totalRequests > 0 {
		successRate = float64(atomic.LoadInt64(&m.SuccessfulRequests)) / float64(totalRequests)
	}
	
	avgResponseTime := time.Duration(0)
	if totalRequests > 0 {
		avgResponseTime = m.TotalResponseTime / time.Duration(totalRequests)
	}
	
	cacheTotal := atomic.LoadInt64(&m.CacheHits) + atomic.LoadInt64(&m.CacheMisses)
	cacheHitRatio := float64(0)
	if cacheTotal > 0 {
		cacheHitRatio = float64(atomic.LoadInt64(&m.CacheHits)) / float64(cacheTotal)
	}
	
	tokenTotal := atomic.LoadInt64(&m.TokenRequests)
	tokenCacheRatio := float64(0)
	if tokenTotal > 0 {
		tokenCacheRatio = float64(atomic.LoadInt64(&m.TokenCacheHits)) / float64(tokenTotal)
	}
	
	return map[string]interface{}{
		"requests": map[string]interface{}{
			"total":        totalRequests,
			"successful":   atomic.LoadInt64(&m.SuccessfulRequests),
			"failed":       atomic.LoadInt64(&m.FailedRequests),
			"success_rate": successRate,
		},
		"response_time": map[string]interface{}{
			"average": avgResponseTime.String(),
			"min":     m.MinResponseTime.String(),
			"max":     m.MaxResponseTime.String(),
		},
		"cache": map[string]interface{}{
			"hits":      atomic.LoadInt64(&m.CacheHits),
			"misses":    atomic.LoadInt64(&m.CacheMisses),
			"hit_ratio": cacheHitRatio,
		},
		"tokens": map[string]interface{}{
			"requests":         tokenTotal,
			"cache_hits":       atomic.LoadInt64(&m.TokenCacheHits),
			"cache_hit_ratio":  tokenCacheRatio,
		},
		"connections": map[string]interface{}{
			"active": atomic.LoadInt64(&m.ActiveConnections),
			"total":  atomic.LoadInt64(&m.TotalConnections),
		},
		"errors": m.ErrorsByType,
	}
}

// Reset resets all metrics
func (m *Metrics) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	atomic.StoreInt64(&m.TotalRequests, 0)
	atomic.StoreInt64(&m.SuccessfulRequests, 0)
	atomic.StoreInt64(&m.FailedRequests, 0)
	atomic.StoreInt64(&m.CacheHits, 0)
	atomic.StoreInt64(&m.CacheMisses, 0)
	atomic.StoreInt64(&m.TokenRequests, 0)
	atomic.StoreInt64(&m.TokenCacheHits, 0)
	atomic.StoreInt64(&m.ActiveConnections, 0)
	atomic.StoreInt64(&m.TotalConnections, 0)
	
	m.TotalResponseTime = 0
	m.MinResponseTime = time.Hour
	m.MaxResponseTime = 0
	m.ErrorsByType = make(map[string]int64)
}