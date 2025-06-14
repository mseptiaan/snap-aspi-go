package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mseptiaan/snap-aspi-go/internal/metrics"
)

// MetricsMiddleware creates a middleware for collecting metrics
func MetricsMiddleware(m *metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		// Record active connection
		m.RecordConnection(true)
		defer m.RecordConnection(false)
		
		c.Next()
		
		// Record request metrics
		duration := time.Since(start)
		success := c.Writer.Status() < 400
		m.RecordRequest(duration, success)
		
		// Record errors if any
		if !success {
			errorType := "client_error"
			if c.Writer.Status() >= 500 {
				errorType = "server_error"
			}
			m.RecordError(errorType)
		}
	}
}