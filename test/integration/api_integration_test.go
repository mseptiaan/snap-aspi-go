package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestServer represents a test server instance
type TestServer struct {
	server  *httptest.Server
	config  *config.Config
	logger  logging.Logger
	client  *http.Client
	baseURL string
}

// NewTestServer creates a new test server for integration testing
func NewTestServer(t *testing.T) *TestServer {
	// Create test configuration
	cfg := &config.Config{
		Server: config.ServerConfig{
			Host: "localhost",
			Port: "0", // Let the test server choose a port
		},
		ASPI: config.ASPIConfig{
			ClientID:       "test-client-id",
			ClientSecret:   "test-client-secret",
			BaseURL:        "https://sandbox.aspi-indonesia.or.id",
			PrivateKeyPath: "../../keys/private_key.pem",
			PublicKeyPath:  "../../keys/public_key.pem",
		},
		Logger: config.LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	// Create logger
	logger := logging.NewLogger(&cfg.Logger)

	// Create HTTP client
	httpClient := client.NewSnapClient(cfg, logger)

	// Create auth manager
	authManager, err := auth.NewAccessTokenManager(cfg, httpClient, logger)
	require.NoError(t, err)

	// Create Virtual Account service
	vaService, err := services.NewVirtualAccountService(httpClient, authManager, cfg, logger)
	require.NoError(t, err)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Setup routes
	setupTestRoutes(router, vaService, authManager, cfg, logger)

	// Create test server
	server := httptest.NewServer(router)

	return &TestServer{
		server:  server,
		config:  cfg,
		logger:  logger,
		client:  &http.Client{Timeout: 30 * time.Second},
		baseURL: server.URL,
	}
}

// Close shuts down the test server
func (ts *TestServer) Close() {
	ts.server.Close()
}

// setupTestRoutes configures the test routes
func setupTestRoutes(
	router *gin.Engine,
	vaService *services.VirtualAccountService,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	api := router.Group("/api/v1")

	// Auth endpoints
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/b2b-token", func(c *gin.Context) {
			ctx := context.Background()
			token, err := authManager.GetAccessToken(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, token)
		})

		authGroup.POST("/b2b2c-token", func(c *gin.Context) {
			var req auth.CustomerTokenRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			ctx := context.Background()
			token, err := authManager.GetCustomerAccessToken(ctx, req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, token)
		})
	}

	// Virtual Account endpoints
	va := api.Group("/virtual-account")
	{
		va.POST("/create", func(c *gin.Context) {
			var payload types.CreateVAPayload
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			ctx := context.Background()
			result, err := vaService.CreateVA(ctx, &payload)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})

		va.POST("/inquiry", func(c *gin.Context) {
			var payload types.InquiryVAPayload
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			ctx := context.Background()
			result, err := vaService.InquiryVA(ctx, &payload)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
}

func TestHealthEndpoint(t *testing.T) {
	server := NewTestServer(t)
	defer server.Close()

	resp, err := server.client.Get(server.baseURL + "/health")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "healthy", response["status"])
	assert.Contains(t, response, "timestamp")
}

func TestCreateVirtualAccountEndpoint(t *testing.T) {
	// Skip if no private key file exists
	if _, err := os.Stat("../../keys/private_key.pem"); os.IsNotExist(err) {
		t.Skip("Private key file not found, skipping integration test")
	}

	server := NewTestServer(t)
	defer server.Close()

	tests := []struct {
		name           string
		payload        types.CreateVAPayload
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid_create_va_request",
			payload: types.CreateVAPayload{
				PartnerServiceId:    "test-service",
				CustomerNo:          "123456789",
				VirtualAccountNo:    "8808001234567890",
				VirtualAccountName:  "Test Account",
				TrxId:               "TRX-001",
				TotalAmount:         &types.Amount{Value: "100000.00", Currency: "IDR"},
				VirtualAccountEmail: "test@example.com",
				VirtualAccountPhone: "081234567890",
				ExpiredDate:         "2024-12-31T23:59:59+07:00",
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "invalid_create_va_request_missing_customer_no",
			payload: types.CreateVAPayload{
				PartnerServiceId:   "test-service",
				VirtualAccountNo:   "8808001234567890",
				VirtualAccountName: "Test Account",
				TrxId:              "TRX-001",
				TotalAmount:        &types.Amount{Value: "100000.00", Currency: "IDR"},
			},
			expectedStatus: http.StatusInternalServerError,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal payload to JSON
			payloadBytes, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			// Make request
			resp, err := server.client.Post(
				server.baseURL+"/api/v1/virtual-account/create",
				"application/json",
				bytes.NewBuffer(payloadBytes),
			)
			require.NoError(t, err)
			defer resp.Body.Close()

			// Check status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Decode response
			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			require.NoError(t, err)

			if tt.expectError {
				assert.Contains(t, response, "error")
			} else {
				// For successful responses, we expect some response data
				// Note: This might fail if the external API is not available
				// In a real integration test, you might want to mock the external API
				assert.NotEmpty(t, response)
			}
		})
	}
}

func TestVirtualAccountInquiryEndpoint(t *testing.T) {
	// Skip if no private key file exists
	if _, err := os.Stat("../../keys/private_key.pem"); os.IsNotExist(err) {
		t.Skip("Private key file not found, skipping integration test")
	}

	server := NewTestServer(t)
	defer server.Close()

	payload := types.InquiryVAPayload{
		PartnerServiceId: "test-service",
		CustomerNo:       "123456789",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      "2024-01-01",
	}

	payloadBytes, err := json.Marshal(payload)
	require.NoError(t, err)

	resp, err := server.client.Post(
		server.baseURL+"/api/v1/virtual-account/inquiry",
		"application/json",
		bytes.NewBuffer(payloadBytes),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// We expect this to return some response (might be an error from external API)
	assert.True(t, resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusInternalServerError)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.NotEmpty(t, response)
}

// Benchmark test for API performance
func BenchmarkCreateVAEndpoint(b *testing.B) {
	// Skip if no private key file exists
	if _, err := os.Stat("../../keys/private_key.pem"); os.IsNotExist(err) {
		b.Skip("Private key file not found, skipping benchmark")
	}

	server := NewTestServer(&testing.T{})
	defer server.Close()

	payload := types.CreateVAPayload{
		PartnerServiceId:   "test-service",
		CustomerNo:         "123456789",
		VirtualAccountNo:   "8808001234567890",
		VirtualAccountName: "Test Account",
		TrxId:              "TRX-001",
		TotalAmount:        &types.Amount{Value: "100000.00", Currency: "IDR"},
	}

	payloadBytes, _ := json.Marshal(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := server.client.Post(
			server.baseURL+"/api/v1/virtual-account/create",
			"application/json",
			bytes.NewBuffer(payloadBytes),
		)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
