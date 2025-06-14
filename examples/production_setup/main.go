package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// ProductionApp represents a production-ready application using SNAP SDK
type ProductionApp struct {
	snapClient *snap.Client
	logger     *log.Logger
	config     *ProductionConfig
	shutdown   chan os.Signal
	wg         sync.WaitGroup
}

// ProductionConfig holds production configuration
type ProductionConfig struct {
	// ASPI Configuration
	ASPIBaseURL        string
	ASPIClientID       string
	ASPIClientSecret   string
	ASPIPrivateKeyPath string
	ASPIPublicKeyPath  string
	ASPIEnvironment    string

	// Application Configuration
	AppName        string
	AppVersion     string
	LogLevel       string
	HealthCheckURL string

	// Performance Configuration
	MaxConcurrentRequests int
	RequestTimeout        int
	MaxRetries            int
	BackoffDuration       int

	// Security Configuration
	TLSMinVersion uint16
	TLSCipherSuites []uint16
}

func main() {
	fmt.Println("=== SNAP ASPI SDK Production Setup ===")

	// Load production configuration
	config := loadProductionConfig()

	// Initialize production application
	app, err := NewProductionApp(config)
	if err != nil {
		log.Fatalf("Failed to initialize production app: %v", err)
	}

	// Start the application
	app.Start()

	// Wait for shutdown signal
	app.WaitForShutdown()

	fmt.Println("Application shutdown complete")
}

// NewProductionApp creates a new production application instance
func NewProductionApp(config *ProductionConfig) (*ProductionApp, error) {
	// Initialize logger
	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", config.AppName), log.LstdFlags|log.Lshortfile)

	// Initialize SNAP client with production settings
	snapClient, err := snap.NewClient(snap.Config{
		BaseURL:         config.ASPIBaseURL,
		ClientID:        config.ASPIClientID,
		ClientSecret:    config.ASPIClientSecret,
		PrivateKeyPath:  config.ASPIPrivateKeyPath,
		PublicKeyPath:   config.ASPIPublicKeyPath,
		Environment:     config.ASPIEnvironment,
		ConnectTimeout:  5,  // Fast connection timeout for production
		RequestTimeout:  config.RequestTimeout,
		MaxRetries:      config.MaxRetries,
		BackoffDuration: config.BackoffDuration,
		LogLevel:        config.LogLevel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize SNAP client: %w", err)
	}

	// Setup shutdown channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	return &ProductionApp{
		snapClient: snapClient,
		logger:     logger,
		config:     config,
		shutdown:   shutdown,
	}, nil
}

// Start starts the production application
func (app *ProductionApp) Start() {
	app.logger.Printf("Starting %s v%s in %s mode", 
		app.config.AppName, 
		app.config.AppVersion, 
		app.config.ASPIEnvironment)

	// Start health check service
	app.wg.Add(1)
	go app.startHealthCheckService()

	// Start background services
	app.wg.Add(1)
	go app.startBackgroundServices()

	// Start metrics collection
	app.wg.Add(1)
	go app.startMetricsCollection()

	// Perform initial health check
	app.performInitialHealthCheck()

	app.logger.Println("Application started successfully")
}

// WaitForShutdown waits for shutdown signal and performs graceful shutdown
func (app *ProductionApp) WaitForShutdown() {
	<-app.shutdown
	app.logger.Println("Shutdown signal received, starting graceful shutdown...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Perform graceful shutdown
	app.gracefulShutdown(ctx)

	// Wait for all goroutines to finish
	done := make(chan struct{})
	go func() {
		app.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		app.logger.Println("Graceful shutdown completed")
	case <-ctx.Done():
		app.logger.Println("Shutdown timeout exceeded, forcing exit")
	}
}

// startHealthCheckService starts the health check HTTP service
func (app *ProductionApp) startHealthCheckService() {
	defer app.wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", app.healthCheckHandler)
	mux.HandleFunc("/ready", app.readinessHandler)
	mux.HandleFunc("/metrics", app.metricsHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion:   app.config.TLSMinVersion,
			CipherSuites: app.config.TLSCipherSuites,
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	app.logger.Println("Health check service started on :8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.logger.Printf("Health check service error: %v", err)
	}
}

// startBackgroundServices starts background processing services
func (app *ProductionApp) startBackgroundServices() {
	defer app.wg.Done()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			app.performPeriodicTasks()
		case <-app.shutdown:
			app.logger.Println("Background services shutting down")
			return
		}
	}
}

// startMetricsCollection starts metrics collection
func (app *ProductionApp) startMetricsCollection() {
	defer app.wg.Done()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			app.collectMetrics()
		case <-app.shutdown:
			app.logger.Println("Metrics collection shutting down")
			return
		}
	}
}

// performInitialHealthCheck performs initial health check
func (app *ProductionApp) performInitialHealthCheck() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.logger.Println("Performing initial health check...")

	// Test SNAP API connectivity
	_, err := app.snapClient.Auth().GetB2BToken(ctx)
	if err != nil {
		app.logger.Printf("WARNING: Initial SNAP API health check failed: %v", err)
	} else {
		app.logger.Println("SNAP API connectivity verified")
	}
}

// performPeriodicTasks performs periodic maintenance tasks
func (app *ProductionApp) performPeriodicTasks() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	app.logger.Println("Performing periodic tasks...")

	// Example: Verify API connectivity
	_, err := app.snapClient.Auth().GetB2BToken(ctx)
	if err != nil {
		app.logger.Printf("Periodic API check failed: %v", err)
	}

	// Example: Process pending transactions
	app.processPendingTransactions(ctx)

	// Example: Cleanup expired data
	app.cleanupExpiredData()
}

// processPendingTransactions processes any pending transactions
func (app *ProductionApp) processPendingTransactions(ctx context.Context) {
	// Example implementation
	app.logger.Println("Processing pending transactions...")

	// This would typically involve:
	// 1. Querying database for pending transactions
	// 2. Checking status with SNAP API
	// 3. Updating transaction status
	// 4. Sending notifications if needed

	// Example status check
	statusPayload := &types.StatusPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		InquiryRequestId: "INQ-001",
	}

	_, err := app.snapClient.VirtualAccount().Status(ctx, statusPayload)
	if err != nil {
		app.logger.Printf("Status check failed: %v", err)
	}
}

// cleanupExpiredData cleans up expired data
func (app *ProductionApp) cleanupExpiredData() {
	app.logger.Println("Cleaning up expired data...")
	// Implementation would clean up expired tokens, cache entries, etc.
}

// collectMetrics collects application metrics
func (app *ProductionApp) collectMetrics() {
	// Example metrics collection
	// In production, you would use proper metrics libraries like Prometheus
	app.logger.Println("Collecting metrics...")
}

// gracefulShutdown performs graceful shutdown
func (app *ProductionApp) gracefulShutdown(ctx context.Context) {
	app.logger.Println("Performing graceful shutdown...")

	// Stop accepting new requests
	// Complete ongoing operations
	// Close database connections
	// Flush logs
	// etc.
}

// HTTP Handlers

func (app *ProductionApp) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Check SNAP API health
	_, err := app.snapClient.Auth().GetB2BToken(ctx)
	
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"status":"unhealthy","error":"%v","timestamp":"%s"}`, 
			err, time.Now().Format(time.RFC3339))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"healthy","timestamp":"%s","version":"%s"}`, 
		time.Now().Format(time.RFC3339), app.config.AppVersion)
}

func (app *ProductionApp) readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Check if application is ready to serve requests
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ready","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
}

func (app *ProductionApp) metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Return application metrics
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "# Application metrics\n")
	fmt.Fprintf(w, "app_uptime_seconds %d\n", int(time.Since(time.Now()).Seconds()))
	// Add more metrics as needed
}

// loadProductionConfig loads production configuration from environment
func loadProductionConfig() *ProductionConfig {
	return &ProductionConfig{
		// ASPI Configuration
		ASPIBaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://api.aspi-indonesia.or.id"),
		ASPIClientID:       getEnvOrDefault("ASPI_CLIENT_ID", ""),
		ASPIClientSecret:   getEnvOrDefault("ASPI_CLIENT_SECRET", ""),
		ASPIPrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "/etc/ssl/private/aspi_private.pem"),
		ASPIPublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "/etc/ssl/certs/aspi_public.pem"),
		ASPIEnvironment:    getEnvOrDefault("ASPI_ENVIRONMENT", "production"),

		// Application Configuration
		AppName:        getEnvOrDefault("APP_NAME", "snap-aspi-production"),
		AppVersion:     getEnvOrDefault("APP_VERSION", "1.0.0"),
		LogLevel:       getEnvOrDefault("LOG_LEVEL", "warn"),
		HealthCheckURL: getEnvOrDefault("HEALTH_CHECK_URL", "http://localhost:8080/health"),

		// Performance Configuration
		MaxConcurrentRequests: getEnvAsInt("MAX_CONCURRENT_REQUESTS", 100),
		RequestTimeout:        getEnvAsInt("REQUEST_TIMEOUT", 15),
		MaxRetries:            getEnvAsInt("MAX_RETRIES", 2),
		BackoffDuration:       getEnvAsInt("BACKOFF_DURATION", 1),

		// Security Configuration
		TLSMinVersion: tls.VersionTLS12,
		TLSCipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
}

// Helper functions

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		// In production, you'd want proper error handling here
		return defaultValue
	}
	return defaultValue
}

// Example business logic functions

func (app *ProductionApp) ProcessVirtualAccountPayment(ctx context.Context, payload *types.PaymentPayload) error {
	// Add request ID for tracing
	requestID := generateRequestID()
	app.logger.Printf("Processing VA payment [%s]: %s", requestID, payload.VirtualAccountNo)

	// Process payment with retry logic
	var result map[string]any
	var err error

	for attempt := 1; attempt <= app.config.MaxRetries; attempt++ {
		result, err = app.snapClient.VirtualAccount().Payment(ctx, payload)
		if err == nil {
			break
		}

		if attempt < app.config.MaxRetries {
			app.logger.Printf("Payment attempt %d failed [%s]: %v", attempt, requestID, err)
			time.Sleep(time.Duration(attempt) * time.Second)
		}
	}

	if err != nil {
		app.logger.Printf("Payment failed after %d attempts [%s]: %v", app.config.MaxRetries, requestID, err)
		return err
	}

	app.logger.Printf("Payment successful [%s]: %+v", requestID, result)
	return nil
}

func (app *ProductionApp) ProcessMPMTransfer(ctx context.Context, payload *types.MPMTransferPayload) error {
	requestID := generateRequestID()
	app.logger.Printf("Processing MPM transfer [%s]: %s", requestID, payload.PartnerReferenceNo)

	result, err := app.snapClient.MPM().Transfer(ctx, payload)
	if err != nil {
		app.logger.Printf("MPM transfer failed [%s]: %v", requestID, err)
		return err
	}

	app.logger.Printf("MPM transfer successful [%s]: %+v", requestID, result)
	return nil
}

func generateRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}