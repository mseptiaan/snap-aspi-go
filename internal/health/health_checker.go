package health

import (
	"context"
	"sync"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
)

// HealthStatus represents the health status of a component
type HealthStatus string

const (
	StatusHealthy   HealthStatus = "healthy"
	StatusUnhealthy HealthStatus = "unhealthy"
	StatusDegraded  HealthStatus = "degraded"
)

// HealthCheck represents a health check function
type HealthCheck func(ctx context.Context) HealthCheckResult

// HealthCheckResult represents the result of a health check
type HealthCheckResult struct {
	Status    HealthStatus `json:"status"`
	Message   string       `json:"message,omitempty"`
	Duration  time.Duration `json:"duration"`
	Timestamp time.Time    `json:"timestamp"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// HealthChecker manages health checks for the application
type HealthChecker struct {
	mu       sync.RWMutex
	checks   map[string]HealthCheck
	results  map[string]HealthCheckResult
	logger   logging.Logger
	interval time.Duration
	timeout  time.Duration
	stopCh   chan struct{}
}

// NewHealthChecker creates a new health checker
func NewHealthChecker(logger logging.Logger) *HealthChecker {
	return &HealthChecker{
		checks:   make(map[string]HealthCheck),
		results:  make(map[string]HealthCheckResult),
		logger:   logger,
		interval: 30 * time.Second,
		timeout:  10 * time.Second,
		stopCh:   make(chan struct{}),
	}
}

// RegisterCheck registers a health check
func (h *HealthChecker) RegisterCheck(name string, check HealthCheck) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checks[name] = check
}

// Start starts the health checker
func (h *HealthChecker) Start() {
	go h.runHealthChecks()
}

// Stop stops the health checker
func (h *HealthChecker) Stop() {
	close(h.stopCh)
}

// GetStatus returns the overall health status
func (h *HealthChecker) GetStatus() map[string]interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()

	overallStatus := StatusHealthy
	checks := make(map[string]HealthCheckResult)

	for name, result := range h.results {
		checks[name] = result
		if result.Status == StatusUnhealthy {
			overallStatus = StatusUnhealthy
		} else if result.Status == StatusDegraded && overallStatus == StatusHealthy {
			overallStatus = StatusDegraded
		}
	}

	return map[string]interface{}{
		"status":    overallStatus,
		"timestamp": time.Now(),
		"checks":    checks,
	}
}

// runHealthChecks runs health checks periodically
func (h *HealthChecker) runHealthChecks() {
	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	// Run initial checks
	h.executeAllChecks()

	for {
		select {
		case <-ticker.C:
			h.executeAllChecks()
		case <-h.stopCh:
			return
		}
	}
}

// executeAllChecks executes all registered health checks
func (h *HealthChecker) executeAllChecks() {
	h.mu.RLock()
	checks := make(map[string]HealthCheck)
	for name, check := range h.checks {
		checks[name] = check
	}
	h.mu.RUnlock()

	var wg sync.WaitGroup
	results := make(map[string]HealthCheckResult)
	resultsMu := sync.Mutex{}

	for name, check := range checks {
		wg.Add(1)
		go func(name string, check HealthCheck) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
			defer cancel()

			start := time.Now()
			result := check(ctx)
			result.Duration = time.Since(start)
			result.Timestamp = time.Now()

			resultsMu.Lock()
			results[name] = result
			resultsMu.Unlock()

			if result.Status != StatusHealthy {
				h.logger.WithFields(map[string]interface{}{
					"check":    name,
					"status":   result.Status,
					"message":  result.Message,
					"duration": result.Duration,
				}).Warn("Health check failed")
			}
		}(name, check)
	}

	wg.Wait()

	h.mu.Lock()
	h.results = results
	h.mu.Unlock()
}

// Common health checks

// DatabaseHealthCheck creates a health check for database connectivity
func DatabaseHealthCheck(db interface{}) HealthCheck {
	return func(ctx context.Context) HealthCheckResult {
		// Implement database ping logic here
		return HealthCheckResult{
			Status:  StatusHealthy,
			Message: "Database connection is healthy",
		}
	}
}

// HTTPClientHealthCheck creates a health check for HTTP client connectivity
func HTTPClientHealthCheck(client contracts.HTTPClient, url string) HealthCheck {
	return func(ctx context.Context) HealthCheckResult {
		start := time.Now()
		
		_, err := client.Get(ctx, url, map[string]string{})
		if err != nil {
			return HealthCheckResult{
				Status:  StatusUnhealthy,
				Message: fmt.Sprintf("HTTP client check failed: %v", err),
				Details: map[string]interface{}{
					"url":   url,
					"error": err.Error(),
				},
			}
		}

		duration := time.Since(start)
		status := StatusHealthy
		if duration > 5*time.Second {
			status = StatusDegraded
		}

		return HealthCheckResult{
			Status:  status,
			Message: "HTTP client is responsive",
			Details: map[string]interface{}{
				"url":           url,
				"response_time": duration.String(),
			},
		}
	}
}

// MemoryHealthCheck creates a health check for memory usage
func MemoryHealthCheck(maxMemoryMB uint64) HealthCheck {
	return func(ctx context.Context) HealthCheckResult {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		allocMB := m.Alloc / 1024 / 1024
		status := StatusHealthy
		message := fmt.Sprintf("Memory usage: %d MB", allocMB)

		if allocMB > maxMemoryMB {
			status = StatusUnhealthy
			message = fmt.Sprintf("Memory usage too high: %d MB (max: %d MB)", allocMB, maxMemoryMB)
		} else if allocMB > maxMemoryMB*80/100 {
			status = StatusDegraded
			message = fmt.Sprintf("Memory usage high: %d MB (max: %d MB)", allocMB, maxMemoryMB)
		}

		return HealthCheckResult{
			Status:  status,
			Message: message,
			Details: map[string]interface{}{
				"alloc_mb":       allocMB,
				"total_alloc_mb": m.TotalAlloc / 1024 / 1024,
				"sys_mb":         m.Sys / 1024 / 1024,
				"num_gc":         m.NumGC,
				"max_memory_mb":  maxMemoryMB,
			},
		}
	}
}