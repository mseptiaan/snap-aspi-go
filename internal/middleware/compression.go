package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CompressionMiddleware provides gzip compression for responses
func CompressionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if client accepts gzip
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// Skip compression for small responses or certain content types
		if shouldSkipCompression(c) {
			c.Next()
			return
		}

		// Set compression headers
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")

		// Create gzip writer
		gz := gzip.NewWriter(c.Writer)
		defer gz.Close()

		// Wrap the response writer
		c.Writer = &gzipWriter{
			ResponseWriter: c.Writer,
			Writer:         gz,
		}

		c.Next()
	}
}

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.Writer.Write(data)
}

func (g *gzipWriter) WriteString(s string) (int, error) {
	return g.Writer.Write([]byte(s))
}

func shouldSkipCompression(c *gin.Context) bool {
	// Skip for certain paths
	skipPaths := []string{"/health", "/metrics"}
	for _, path := range skipPaths {
		if c.Request.URL.Path == path {
			return true
		}
	}

	// Skip for certain content types
	contentType := c.GetHeader("Content-Type")
	skipTypes := []string{"image/", "video/", "audio/"}
	for _, skipType := range skipTypes {
		if strings.HasPrefix(contentType, skipType) {
			return true
		}
	}

	return false
}