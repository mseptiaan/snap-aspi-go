# Custom Endpoints Configuration Guide

This guide explains how to configure custom endpoints for different banks when using the SNAP ASPI Go SDK.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Default vs Custom Endpoints](#default-vs-custom-endpoints)
3. [Bank Presets](#bank-presets)
4. [Custom Configuration](#custom-configuration)
5. [Mixed Configuration](#mixed-configuration)
6. [Integration Examples](#integration-examples)
7. [Best Practices](#best-practices)

## 🔍 Overview

Different banks may have varying API endpoint structures for ASPI services. The SDK supports:

- **Default endpoints** - Standard ASPI endpoints
- **Bank presets** - Predefined configurations for major Indonesian banks
- **Custom endpoints** - Fully customizable endpoint configuration
- **Mixed configuration** - Combining bank presets with custom overrides

## 🔄 Default vs Custom Endpoints

### Default Endpoints (Standard ASPI)

```go
// Default configuration - uses standard ASPI endpoints
client, err := snap.NewClient(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "your-client-id",
    ClientSecret:   "your-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
    // No CustomEndpoints specified - uses defaults
})
```

**Default Virtual Account Endpoints:**
- Create VA: `/api/v1.0/transfer-va/create-va`
- Update VA: `/api/v1.0/transfer-va/update-va`
- Delete VA: `/api/v1.0/transfer-va/delete-va`
- Inquiry VA: `/api/v1.0/transfer-va/inquiry-va`
- Payment: `/api/v1.0/transfer-va/payment`

**Default MPM Endpoints:**
- Transfer: `/api/v1.0/transfer-kredit/mpm`
- Inquiry: `/api/v1.0/transfer-kredit/mpm/inquiry`
- Generate QR: `/api/v1.0/qr/qr-mpm-generate`

### Custom Endpoints

```go
// Custom configuration - override specific endpoints
customEndpoints := &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA:  "/api/v2.0/custom-bank/va/create",
        InquiryVA: "/api/v2.0/custom-bank/va/inquiry",
        // Other endpoints will use defaults
    },
    MPM: &snap.MPMEndpoints{
        Transfer: "/api/v2.0/custom-bank/mpm/transfer",
        // Other endpoints will use defaults
    },
}

client, err := snap.NewClient(snap.Config{
    BaseURL:         "https://custom-bank-api.example.com",
    ClientID:        "your-client-id",
    ClientSecret:    "your-client-secret",
    PrivateKeyPath:  "path/to/private_key.pem",
    CustomEndpoints: customEndpoints,
})
```

## 🏦 Bank Presets

The SDK includes predefined configurations for major Indonesian banks:

### Supported Banks

- **BCA** (Bank Central Asia)
- **BNI** (Bank Negara Indonesia)
- **BRI** (Bank Rakyat Indonesia)
- **MANDIRI** (Bank Mandiri)
- **CIMB** (CIMB Niaga)
- **PERMATA** (Bank Permata)

### Using Bank Presets

```go
// Method 1: Using NewClientForBank
client, err := snap.NewClientForBank(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bca-client-id",
    ClientSecret:   "bca-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
}, "BCA") // Bank code

// Method 2: Getting bank configuration manually
bankPresets := &snap.BankPresets{}
bcaEndpoints := bankPresets.GetBankConfig("BCA")

client, err := snap.NewClient(snap.Config{
    BaseURL:         "https://sandbox.aspi-indonesia.or.id",
    ClientID:        "bca-client-id",
    ClientSecret:    "bca-client-secret",
    PrivateKeyPath:  "path/to/private_key.pem",
    CustomEndpoints: bcaEndpoints,
})
```

### Bank-Specific Endpoints

#### BCA Endpoints
```
Virtual Account:
- Create: /api/v1.0/bca/transfer-va/create-va
- Inquiry: /api/v1.0/bca/transfer-va/inquiry-va
- Payment: /api/v1.0/bca/transfer-va/payment

MPM:
- Transfer: /api/v1.0/bca/transfer-kredit/mpm
- QR Generate: /api/v1.0/bca/qr/qr-mpm-generate
```

#### BNI Endpoints
```
Virtual Account:
- Create: /api/v1.0/bni/virtual-account/create
- Inquiry: /api/v1.0/bni/virtual-account/inquiry-va
- Payment: /api/v1.0/bni/virtual-account/payment

MPM:
- Transfer: /api/v1.0/bni/mpm/transfer
- QR Generate: /api/v1.0/bni/qr/generate
```

#### BRI Endpoints
```
Virtual Account:
- Create: /api/v1.0/bri/va/create
- Inquiry: /api/v1.0/bri/va/inquiry-va
- Payment: /api/v1.0/bri/va/payment

MPM:
- Transfer: /api/v1.0/bri/mpm-transfer/transfer
- QR Generate: /api/v1.0/bri/qr-mpm/generate
```

## 🔧 Custom Configuration

### Complete Custom Configuration

```go
customEndpoints := &snap.CustomEndpoints{
    // Authentication endpoints
    B2BToken:   "/api/v2.0/auth/b2b-token",
    B2B2CToken: "/api/v2.0/auth/b2b2c-token",
    
    // Virtual Account endpoints
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA:     "/api/v2.0/custom/va/create",
        UpdateVA:     "/api/v2.0/custom/va/update",
        DeleteVA:     "/api/v2.0/custom/va/delete",
        InquiryVA:    "/api/v2.0/custom/va/inquiry-va",
        Inquiry:      "/api/v2.0/custom/va/inquiry",
        Payment:      "/api/v2.0/custom/va/payment",
        Status:       "/api/v2.0/custom/va/status",
        Report:       "/api/v2.0/custom/va/report",
        UpdateStatus: "/api/v2.0/custom/va/update-status",
    },
    
    // MPM endpoints
    MPM: &snap.MPMEndpoints{
        Transfer:       "/api/v2.0/custom/mpm/transfer",
        Inquiry:        "/api/v2.0/custom/mpm/inquiry",
        Status:         "/api/v2.0/custom/mpm/status",
        Refund:         "/api/v2.0/custom/mpm/refund",
        BalanceInquiry: "/api/v2.0/custom/mpm/balance",
        AccountInquiry: "/api/v2.0/custom/mpm/account",
        History:        "/api/v2.0/custom/mpm/history",
        GenerateQR:     "/api/v2.0/custom/qr/generate",
        NotifyQR:       "/api/v2.0/custom/qr/notify",
    },
}
```

### Partial Custom Configuration

```go
// Override only specific endpoints
customEndpoints := &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA: "/api/v1.0/special/va/create",
        Payment:  "/api/v1.0/special/va/payment",
        // Other VA endpoints will use defaults
    },
    // MPM endpoints will use defaults
}
```

## 🔀 Mixed Configuration

Combine bank presets with custom overrides:

```go
// Start with BNI preset, but override specific endpoints
customEndpoints := &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA: "/api/v1.0/bni/custom-va/create", // Override
        // Other endpoints will use BNI defaults
    },
    MPM: &snap.MPMEndpoints{
        Transfer: "/api/v1.0/bni/custom-mpm/transfer", // Override
        // Other endpoints will use BNI defaults
    },
}

client, err := snap.NewClientForBank(snap.Config{
    BaseURL:         "https://sandbox.aspi-indonesia.or.id",
    ClientID:        "bni-client-id",
    ClientSecret:    "bni-client-secret",
    PrivateKeyPath:  "path/to/private_key.pem",
    CustomEndpoints: customEndpoints, // Merged with BNI presets
}, "BNI")
```

## 💡 Integration Examples

### Example 1: Multi-Bank Application

```go
func createMultiBankClients() map[string]*snap.Client {
    banks := []string{"BCA", "BNI", "BRI"}
    clients := make(map[string]*snap.Client)
    
    for _, bankCode := range banks {
        client, err := snap.NewClientForBank(snap.Config{
            BaseURL:        "https://sandbox.aspi-indonesia.or.id",
            ClientID:       fmt.Sprintf("%s_CLIENT_ID", bankCode),
            ClientSecret:   fmt.Sprintf("%s_CLIENT_SECRET", bankCode),
            PrivateKeyPath: "path/to/private_key.pem",
        }, bankCode)
        
        if err != nil {
            log.Printf("Failed to create %s client: %v", bankCode, err)
            continue
        }
        
        clients[bankCode] = client
    }
    
    return clients
}
```

### Example 2: Dynamic Bank Selection

```go
func routeToBank(accountNumber string) string {
    switch {
    case strings.HasPrefix(accountNumber, "8808"):
        return "BCA"
    case strings.HasPrefix(accountNumber, "8809"):
        return "BNI"
    case strings.HasPrefix(accountNumber, "8810"):
        return "BRI"
    default:
        return "BCA" // Default
    }
}

func processTransaction(accountNumber string, clients map[string]*snap.Client) {
    bankCode := routeToBank(accountNumber)
    client := clients[bankCode]
    
    // Process with bank-specific client
    result, err := client.VirtualAccount().CreateVA(ctx, payload)
    // ...
}
```

### Example 3: Environment-Based Configuration

```go
func createClientForEnvironment(env string, bankCode string) (*snap.Client, error) {
    var baseURL string
    var customEndpoints *snap.CustomEndpoints
    
    switch env {
    case "development":
        baseURL = "https://dev.aspi-indonesia.or.id"
        customEndpoints = &snap.CustomEndpoints{
            VirtualAccount: &snap.VirtualAccountEndpoints{
                CreateVA: "/api/dev/va/create",
            },
        }
    case "staging":
        baseURL = "https://staging.aspi-indonesia.or.id"
    case "production":
        baseURL = "https://api.aspi-indonesia.or.id"
    default:
        baseURL = "https://sandbox.aspi-indonesia.or.id"
    }
    
    return snap.NewClientForBank(snap.Config{
        BaseURL:         baseURL,
        ClientID:        os.Getenv("ASPI_CLIENT_ID"),
        ClientSecret:    os.Getenv("ASPI_CLIENT_SECRET"),
        PrivateKeyPath:  os.Getenv("ASPI_PRIVATE_KEY_PATH"),
        Environment:     env,
        CustomEndpoints: customEndpoints,
    }, bankCode)
}
```

## 🎯 Best Practices

### 1. Configuration Management

```go
// Use environment variables for sensitive data
type BankConfig struct {
    Code         string `env:"BANK_CODE"`
    ClientID     string `env:"BANK_CLIENT_ID"`
    ClientSecret string `env:"BANK_CLIENT_SECRET"`
    BaseURL      string `env:"BANK_BASE_URL"`
}

// Load from configuration file
type EndpointsConfig struct {
    Banks map[string]snap.CustomEndpoints `yaml:"banks"`
}
```

### 2. Validation

```go
func validateEndpoints(endpoints *snap.CustomEndpoints) error {
    if endpoints.VirtualAccount != nil {
        if endpoints.VirtualAccount.CreateVA == "" {
            return fmt.Errorf("CreateVA endpoint is required")
        }
    }
    return nil
}
```

### 3. Fallback Strategy

```go
func createClientWithFallback(bankCode string) (*snap.Client, error) {
    // Try bank-specific configuration first
    client, err := snap.NewClientForBank(config, bankCode)
    if err != nil {
        // Fallback to default configuration
        return snap.NewClient(defaultConfig)
    }
    return client, nil
}
```

### 4. Testing Different Banks

```go
func TestBankEndpoints(t *testing.T) {
    banks := []string{"BCA", "BNI", "BRI"}
    
    for _, bankCode := range banks {
        t.Run(bankCode, func(t *testing.T) {
            client, err := snap.NewClientForBank(testConfig, bankCode)
            require.NoError(t, err)
            
            // Test bank-specific functionality
            testVirtualAccountOperations(t, client)
            testMPMOperations(t, client)
        })
    }
}
```

### 5. Monitoring and Logging

```go
func logEndpointUsage(bankCode string, endpoint string, success bool) {
    log.Printf("Bank: %s, Endpoint: %s, Success: %v", bankCode, endpoint, success)
}
```

## 🔍 Troubleshooting

### Common Issues

1. **Endpoint Not Found (404)**
   - Verify the custom endpoint URL is correct
   - Check if the bank supports the specific endpoint
   - Ensure the base URL is correct for the bank

2. **Authentication Failures**
   - Verify client credentials are correct for the specific bank
   - Check if the bank uses different authentication endpoints

3. **Different Response Formats**
   - Some banks may return slightly different response structures
   - Implement bank-specific response handling if needed

### Debug Configuration

```go
// Enable debug logging to see actual endpoints being used
client, err := snap.NewClient(snap.Config{
    // ... other config
    LogLevel: "debug",
})

// Check current configuration
config := client.GetConfig()
fmt.Printf("Current endpoints: %+v\n", config.CustomEndpoints)
```

This custom endpoints feature provides maximum flexibility for integrating with different banks while maintaining a consistent SDK interface.