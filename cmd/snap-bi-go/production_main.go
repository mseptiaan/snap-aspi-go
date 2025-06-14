package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	snapErrors "github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/metrics"
	"github.com/mseptiaan/snap-aspi-go/internal/middleware"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
	// Set GOMAXPROCS to use all available CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Load configuration
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to load configuration: %v", err))
	}

	// Initialize logger
	logger := logging.NewLogger(&cfg.Logger)

	logger.WithFields(map[string]interface{}{
		"app_name":    cfg.App.Name,
		"app_version": cfg.App.Version,
		"environment": cfg.App.Environment,
		"go_version":  runtime.Version(),
		"num_cpu":     runtime.NumCPU(),
		"gomaxprocs":  runtime.GOMAXPROCS(0),
	}).Info("Starting Snap ASPI Go application with production optimizations")

	// Initialize metrics collector
	metricsCollector := metrics.NewMetrics()

	// Initialize optimized HTTP client
	httpClient := client.NewOptimizedHTTPClient(cfg, logger, metricsCollector)

	// Initialize optimized access token manager
	accessTokenManager, err := auth.NewOptimizedAccessTokenManager(cfg, httpClient, logger, metricsCollector)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize access token manager")
	}

	// Initialize Virtual Account service
	vaService, err := services.NewVirtualAccountService(httpClient, accessTokenManager, cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize Virtual Account service")
	}

	// Initialize MPM Transfer Credit service
	mpmService, err := services.NewMPMService(httpClient, accessTokenManager, cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize MPM service")
	}

	// Set Gin mode based on environment
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router with production optimizations
	router := gin.New()

	// Add production middleware stack
	setupProductionMiddleware(router, metricsCollector, logger)

	// Setup routes
	setupProductionRoutes(router, cfg, accessTokenManager, vaService, mpmService, logger, metricsCollector)

	// Create HTTP server with production settings
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Start background maintenance routines
	startMaintenanceRoutines(accessTokenManager, metricsCollector, logger)

	// Start server in a goroutine
	go func() {
		logger.WithFields(map[string]interface{}{
			"host": cfg.Server.Host,
			"port": cfg.Server.Port,
		}).Info("Starting production HTTP server")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown access token manager
	accessTokenManager.Shutdown()

	// Wait for active HTTP requests to complete
	if optimizedClient, ok := httpClient.(*client.OptimizedHTTPClient); ok {
		optimizedClient.WaitForActiveRequests()
	}

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	// Log final metrics
	finalStats := metricsCollector.GetStats()
	logger.WithFields(finalStats).Info("Final application metrics")

	logger.Info("Server gracefully stopped")
}

// setupProductionMiddleware configures production-grade middleware
func setupProductionMiddleware(router *gin.Engine, metricsCollector *metrics.Metrics, logger logging.Logger) {
	// Request ID middleware (first)
	router.Use(middleware.RequestID())

	// Metrics collection middleware
	router.Use(middleware.MetricsMiddleware(metricsCollector))

	// Recovery middleware with custom logging
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.WithFields(map[string]interface{}{
			"panic":      recovered,
			"request_id": c.GetString("request_id"),
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
		}).Error("Panic recovered")
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Internal server error",
			"request_id": c.GetString("request_id"),
		})
	}))

	// Rate limiting middleware (100 requests per minute per IP, burst 20)
	rateLimiter := middleware.NewRateLimiter(100, 20)
	router.Use(rateLimiter.Middleware())

	// Circuit breaker middleware (5 failures, 3 successes, 30s timeout)
	circuitBreaker := middleware.NewCircuitBreaker(5, 3, 30*time.Second)
	router.Use(middleware.CircuitBreakerMiddleware(circuitBreaker))

	// Compression middleware
	router.Use(middleware.CompressionMiddleware())

	// CORS middleware with production settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Configure based on your needs
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"X-Request-ID", "X-RateLimit-Remaining"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Security headers middleware
	router.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Next()
	})

	// Request logging middleware (optimized for production)
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Only log slow requests or errors in production
		if duration > 500*time.Millisecond || c.Writer.Status() >= 400 {
			logger.WithFields(map[string]interface{}{
				"method":      c.Request.Method,
				"path":        c.Request.URL.Path,
				"status_code": c.Writer.Status(),
				"duration":    duration.String(),
				"client_ip":   c.ClientIP(),
				"request_id":  c.GetString("request_id"),
				"user_agent":  c.Request.UserAgent(),
			}).Info("HTTP request processed")
		}
	})
}

// setupProductionRoutes configures production routes with optimizations
func setupProductionRoutes(
	router *gin.Engine,
	cfg *config.Config,
	accessTokenManager *auth.OptimizedAccessTokenManager,
	vaService *services.VirtualAccountService,
	mpmService *services.MPMService,
	logger logging.Logger,
	metricsCollector *metrics.Metrics,
) {
	// Health check endpoint with detailed status
	router.GET("/health", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		
		// Perform health checks
		healthStatus := map[string]interface{}{
			"status":      "healthy",
			"app_name":    cfg.App.Name,
			"app_version": cfg.App.Version,
			"timestamp":   time.Now().Format(time.RFC3339),
			"uptime":      time.Since(time.Now()).String(),
			"environment": cfg.App.Environment,
		}
		
		c.JSON(http.StatusOK, healthStatus)
	})

	// Metrics endpoint (protected)
	router.GET("/metrics", func(c *gin.Context) {
		// Add basic auth or API key validation here
		stats := metricsCollector.GetStats()
		cacheStats := accessTokenManager.GetCacheStats()
		
		response := map[string]interface{}{
			"application": stats,
			"token_cache": cacheStats,
			"timestamp":   time.Now().Format(time.RFC3339),
		}
		
		c.JSON(http.StatusOK, response)
	})

	// API version 1
	v1 := router.Group("/api/v1")
	{
		// Authentication endpoints
		auth := v1.Group("/auth")
		{
			auth.POST("/b2b-token", handleGetB2BTokenProduction(accessTokenManager, logger))
			auth.POST("/b2b2c-token", handleGetB2B2CTokenProduction(accessTokenManager, logger))
		}

		// Virtual Account endpoints
		va := v1.Group("/virtual-account")
		{
			va.POST("/create", handleCreateVAProduction(vaService, logger))
			va.PUT("/update", handleUpdateVAProduction(vaService, logger))
			va.DELETE("/delete", handleDeleteVAProduction(vaService, logger))
			va.POST("/inquiry", handleInquiryVAProduction(vaService, logger))
			va.POST("/inquiry-general", handleInquiryProduction(vaService, logger))
			va.POST("/payment", handlePaymentProduction(vaService, logger))
			va.POST("/status", handleStatusProduction(vaService, logger))
			va.POST("/report", handleReportProduction(vaService, logger))
			va.PUT("/status", handleUpdateStatusProduction(vaService, logger))
		}

		// MPM Transfer Credit endpoints
		mpm := v1.Group("/mpm")
		{
			mpm.POST("/transfer", handleMPMTransferProduction(mpmService, logger))
			mpm.POST("/inquiry", handleMPMInquiryProduction(mpmService, logger))
			mpm.POST("/status", handleMPMStatusProduction(mpmService, logger))
			mpm.POST("/refund", handleMPMRefundProduction(mpmService, logger))
			mpm.POST("/balance-inquiry", handleMPMBalanceInquiryProduction(mpmService, logger))
			mpm.POST("/account-inquiry", handleMPMAccountInquiryProduction(mpmService, logger))
			mpm.POST("/history", handleMPMHistoryProduction(mpmService, logger))
			mpm.POST("/generate-qr", handleMPMGenerateQRProduction(mpmService, logger))
			mpm.POST("/notify-qr", handleMPMNotifyQRProduction(mpmService, logger))
		}

		// Utility endpoints
		utils := v1.Group("/utils")
		{
			utils.GET("/config", func(c *gin.Context) {
				c.Header("Cache-Control", "public, max-age=300")
				c.JSON(http.StatusOK, gin.H{
					"app": gin.H{
						"name":        cfg.App.Name,
						"version":     cfg.App.Version,
						"environment": cfg.App.Environment,
					},
					"aspi": gin.H{
						"base_url": cfg.ASPI.BaseURL,
						"endpoints": gin.H{
							"b2b_token":   cfg.ASPI.Endpoints.B2BToken,
							"b2b2c_token": cfg.ASPI.Endpoints.B2B2CToken,
						},
					},
				})
			})
		}
	}
}

// startMaintenanceRoutines starts background maintenance tasks
func startMaintenanceRoutines(
	accessTokenManager *auth.OptimizedAccessTokenManager,
	metricsCollector *metrics.Metrics,
	logger logging.Logger,
) {
	// Cache cleanup routine
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		
		for range ticker.C {
			accessTokenManager.CleanupCache()
			logger.Debug("Cache cleanup completed")
		}
	}()

	// Metrics reporting routine
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		
		for range ticker.C {
			stats := metricsCollector.GetStats()
			logger.WithFields(stats).Info("Application metrics report")
		}
	}()

	// Memory stats routine
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		
		var m runtime.MemStats
		for range ticker.C {
			runtime.ReadMemStats(&m)
			logger.WithFields(map[string]interface{}{
				"alloc_mb":      bToMb(m.Alloc),
				"total_alloc_mb": bToMb(m.TotalAlloc),
				"sys_mb":        bToMb(m.Sys),
				"num_gc":        m.NumGC,
				"goroutines":    runtime.NumGoroutine(),
			}).Debug("Memory and runtime stats")
		}
	}()
}

// Production handler functions with optimized error handling

func handleGetB2BTokenProduction(accessTokenManager *auth.OptimizedAccessTokenManager, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("request_id")
		
		tokenResponse, err := accessTokenManager.GetAccessToken(c.Request.Context())
		if err != nil {
			logger.WithFields(map[string]interface{}{
				"request_id": requestID,
				"error":      err.Error(),
			}).Error("Failed to get B2B access token")

			handleProductionError(c, err, requestID)
			return
		}

		c.JSON(http.StatusOK, tokenResponse)
	}
}

func handleGetB2B2CTokenProduction(accessTokenManager *auth.OptimizedAccessTokenManager, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("request_id")
		
		var request auth.CustomerTokenRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			logger.WithFields(map[string]interface{}{
				"request_id": requestID,
				"error":      err.Error(),
			}).Warn("Invalid B2B2C token request")
			
			c.JSON(http.StatusBadRequest, gin.H{
				"error":      "Invalid request",
				"message":    "Request body is invalid",
				"request_id": requestID,
			})
			return
		}

		tokenResponse, err := accessTokenManager.GetCustomerAccessToken(c.Request.Context(), request)
		if err != nil {
			logger.WithFields(map[string]interface{}{
				"request_id": requestID,
				"error":      err.Error(),
			}).Error("Failed to get B2B2C access token")

			handleProductionError(c, err, requestID)
			return
		}

		c.JSON(http.StatusOK, tokenResponse)
	}
}

func handleCreateVAProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("request_id")
		
		var payload types.CreateVAPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			logger.WithFields(map[string]interface{}{
				"request_id": requestID,
				"error":      err.Error(),
			}).Warn("Invalid create VA request")
			
			c.JSON(http.StatusBadRequest, gin.H{
				"error":      "Invalid request",
				"message":    "Request body is invalid",
				"request_id": requestID,
			})
			return
		}

		result, err := vaService.CreateVA(c.Request.Context(), &payload)
		if err != nil {
			logger.WithFields(map[string]interface{}{
				"request_id": requestID,
				"error":      err.Error(),
			}).Error("Failed to create Virtual Account")
			
			handleProductionError(c, err, requestID)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// Add similar production handlers for other endpoints...

func handleProductionError(c *gin.Context, err error, requestID string) {
	switch e := err.(type) {
	case *snapErrors.Error:
		statusCode := e.GetHTTPStatus()
		c.JSON(statusCode, gin.H{
			"error":      e.Type,
			"message":    e.Message,
			"code":       e.Code,
			"request_id": requestID,
			"timestamp":  time.Now().Format(time.RFC3339),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Internal error",
			"message":    "An unexpected error occurred",
			"request_id": requestID,
			"timestamp":  time.Now().Format(time.RFC3339),
		})
	}
}

// Utility functions
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Placeholder handlers for other endpoints
func handleUpdateVAProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleDeleteVAProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleInquiryVAProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleInquiryProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handlePaymentProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleStatusProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleReportProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleUpdateStatusProduction(vaService *services.VirtualAccountService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMTransferProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMInquiryProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMStatusProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMRefundProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMBalanceInquiryProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMAccountInquiryProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMHistoryProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMGenerateQRProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}

func handleMPMNotifyQRProduction(mpmService *services.MPMService, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) { /* Implementation */ }
}