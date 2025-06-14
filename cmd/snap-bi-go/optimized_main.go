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
	}).Info("Starting Snap ASPI Go application")

	// Initialize optimized HTTP client
	httpClient := client.NewOptimizedSnapClient(cfg, logger)

	// Initialize optimized access token manager
	accessTokenManager, err := auth.NewOptimizedAccessTokenManager(cfg, httpClient, logger)
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

	// Create Gin router with optimized middleware
	router := gin.New()

	// Add middleware
	router.Use(middleware.RequestID())
	router.Use(gin.Recovery())

	// Add rate limiting (100 requests per minute per IP)
	rateLimiter := middleware.NewRateLimiter(100, 10)
	router.Use(rateLimiter.Middleware())

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Configure based on your needs
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Add optimized request logging
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Only log if duration is significant or there's an error
		if duration > 100*time.Millisecond || c.Writer.Status() >= 400 {
			logger.WithFields(map[string]interface{}{
				"method":      c.Request.Method,
				"path":        c.Request.URL.Path,
				"status_code": c.Writer.Status(),
				"duration":    duration.String(),
				"client_ip":   c.ClientIP(),
				"request_id":  c.GetString("request_id"),
			}).Info("HTTP request processed")
		}
	})

	// Setup routes
	setupOptimizedRoutes(router, cfg, accessTokenManager, vaService, mpmService, logger)

	// Create HTTP server with optimizations
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Start cache cleanup routine
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				accessTokenManager.CleanupCache()
				rateLimiter.Cleanup()
			}
		}
	}()

	// Start server in a goroutine
	go func() {
		logger.WithFields(map[string]interface{}{
			"host": cfg.Server.Host,
			"port": cfg.Server.Port,
		}).Info("Starting HTTP server")

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

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shutdown")
	}

	logger.Info("Server gracefully stopped")
}

// setupOptimizedRoutes configures the application routes with optimizations
func setupOptimizedRoutes(
	router *gin.Engine,
	cfg *config.Config,
	accessTokenManager *auth.OptimizedAccessTokenManager,
	vaService *services.VirtualAccountService,
	mpmService *services.MPMService,
	logger logging.Logger,
) {
	// Health check endpoint with caching
	router.GET("/health", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=60")
		c.JSON(http.StatusOK, gin.H{
			"status":      "healthy",
			"app_name":    cfg.App.Name,
			"app_version": cfg.App.Version,
			"timestamp":   time.Now().Format(time.RFC3339),
			"uptime":      time.Since(time.Now()).String(),
		})
	})

	// API version 1
	v1 := router.Group("/api/v1")
	{
		// Authentication endpoints
		auth := v1.Group("/auth")
		{
			// Get B2B access token
			auth.POST("/b2b-token", func(c *gin.Context) {
				handleGetB2BTokenOptimized(c, accessTokenManager, logger)
			})

			// Get B2B2C access token
			auth.POST("/b2b2c-token", func(c *gin.Context) {
				handleGetB2B2CTokenOptimized(c, accessTokenManager, logger)
			})
		}

		// Virtual Account endpoints
		va := v1.Group("/virtual-account")
		{
			va.POST("/create", func(c *gin.Context) {
				handleCreateVAOptimized(c, vaService, logger)
			})

			va.PUT("/update", func(c *gin.Context) {
				handleUpdateVAOptimized(c, vaService, logger)
			})

			va.DELETE("/delete", func(c *gin.Context) {
				handleDeleteVAOptimized(c, vaService, logger)
			})

			va.POST("/inquiry", func(c *gin.Context) {
				handleInquiryVAOptimized(c, vaService, logger)
			})

			va.POST("/inquiry-general", func(c *gin.Context) {
				handleInquiryOptimized(c, vaService, logger)
			})

			va.POST("/payment", func(c *gin.Context) {
				handlePaymentOptimized(c, vaService, logger)
			})

			va.POST("/status", func(c *gin.Context) {
				handleStatusOptimized(c, vaService, logger)
			})

			va.POST("/report", func(c *gin.Context) {
				handleReportOptimized(c, vaService, logger)
			})

			va.PUT("/status", func(c *gin.Context) {
				handleUpdateStatusOptimized(c, vaService, logger)
			})
		}

		// MPM Transfer Credit endpoints
		mpm := v1.Group("/mpm")
		{
			mpm.POST("/transfer", func(c *gin.Context) {
				handleMPMTransferOptimized(c, mpmService, logger)
			})

			mpm.POST("/inquiry", func(c *gin.Context) {
				handleMPMInquiryOptimized(c, mpmService, logger)
			})

			mpm.POST("/status", func(c *gin.Context) {
				handleMPMStatusOptimized(c, mpmService, logger)
			})

			mpm.POST("/refund", func(c *gin.Context) {
				handleMPMRefundOptimized(c, mpmService, logger)
			})

			mpm.POST("/balance-inquiry", func(c *gin.Context) {
				handleMPMBalanceInquiryOptimized(c, mpmService, logger)
			})

			mpm.POST("/account-inquiry", func(c *gin.Context) {
				handleMPMAccountInquiryOptimized(c, mpmService, logger)
			})

			mpm.POST("/history", func(c *gin.Context) {
				handleMPMHistoryOptimized(c, mpmService, logger)
			})

			mpm.POST("/generate-qr", func(c *gin.Context) {
				handleMPMGenerateQROptimized(c, mpmService, logger)
			})

			mpm.POST("/notify-qr", func(c *gin.Context) {
				handleMPMNotifyQROptimized(c, mpmService, logger)
			})
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

// Optimized handler functions with better error handling and logging

func handleGetB2BTokenOptimized(c *gin.Context, accessTokenManager *auth.OptimizedAccessTokenManager, logger logging.Logger) {
	requestID := c.GetString("request_id")
	
	tokenResponse, err := accessTokenManager.GetAccessToken(c.Request.Context())
	if err != nil {
		logger.WithFields(map[string]interface{}{
			"request_id": requestID,
			"error":      err.Error(),
		}).Error("Failed to get B2B access token")

		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

func handleGetB2B2CTokenOptimized(c *gin.Context, accessTokenManager *auth.OptimizedAccessTokenManager, logger logging.Logger) {
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

		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

func handleCreateVAOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
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
		
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Add similar optimized handlers for other endpoints...

func handleErrorResponse(c *gin.Context, err error) {
	requestID := c.GetString("request_id")
	
	switch e := err.(type) {
	case *snapErrors.Error:
		statusCode := e.GetHTTPStatus()
		c.JSON(statusCode, gin.H{
			"error":      e.Type,
			"message":    e.Message,
			"code":       e.Code,
			"request_id": requestID,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Internal error",
			"message":    "An unexpected error occurred",
			"request_id": requestID,
		})
	}
}

// Add placeholder handlers for other endpoints...
func handleUpdateVAOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)           {}
func handleDeleteVAOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)           {}
func handleInquiryVAOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)          {}
func handleInquiryOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)            {}
func handlePaymentOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)            {}
func handleStatusOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)             {}
func handleReportOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)             {}
func handleUpdateStatusOptimized(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger)       {}
func handleMPMTransferOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                  {}
func handleMPMInquiryOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                   {}
func handleMPMStatusOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                    {}
func handleMPMRefundOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                    {}
func handleMPMBalanceInquiryOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)            {}
func handleMPMAccountInquiryOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)            {}
func handleMPMHistoryOptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                   {}
func handleMPMGenerateQROptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                {}
func handleMPMNotifyQROptimized(c *gin.Context, mpmService *services.MPMService, logger logging.Logger)                  {}