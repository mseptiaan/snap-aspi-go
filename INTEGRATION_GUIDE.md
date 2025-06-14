# SNAP ASPI Go SDK Integration Guide

This guide will help you integrate the SNAP ASPI Go SDK into your application for seamless payment processing with Indonesia's ASPI (Payment System Administration) services.

## 📋 Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Basic Usage](#basic-usage)
5. [Virtual Account Operations](#virtual-account-operations)
6. [MPM (Merchant Payment Management)](#mpm-merchant-payment-management)
7. [QR Code Operations](#qr-code-operations)
8. [Authentication](#authentication)
9. [Error Handling](#error-handling)
10. [Best Practices](#best-practices)
11. [Testing](#testing)
12. [Production Deployment](#production-deployment)

## 🔧 Prerequisites

Before integrating the SDK, ensure you have:

- **Go 1.23 or later**
- **ASPI API credentials** (Client ID, Client Secret)
- **RSA Key Pair** (Private and Public keys for signing)
- **Valid ASPI account** with appropriate permissions

## 📦 Installation

### Step 1: Add the SDK to your project

```bash
go mod init your-project-name
go get github.com/mseptiaan/snap-aspi-go
```

### Step 2: Import the SDK

```go
import (
    "github.com/mseptiaan/snap-aspi-go/pkg/snap"
    "github.com/mseptiaan/snap-aspi-go/pkg/types"
)
```

## ⚙️ Configuration

### Step 1: Prepare RSA Keys

Generate RSA key pair for API signing:

```bash
# Generate private key
openssl genrsa -out private_key.pem 2048

# Generate public key
openssl rsa -in private_key.pem -pubout -out public_key.pem
```

### Step 2: Initialize the SDK Client

```go
package main

import (
    "context"
    "log"
    
    "github.com/mseptiaan/snap-aspi-go/pkg/snap"
)

func main() {
    // Initialize SDK client
    client, err := snap.NewClient(snap.Config{
        BaseURL:        "https://sandbox.aspi-indonesia.or.id", // Use production URL for live
        ClientID:       "your-client-id",
        ClientSecret:   "your-client-secret",
        PrivateKeyPath: "path/to/private_key.pem",
        PublicKeyPath:  "path/to/public_key.pem",
        Environment:    "sandbox", // "sandbox" or "production"
        
        // Optional configurations
        ConnectTimeout:  10, // seconds
        RequestTimeout:  30, // seconds
        MaxRetries:      3,
        BackoffDuration: 1,  // seconds
        LogLevel:        "info", // debug, info, warn, error
    })
    if err != nil {
        log.Fatalf("Failed to initialize SNAP client: %v", err)
    }
    
    // Your application logic here
}
```

### Step 3: Environment Configuration

For different environments, use these base URLs:

```go
// Sandbox (Testing)
BaseURL: "https://sandbox.aspi-indonesia.or.id"

// Production
BaseURL: "https://api.aspi-indonesia.or.id"
```

## 🚀 Basic Usage

### Context Management

Always use context for timeout and cancellation control:

```go
ctx := context.Background()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

## 💳 Virtual Account Operations

### Create Virtual Account

```go
func createVirtualAccount(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.CreateVAPayload{
        PartnerServiceId:   "12345",
        CustomerNo:         "CUST001",
        VirtualAccountNo:   "8808001234567890",
        VirtualAccountName: "John Doe",
        VirtualAccountEmail: "john@example.com",
        VirtualAccountPhone: "081234567890",
        TrxId:              "TRX-001",
        TotalAmount:        types.NewAmount(100000, "IDR"),
        ExpiredDate:        "2024-12-31T23:59:59+07:00",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "WEB",
        },
    }
    
    result, err := client.VirtualAccount().CreateVA(ctx, payload)
    if err != nil {
        log.Printf("Error creating VA: %v", err)
        return
    }
    
    fmt.Printf("VA Created: %+v\n", result)
}
```

### Inquiry Virtual Account

```go
func inquiryVirtualAccount(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.InquiryVAPayload{
        PartnerServiceId: "12345",
        CustomerNo:       "CUST001",
        VirtualAccountNo: "8808001234567890",
        TrxDateInit:      "2024-01-01",
    }
    
    result, err := client.VirtualAccount().InquiryVA(ctx, payload)
    if err != nil {
        log.Printf("Error inquiring VA: %v", err)
        return
    }
    
    fmt.Printf("VA Inquiry: %+v\n", result)
}
```

### Process Payment

```go
func processPayment(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.PaymentPayload{
        PartnerServiceId:        "12345",
        CustomerNo:              "CUST001",
        VirtualAccountNo:        "8808001234567890",
        VirtualAccountName:      "John Doe",
        TrxId:                   "TRX-001",
        PaymentRequestId:        "PAY-001",
        ChannelCode:             "6014",
        SourceAccountNo:         "1234567890",
        PaidAmount:              types.NewAmount(100000, "IDR"),
        CumulativePaymentAmount: types.NewAmount(100000, "IDR"),
        TotalAmount:             types.NewAmount(100000, "IDR"),
        TrxDateTime:             "2024-01-15T10:30:00+07:00",
        ReferenceNo:             "REF001",
        PaymentType:             "1",
        FlagAdvise:              "N",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "ATM",
        },
    }
    
    result, err := client.VirtualAccount().Payment(ctx, payload)
    if err != nil {
        log.Printf("Error processing payment: %v", err)
        return
    }
    
    fmt.Printf("Payment Result: %+v\n", result)
}
```

## 🏪 MPM (Merchant Payment Management)

### Transfer Credit

```go
func transferCredit(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.MPMTransferPayload{
        PartnerServiceId:      "12345",
        CustomerNo:            "CUST001",
        PartnerReferenceNo:    "REF001",
        MerchantId:            "MERCHANT001",
        SubMerchantId:         "SUB001",
        ExternalStoreId:       "STORE001",
        Amount:                types.NewAmount(50000, "IDR"),
        BeneficiaryAccountNo:  "1234567890",
        BeneficiaryName:       "Jane Doe",
        BeneficiaryCustomerNo: "CUST002",
        BeneficiaryEmail:      "jane@example.com",
        Currency:              "IDR",
        CustomerReference:     "CUSTREF001",
        FeeType:               "OUR",
        Remark:                "Transfer for payment",
        SourceAccountNo:       "0987654321",
        TransactionDate:       "2024-01-15T10:30:00+07:00",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "MOBILE",
        },
    }
    
    result, err := client.MPM().Transfer(ctx, payload)
    if err != nil {
        log.Printf("Error transferring credit: %v", err)
        return
    }
    
    fmt.Printf("Transfer Result: %+v\n", result)
}
```

### Balance Inquiry

```go
func checkBalance(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.MPMBalanceInquiryPayload{
        PartnerServiceId: "12345",
        CustomerNo:       "CUST001",
        AccountNo:        "1234567890",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "MOBILE",
        },
    }
    
    result, err := client.MPM().BalanceInquiry(ctx, payload)
    if err != nil {
        log.Printf("Error checking balance: %v", err)
        return
    }
    
    fmt.Printf("Balance: %+v\n", result)
}
```

## 📱 QR Code Operations

### Generate QR Code

```go
func generateQRCode(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.MPMGenerateQRPayload{
        PartnerServiceId:   "12345",
        CustomerNo:         "CUST001",
        PartnerReferenceNo: "QR001",
        MerchantId:         "MERCHANT001",
        SubMerchantId:      "SUB001",
        ExternalStoreId:    "STORE001",
        Amount:             types.NewAmount(25000, "IDR"),
        MerchantName:       "Test Merchant",
        MerchantCity:       "Jakarta",
        MerchantPostalCode: "12345",
        TerminalLabel:      "Terminal 1",
        PurposeOfPayment:   "Payment for goods",
        QRType:             "DYNAMIC",
        ValidityPeriod:     "3600", // 1 hour
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "POS",
        },
    }
    
    result, err := client.MPM().GenerateQR(ctx, payload)
    if err != nil {
        log.Printf("Error generating QR: %v", err)
        return
    }
    
    fmt.Printf("QR Code: %s\n", result.QRContent)
    fmt.Printf("Valid Until: %s\n", result.ValidUntil)
}
```

## 🔐 Authentication

### Get B2B Token

```go
func getB2BToken(client *snap.Client) {
    ctx := context.Background()
    
    token, err := client.Auth().GetB2BToken(ctx)
    if err != nil {
        log.Printf("Error getting B2B token: %v", err)
        return
    }
    
    fmt.Printf("Access Token: %s\n", token.AccessToken)
    fmt.Printf("Expires In: %s seconds\n", token.ExpiresIn)
}
```

### Get B2B2C Token

```go
func getB2B2CToken(client *snap.Client) {
    ctx := context.Background()
    
    req := auth.CustomerTokenRequest{
        AuthCode:       "auth-code-from-customer",
        RefreshToken:   "refresh-token",
        AdditionalInfo: map[string]interface{}{
            "customerId": "CUST001",
        },
    }
    
    token, err := client.Auth().GetB2B2CToken(ctx, req)
    if err != nil {
        log.Printf("Error getting B2B2C token: %v", err)
        return
    }
    
    fmt.Printf("Customer Token: %s\n", token.AccessToken)
}
```

## ⚠️ Error Handling

### Comprehensive Error Handling

```go
func handleErrors(client *snap.Client) {
    ctx := context.Background()
    
    payload := &types.CreateVAPayload{
        // ... payload data
    }
    
    result, err := client.VirtualAccount().CreateVA(ctx, payload)
    if err != nil {
        // Check error type
        switch {
        case strings.Contains(err.Error(), "validation"):
            log.Printf("Validation error: %v", err)
            // Handle validation errors
            
        case strings.Contains(err.Error(), "authentication"):
            log.Printf("Authentication error: %v", err)
            // Handle auth errors - maybe refresh token
            
        case strings.Contains(err.Error(), "network"):
            log.Printf("Network error: %v", err)
            // Handle network errors - maybe retry
            
        case strings.Contains(err.Error(), "timeout"):
            log.Printf("Timeout error: %v", err)
            // Handle timeout - maybe increase timeout
            
        default:
            log.Printf("Unknown error: %v", err)
            // Handle other errors
        }
        return
    }
    
    // Process successful result
    fmt.Printf("Success: %+v\n", result)
}
```

### Retry Logic

```go
func withRetry(operation func() error, maxRetries int) error {
    var err error
    for i := 0; i < maxRetries; i++ {
        err = operation()
        if err == nil {
            return nil
        }
        
        // Check if error is retryable
        if strings.Contains(err.Error(), "network") || 
           strings.Contains(err.Error(), "timeout") {
            time.Sleep(time.Duration(i+1) * time.Second)
            continue
        }
        
        // Non-retryable error
        break
    }
    return err
}

// Usage
err := withRetry(func() error {
    _, err := client.VirtualAccount().CreateVA(ctx, payload)
    return err
}, 3)
```

## 🎯 Best Practices

### 1. Configuration Management

```go
// Use environment variables for sensitive data
type Config struct {
    ClientID       string `env:"ASPI_CLIENT_ID"`
    ClientSecret   string `env:"ASPI_CLIENT_SECRET"`
    PrivateKeyPath string `env:"ASPI_PRIVATE_KEY_PATH"`
    Environment    string `env:"ASPI_ENVIRONMENT" envDefault:"sandbox"`
}

// Load from environment
func loadConfig() (*Config, error) {
    cfg := &Config{}
    if err := env.Parse(cfg); err != nil {
        return nil, err
    }
    return cfg, nil
}
```

### 2. Connection Pooling

```go
// The SDK automatically handles connection pooling
// But you can create a singleton client
var (
    snapClient *snap.Client
    once       sync.Once
)

func GetSnapClient() *snap.Client {
    once.Do(func() {
        var err error
        snapClient, err = snap.NewClient(snap.Config{
            // ... configuration
        })
        if err != nil {
            log.Fatalf("Failed to initialize SNAP client: %v", err)
        }
    })
    return snapClient
}
```

### 3. Logging and Monitoring

```go
// Add request ID for tracing
func withRequestID(ctx context.Context) context.Context {
    requestID := generateRequestID()
    return context.WithValue(ctx, "request_id", requestID)
}

// Log operations
func logOperation(operation string, payload interface{}, result interface{}, err error) {
    log.Printf("Operation: %s, Payload: %+v, Result: %+v, Error: %v", 
        operation, payload, result, err)
}
```

### 4. Graceful Shutdown

```go
func gracefulShutdown(client *snap.Client) {
    // The SDK handles cleanup automatically
    // But you can implement graceful shutdown for your app
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    
    <-c
    log.Println("Shutting down gracefully...")
    
    // Cancel ongoing operations
    // Close resources
    // etc.
}
```

## 🧪 Testing

### Unit Testing

```go
func TestVirtualAccountCreation(t *testing.T) {
    // Use test configuration
    client, err := snap.NewClient(snap.Config{
        BaseURL:        "https://sandbox.aspi-indonesia.or.id",
        ClientID:       "test-client-id",
        ClientSecret:   "test-client-secret",
        PrivateKeyPath: "test_keys/private_key.pem",
        PublicKeyPath:  "test_keys/public_key.pem",
        Environment:    "sandbox",
        LogLevel:       "debug",
    })
    require.NoError(t, err)
    
    ctx := context.Background()
    payload := &types.CreateVAPayload{
        // ... test payload
    }
    
    result, err := client.VirtualAccount().CreateVA(ctx, payload)
    
    // Assertions
    assert.NoError(t, err)
    assert.NotNil(t, result)
    // Add more specific assertions based on expected response
}
```

### Integration Testing

```go
func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    
    // Use real sandbox credentials
    client := setupRealClient(t)
    
    // Test full flow
    testVirtualAccountFlow(t, client)
    testMPMFlow(t, client)
    testQRFlow(t, client)
}
```

## 🚀 Production Deployment

### 1. Environment Setup

```bash
# Production environment variables
export ASPI_CLIENT_ID="your-production-client-id"
export ASPI_CLIENT_SECRET="your-production-client-secret"
export ASPI_PRIVATE_KEY_PATH="/secure/path/to/private_key.pem"
export ASPI_PUBLIC_KEY_PATH="/secure/path/to/public_key.pem"
export ASPI_ENVIRONMENT="production"
export ASPI_LOG_LEVEL="warn"
```

### 2. Security Considerations

- **Key Management**: Store RSA keys securely (e.g., AWS KMS, HashiCorp Vault)
- **Network Security**: Use HTTPS and validate certificates
- **Access Control**: Limit API access to authorized services only
- **Monitoring**: Implement comprehensive logging and monitoring

### 3. Performance Optimization

```go
// Production client configuration
client, err := snap.NewClient(snap.Config{
    BaseURL:         "https://api.aspi-indonesia.or.id",
    ClientID:        os.Getenv("ASPI_CLIENT_ID"),
    ClientSecret:    os.Getenv("ASPI_CLIENT_SECRET"),
    PrivateKeyPath:  os.Getenv("ASPI_PRIVATE_KEY_PATH"),
    PublicKeyPath:   os.Getenv("ASPI_PUBLIC_KEY_PATH"),
    Environment:     "production",
    ConnectTimeout:  5,  // Faster timeout for production
    RequestTimeout:  15, // Reasonable timeout
    MaxRetries:      2,  // Limited retries
    BackoffDuration: 1,
    LogLevel:        "warn", // Reduce log verbosity
})
```

### 4. Health Checks

```go
func healthCheck(client *snap.Client) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Simple health check - get B2B token
    _, err := client.Auth().GetB2BToken(ctx)
    return err
}
```

## 📞 Support and Resources

- **Documentation**: [ASPI API Documentation](https://aspi-indonesia.or.id/docs)
- **SDK Repository**: [GitHub Repository](https://github.com/mseptiaan/snap-aspi-go)
- **Issues**: Report issues on GitHub
- **Examples**: Check the `examples/` directory for more use cases

## 🔄 Migration from HTTP Server

If you're migrating from the HTTP server version:

1. **Remove HTTP dependencies**: No need for Gin, HTTP handlers, etc.
2. **Direct service calls**: Use SDK methods directly instead of HTTP endpoints
3. **Configuration changes**: Use SDK config instead of YAML config files
4. **Error handling**: Handle errors directly instead of HTTP status codes

### Before (HTTP Server)
```go
// HTTP request to create VA
resp, err := http.Post("/api/v1/virtual-account/create", payload)
```

### After (SDK)
```go
// Direct SDK call
result, err := client.VirtualAccount().CreateVA(ctx, payload)
```

This integration guide provides everything you need to successfully integrate the SNAP ASPI Go SDK into your application. Start with the basic examples and gradually implement more complex features as needed.