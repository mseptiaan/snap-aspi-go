# SNAP ASPI Go SDK

A comprehensive, production-ready Go SDK for integrating with Indonesia's ASPI (Payment System Administration) services. This SDK provides a clean, type-safe interface for all ASPI endpoints including Virtual Account, MPM, Registration, Balance Inquiry, Transaction History, Transfer Credit, and Transfer Debit operations.

## 🚀 Quick Start

### Installation

```bash
go get github.com/mseptiaan/snap-aspi-go
```

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/mseptiaan/snap-aspi-go/pkg/snap"
    "github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
    // Initialize the SDK
    client, err := snap.NewClient(snap.Config{
        BaseURL:        "https://sandbox.aspi-indonesia.or.id",
        ClientID:       "your-client-id",
        ClientSecret:   "your-client-secret",
        PrivateKeyPath: "path/to/private_key.pem",
        PublicKeyPath:  "path/to/public_key.pem",
        Environment:    "sandbox",
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Create a Virtual Account
    result, err := client.VirtualAccount().CreateVA(ctx, &types.CreateVAPayload{
        PartnerServiceId:   "12345",
        CustomerNo:         "CUST001",
        VirtualAccountNo:   "8808001234567890",
        VirtualAccountName: "John Doe",
        TotalAmount:        types.NewAmount(100000, "IDR"),
        ExpiredDate:        "2024-12-31T23:59:59+07:00",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "DEVICE001",
            Channel:  "WEB",
        },
    })
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    
    log.Printf("VA Created: %+v", result)
}
```

### BCA-Specific Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/mseptiaan/snap-aspi-go/pkg/snap"
    "github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
    // Initialize the BCA client
    client, err := snap.NewBCAClient(snap.Config{
        BaseURL:        "https://sandbox.aspi-indonesia.or.id",
        ClientID:       "your-bca-client-id",
        ClientSecret:   "your-bca-client-secret",
        PrivateKeyPath: "path/to/private_key.pem",
        PublicKeyPath:  "path/to/public_key.pem",
        Environment:    "sandbox",
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Create a BCA Virtual Account
    result, err := client.CreateBCAVirtualAccount(ctx, &types.BCACreateVAPayload{
        PartnerServiceId:    "BCA12345",
        CustomerNo:          "BCA_CUSTOMER_001",
        VirtualAccountNo:    "8808001234567890",
        VirtualAccountName:  "BCA Customer",
        TotalAmount:         types.NewAmount(150000, "IDR"),
        ExpiredDate:         "2024-12-31T23:59:59+07:00",
        AdditionalInfo: &types.AdditionalInfo{
            DeviceId: "BCA_DEVICE_001",
            Channel:  "BCA_WEB",
        },
        // BCA-specific fields
        BcaCustomerNo:      "BCA123456789",
        BcaSubCompany:      "BCA Finance",
        BcaTransactionType: "1",
    })
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    
    log.Printf("BCA VA Created: %+v", result)
}
```

## 📚 Features

- **🔐 Authentication**: Automatic B2B and B2B2C token management
- **💳 Virtual Accounts**: Complete VA lifecycle management
- **🏪 MPM Operations**: Merchant payment management and transfers
- **📱 QR Codes**: Dynamic QR code generation and notifications
- **👤 Registration**: User registration, card registration, and account binding
- **💰 Balance Inquiry**: Account balance information
- **📊 Transaction History**: Transaction history and bank statements
- **💸 Transfer Credit**: Account inquiry, transfers, and top-ups
- **💳 Transfer Debit**: Direct debit, CPM, auth payment, and BI-FAST
- **🔄 Auto-retry**: Built-in retry logic with exponential backoff
- **🚀 Performance**: Connection pooling and request optimization
- **🛡️ Security**: RSA signature validation and secure communication
- **📝 Type Safety**: Full Go type definitions for all API operations
- **🏦 Multi-Bank**: Support for different bank endpoints including BCA-specific integration

## 🏗️ Architecture

```
┌─────────────────┐
│   Your App      │
└─────────┬───────┘
          │
┌─────────▼───────┐
│   SNAP SDK      │
├─────────────────┤
│ • VirtualAccount│
│ • MPM           │
│ • Registration  │
│ • BalanceInquiry│
│ • TxHistory     │
│ • TransferCredit│
│ • TransferDebit │
│ • Auth          │
│ • BCA           │ ← New BCA-specific integration
└─────────┬───────┘
          │
┌─────────▼───────┐
│   ASPI API      │
└─────────────────┘
```

## 🔧 Configuration

### Environment Variables

```bash
export ASPI_CLIENT_ID="your-client-id"
export ASPI_CLIENT_SECRET="your-client-secret"
export ASPI_PRIVATE_KEY_PATH="/path/to/private_key.pem"
export ASPI_PUBLIC_KEY_PATH="/path/to/public_key.pem"
export ASPI_ENVIRONMENT="sandbox"  # or "production"

# BCA-specific environment variables
export BCA_CLIENT_ID="your-bca-client-id"
export BCA_CLIENT_SECRET="your-bca-client-secret"
```

## 🏦 Multi-Bank Support

The SDK supports different bank endpoints through bank presets or custom configurations:

### Using Bank Presets

```go
// BCA Bank
client, err := snap.NewClientForBank(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bca-client-id",
    ClientSecret:   "bca-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
}, "BCA")

// Or use the specialized BCA client
bcaClient, err := snap.NewBCAClient(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bca-client-id",
    ClientSecret:   "bca-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
})
```

### Supported Banks

- **BCA** (Bank Central Asia) - Enhanced with BCA-specific endpoints and types
- **BNI** (Bank Negara Indonesia)
- **BRI** (Bank Rakyat Indonesia)
- **MANDIRI** (Bank Mandiri)
- **CIMB** (CIMB Niaga)
- **PERMATA** (Bank Permata)

## 📝 Examples

Check out the `examples` directory for complete examples of using the SDK:

- `examples/basic/main.go`: Basic usage examples
- `examples/multi_bank/main.go`: Multi-bank integration examples
- `examples/custom_endpoints/main.go`: Custom endpoint configuration examples
- `examples/complete/main.go`: Comprehensive examples of all features
- `examples/bca/main.go`: BCA-specific integration examples

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.