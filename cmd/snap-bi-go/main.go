package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	snapErrors "github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
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
	}).Info("Starting Snap ASPI Go application")

	// Initialize HTTP client
	httpClient := client.NewSnapClient(cfg, logger)

	// Initialize access token manager
	accessTokenManager, err := auth.NewAccessTokenManager(cfg, httpClient, logger)
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

	// Create Gin router with middleware
	router := gin.New()

	// Add logging middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Add recovery middleware
	router.Use(gin.Recovery())

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Configure based on your needs
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	// Add request logging
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		logger.WithFields(map[string]interface{}{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status_code": c.Writer.Status(),
			"duration":    duration.String(),
			"client_ip":   c.ClientIP(),
			"user_agent":  c.Request.UserAgent(),
		}).Info("HTTP request processed")
	})

	// Setup routes
	setupRoutes(router, cfg, accessTokenManager, vaService, mpmService, logger)

	// Create HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

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

// setupRoutes configures the application routes
func setupRoutes(
	router *gin.Engine,
	cfg *config.Config,
	accessTokenManager *auth.AccessTokenManager,
	vaService *services.VirtualAccountService,
	mpmService *services.MPMService,
	logger logging.Logger,
) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "healthy",
			"app_name":    cfg.App.Name,
			"app_version": cfg.App.Version,
			"timestamp":   time.Now().Format(time.RFC3339),
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
				handleGetB2BToken(c, accessTokenManager, logger)
			})

			// Get B2B2C access token
			auth.POST("/b2b2c-token", func(c *gin.Context) {
				handleGetB2B2CToken(c, accessTokenManager, logger)
			})
		}

		// Virtual Account endpoints
		va := v1.Group("/virtual-account")
		{
			// Create Virtual Account
			va.POST("/create", func(c *gin.Context) {
				handleCreateVA(c, vaService, logger)
			})

			// Update Virtual Account
			va.PUT("/update", func(c *gin.Context) {
				handleUpdateVA(c, vaService, logger)
			})

			// Delete Virtual Account
			va.DELETE("/delete", func(c *gin.Context) {
				handleDeleteVA(c, vaService, logger)
			})

			// Virtual Account inquiry
			va.POST("/inquiry", func(c *gin.Context) {
				handleInquiryVA(c, vaService, logger)
			})

			// General inquiry
			va.POST("/inquiry-general", func(c *gin.Context) {
				handleInquiry(c, vaService, logger)
			})

			// Payment
			va.POST("/payment", func(c *gin.Context) {
				handlePayment(c, vaService, logger)
			})

			// Status check
			va.POST("/status", func(c *gin.Context) {
				handleStatus(c, vaService, logger)
			})

			// Report generation
			va.POST("/report", func(c *gin.Context) {
				handleReport(c, vaService, logger)
			})

			// Update status
			va.PUT("/status", func(c *gin.Context) {
				handleUpdateStatus(c, vaService, logger)
			})
		}

		// MPM Transfer Credit endpoints
		mpm := v1.Group("/mpm")
		{
			// Transfer Credit
			mpm.POST("/transfer", func(c *gin.Context) {
				handleMPMTransfer(c, mpmService, logger)
			})

			// Transfer Inquiry
			mpm.POST("/inquiry", func(c *gin.Context) {
				handleMPMInquiry(c, mpmService, logger)
			})

			// Transfer Status
			mpm.POST("/status", func(c *gin.Context) {
				handleMPMStatus(c, mpmService, logger)
			})

			// Transfer Refund
			mpm.POST("/refund", func(c *gin.Context) {
				handleMPMRefund(c, mpmService, logger)
			})

			// Balance Inquiry
			mpm.POST("/balance-inquiry", func(c *gin.Context) {
				handleMPMBalanceInquiry(c, mpmService, logger)
			})

			// Account Inquiry
			mpm.POST("/account-inquiry", func(c *gin.Context) {
				handleMPMAccountInquiry(c, mpmService, logger)
			})

			// Transaction History
			mpm.POST("/history", func(c *gin.Context) {
				handleMPMHistory(c, mpmService, logger)
			})

			// QR MPM endpoints
			mpm.POST("/generate-qr", func(c *gin.Context) {
				handleMPMGenerateQR(c, mpmService, logger)
			})

			mpm.POST("/notify-qr", func(c *gin.Context) {
				handleMPMNotifyQR(c, mpmService, logger)
			})
		}

		// Utility endpoints
		utils := v1.Group("/utils")
		{
			// Configuration info (non-sensitive)
			utils.GET("/config", func(c *gin.Context) {
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

// handleGetB2BToken handles B2B access token requests
func handleGetB2BToken(c *gin.Context, accessTokenManager *auth.AccessTokenManager, logger logging.Logger) {
	// Get access token
	tokenResponse, err := accessTokenManager.GetAccessToken(c.Request.Context())
	if err != nil {
		logger.WithError(err).Error("Failed to get B2B access token")

		// Handle different error types
		switch e := err.(type) {
		case *snapErrors.Error:
			switch e.Type {
			case snapErrors.ErrorTypeConfiguration:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Configuration error",
					"message": "Service configuration is invalid",
				})
			case snapErrors.ErrorTypeExternal:
				c.JSON(http.StatusBadGateway, gin.H{
					"error":   "External service error",
					"message": "Failed to communicate with ASPI API",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal error",
					"message": "An unexpected error occurred",
				})
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal error",
				"message": "An unexpected error occurred",
			})
		}
		return
	}

	// Return token response
	c.JSON(http.StatusOK, tokenResponse)
}

// handleGetB2B2CToken handles B2B2C access token requests
func handleGetB2B2CToken(c *gin.Context, accessTokenManager *auth.AccessTokenManager, logger logging.Logger) {
	// Parse request body
	var request auth.CustomerTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.WithError(err).Warn("Invalid B2B2C token request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	// Get customer access token
	tokenResponse, err := accessTokenManager.GetCustomerAccessToken(c.Request.Context(), request)
	if err != nil {
		logger.WithError(err).Error("Failed to get B2B2C access token")

		// Handle different error types
		switch e := err.(type) {
		case *snapErrors.Error:
			switch e.Type {
			case snapErrors.ErrorTypeValidation:
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Validation error",
					"message": e.Message,
				})
			case snapErrors.ErrorTypeConfiguration:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Configuration error",
					"message": "Service configuration is invalid",
				})
			case snapErrors.ErrorTypeExternal:
				c.JSON(http.StatusBadGateway, gin.H{
					"error":   "External service error",
					"message": "Failed to communicate with ASPI API",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal error",
					"message": "An unexpected error occurred",
				})
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal error",
				"message": "An unexpected error occurred",
			})
		}
		return
	}

	// Return token response
	c.JSON(http.StatusOK, tokenResponse)
}

// Virtual Account endpoint handlers

// handleCreateVA handles Virtual Account creation requests
func handleCreateVA(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.CreateVAPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid create VA request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.CreateVA(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to create Virtual Account")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleUpdateVA handles Virtual Account update requests
func handleUpdateVA(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.CreateVAPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid update VA request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.UpdateVA(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to update Virtual Account")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleDeleteVA handles Virtual Account deletion requests
func handleDeleteVA(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.InquiryVAPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid delete VA request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.DeleteVA(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to delete Virtual Account")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleInquiryVA handles Virtual Account inquiry requests
func handleInquiryVA(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.InquiryVAPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid inquiry VA request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.InquiryVA(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to inquiry Virtual Account")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleInquiry handles general inquiry requests
func handleInquiry(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.InquiryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid inquiry request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.Inquiry(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process inquiry")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handlePayment handles Virtual Account payment requests
func handlePayment(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.PaymentPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid payment request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.Payment(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process payment")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleStatus handles status check requests
func handleStatus(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.StatusPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid status request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.Status(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to check status")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleReport handles report generation requests
func handleReport(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.ReportPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid report request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.Report(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to generate report")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleUpdateStatus handles status update requests
func handleUpdateStatus(c *gin.Context, vaService *services.VirtualAccountService, logger logging.Logger) {
	var payload types.StatusPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid update status request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := vaService.UpdateStatus(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to update status")
		handleVAError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleVAError handles Virtual Account service errors
func handleVAError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *snapErrors.Error:
		switch e.Type {
		case snapErrors.ErrorTypeValidation:
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Validation error",
				"message": e.Message,
			})
		case snapErrors.ErrorTypeAuthentication:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Authentication error",
				"message": "Authentication failed",
			})
		case snapErrors.ErrorTypeConfiguration:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Configuration error",
				"message": "Service configuration is invalid",
			})
		case snapErrors.ErrorTypeExternal:
			c.JSON(http.StatusBadGateway, gin.H{
				"error":   "External service error",
				"message": "Failed to communicate with ASPI API",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal error",
				"message": "An unexpected error occurred",
			})
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal error",
			"message": "An unexpected error occurred",
		})
	}
}

// MPM Transfer Credit endpoint handlers

// handleMPMTransfer handles MPM transfer credit requests
func handleMPMTransfer(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMTransferPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM transfer request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.Transfer(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM transfer")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMInquiry handles MPM transfer inquiry requests
func handleMPMInquiry(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMInquiryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM inquiry request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.Inquiry(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM inquiry")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMStatus handles MPM transfer status requests
func handleMPMStatus(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMStatusPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM status request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.Status(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to check MPM status")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMRefund handles MPM transfer refund requests
func handleMPMRefund(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMRefundPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM refund request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.Refund(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM refund")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMBalanceInquiry handles MPM balance inquiry requests
func handleMPMBalanceInquiry(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMBalanceInquiryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM balance inquiry request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.BalanceInquiry(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM balance inquiry")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMAccountInquiry handles MPM account inquiry requests
func handleMPMAccountInquiry(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMAccountInquiryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM account inquiry request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.AccountInquiry(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM account inquiry")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMHistory handles MPM transaction history requests
func handleMPMHistory(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMHistoryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM history request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.History(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to retrieve MPM history")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMError handles MPM service errors
func handleMPMError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *snapErrors.Error:
		switch e.Type {
		case snapErrors.ErrorTypeValidation:
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Validation error",
				"message": e.Message,
			})
		case snapErrors.ErrorTypeAuthentication:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Authentication error",
				"message": "Authentication failed",
			})
		case snapErrors.ErrorTypeConfiguration:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Configuration error",
				"message": "Service configuration is invalid",
			})
		case snapErrors.ErrorTypeExternal:
			c.JSON(http.StatusBadGateway, gin.H{
				"error":   "External service error",
				"message": "Failed to communicate with ASPI API",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal error",
				"message": "An unexpected error occurred",
			})
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal error",
			"message": "An unexpected error occurred",
		})
	}
}

// handleMPMGenerateQR handles MPM QR generation requests
func handleMPMGenerateQR(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMGenerateQRPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM generate QR request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.GenerateQR(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to generate MPM QR")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleMPMNotifyQR handles MPM QR notification requests
func handleMPMNotifyQR(c *gin.Context, mpmService *services.MPMService, logger logging.Logger) {
	var payload types.MPMNotifyQRPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.WithError(err).Warn("Invalid MPM notify QR request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Request body is invalid",
		})
		return
	}

	result, err := mpmService.NotifyQR(c.Request.Context(), &payload)
	if err != nil {
		logger.WithError(err).Error("Failed to process MPM QR notification")
		handleMPMError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
