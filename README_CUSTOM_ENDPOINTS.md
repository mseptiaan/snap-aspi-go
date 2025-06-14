# Custom Endpoints Support

The SNAP ASPI Go SDK now supports custom endpoints for different banks, allowing you to integrate with various Indonesian banks that may have different API endpoint structures.

## 🚀 Quick Start

### Using Bank Presets

```go
// BCA Bank
client, err := snap.NewClientForBank(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bca-client-id",
    ClientSecret:   "bca-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
}, "BCA")

// BNI Bank
client, err := snap.NewClientForBank(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bni-client-id",
    ClientSecret:   "bni-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
}, "BNI")
```

### Using Custom Endpoints

```go
customEndpoints := &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA:  "/api/v2.0/custom-bank/va/create",
        InquiryVA: "/api/v2.0/custom-bank/va/inquiry",
        Payment:   "/api/v2.0/custom-bank/va/payment",
    },
    MPM: &snap.MPMEndpoints{
        Transfer:   "/api/v2.0/custom-bank/mpm/transfer",
        GenerateQR: "/api/v2.0/custom-bank/qr/generate",
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

## 🏦 Supported Banks

The SDK includes predefined configurations for:

- **BCA** (Bank Central Asia)
- **BNI** (Bank Negara Indonesia)  
- **BRI** (Bank Rakyat Indonesia)
- **MANDIRI** (Bank Mandiri)
- **CIMB** (CIMB Niaga)
- **PERMATA** (Bank Permata)

## 📊 Endpoint Comparison

| Bank | VA Create Endpoint | MPM Transfer Endpoint |
|------|-------------------|----------------------|
| **Default** | `/api/v1.0/transfer-va/create-va` | `/api/v1.0/transfer-kredit/mpm` |
| **BCA** | `/api/v1.0/bca/transfer-va/create-va` | `/api/v1.0/bca/transfer-kredit/mpm` |
| **BNI** | `/api/v1.0/bni/virtual-account/create` | `/api/v1.0/bni/mpm/transfer` |
| **BRI** | `/api/v1.0/bri/va/create` | `/api/v1.0/bri/mpm-transfer/transfer` |

## 🔧 Configuration Options

### 1. Bank Presets (Recommended)

```go
client, err := snap.NewClientForBank(config, "BCA")
```

### 2. Custom Endpoints

```go
config.CustomEndpoints = &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA: "/custom/endpoint",
    },
}
```

### 3. Mixed Configuration

```go
// Start with BNI preset, override specific endpoints
config.CustomEndpoints = &snap.CustomEndpoints{
    VirtualAccount: &snap.VirtualAccountEndpoints{
        CreateVA: "/api/v1.0/bni/custom-va/create",
    },
}
client, err := snap.NewClientForBank(config, "BNI")
```

## 💼 Multi-Bank Support

```go
// Create clients for multiple banks
banks := []string{"BCA", "BNI", "BRI"}
clients := make(map[string]*snap.Client)

for _, bankCode := range banks {
    client, err := snap.NewClientForBank(snap.Config{
        BaseURL:        "https://sandbox.aspi-indonesia.or.id",
        ClientID:       fmt.Sprintf("%s_CLIENT_ID", bankCode),
        ClientSecret:   fmt.Sprintf("%s_CLIENT_SECRET", bankCode),
        PrivateKeyPath: "path/to/private_key.pem",
    }, bankCode)
    
    if err == nil {
        clients[bankCode] = client
    }
}

// Route transactions to appropriate bank
func processTransaction(accountNumber string) {
    bankCode := routeToBank(accountNumber)
    client := clients[bankCode]
    
    result, err := client.VirtualAccount().CreateVA(ctx, payload)
    // ...
}
```

## 🎯 Use Cases

### 1. Bank-Specific Integration
When integrating with a specific bank that has custom endpoint structures.

### 2. Multi-Bank Applications
Applications that need to support multiple banks with different endpoint configurations.

### 3. Testing and Development
Using different endpoints for development, staging, and production environments.

### 4. Legacy System Integration
Integrating with existing bank systems that may have non-standard endpoint structures.

## 📚 Examples

Check out the examples directory:

- **[Custom Endpoints Example](examples/custom_endpoints/)** - Complete custom endpoint configuration
- **[Bank Specific Example](examples/bank_specific/)** - Bank-specific integration patterns
- **[Multi-Bank Example](examples/multi_bank/)** - Supporting multiple banks

## 🔍 Documentation

- **[Custom Endpoints Guide](CUSTOM_ENDPOINTS_GUIDE.md)** - Comprehensive configuration guide
- **[Integration Guide](INTEGRATION_GUIDE.md)** - General integration instructions
- **[API Reference](docs/API.md)** - Complete API documentation

## ⚡ Benefits

- **✅ Flexibility** - Support any bank's endpoint structure
- **✅ Backward Compatibility** - Default endpoints still work
- **✅ Easy Migration** - Simple bank preset selection
- **✅ Multi-Bank Support** - Handle multiple banks in one application
- **✅ Production Ready** - Tested with major Indonesian banks

## 🚀 Migration

### From Default Endpoints

```go
// Before
client, err := snap.NewClient(config)

// After (using bank preset)
client, err := snap.NewClientForBank(config, "BCA")
```

### From HTTP Server Version

The custom endpoints feature maintains full compatibility with the existing SDK interface while adding bank-specific endpoint support.

---

**Ready to integrate with specific banks?** Check out the [Custom Endpoints Guide](CUSTOM_ENDPOINTS_GUIDE.md) for detailed configuration instructions.