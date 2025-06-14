# SNAP ASPI Go SDK Documentation

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Services](#services)
   - [Authentication](#authentication)
   - [Virtual Account](#virtual-account)
   - [MPM (Merchant Payment Management)](#mpm-merchant-payment-management)
   - [Registration](#registration)
   - [Balance Inquiry](#balance-inquiry)
   - [Transaction History](#transaction-history)
   - [Transfer Credit](#transfer-credit)
   - [Transfer Debit](#transfer-debit)
5. [Multi-Bank Support](#multi-bank-support)
6. [Error Handling](#error-handling)
7. [Best Practices](#best-practices)
8. [Examples](#examples)
9. [API Reference](#api-reference)

## Introduction

The SNAP ASPI Go SDK provides a comprehensive, production-ready interface for integrating with Indonesia's ASPI (Payment System Administration) services. This SDK supports all ASPI endpoints including Virtual Account, MPM, Registration, Balance Inquiry, Transaction History, Transfer Credit, and Transfer Debit operations.

## Installation

```bash
go get github.com/mseptiaan/snap-aspi-go
```

## Configuration

### Basic Configuration

```go
client, err := snap.NewClient(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "your-client-id",
    ClientSecret:   "your-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
    PublicKeyPath:  "path/to/public_key.pem",
    Environment:    "sandbox", // or "production"
    LogLevel:       "info",    // debug, info, warn, error
})
```

### Advanced Configuration

```go
client, err := snap.NewClient(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "your-client-id",
    ClientSecret:   "your-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
    PublicKeyPath:  "path/to/public_key.pem",
    Environment:    "sandbox",
    
    // Timeout configurations
    ConnectTimeout: 10, // seconds
    RequestTimeout: 30, // seconds
    
    // Retry configurations
    MaxRetries:      3,
    BackoffDuration: 1, // seconds
    
    // Logging
    LogLevel: "info", // debug, info, warn, error
    
    // Custom endpoints
    CustomEndpoints: &snap.CustomEndpoints{
        VirtualAccount: &snap.VirtualAccountEndpoints{
            CreateVA: "/api/v2.0/custom/va/create",
        },
        MPM: &snap.MPMEndpoints{
            GenerateQR: "/api/v2.0/custom/qr/generate",
        },
    },
})
```

## Services

### Authentication

The Authentication service handles token management for API access.

```go
// Get B2B token
token, err := client.Auth().GetB2BToken(ctx)
if err != nil {
    log.Printf("Error: %v", err)
    return
}
fmt.Printf("Access Token: %s\n", token.AccessToken)

// Get B2B2C token
customerToken, err := client.Auth().GetB2B2CToken(ctx, auth.CustomerTokenRequest{
    AuthCode:     "auth-code",
    RefreshToken: "refresh-token",
    AdditionalInfo: map[string]interface{}{
        "deviceId": "DEVICE001",
    },
})
```

### Virtual Account

The Virtual Account service provides methods for managing virtual accounts.

```go
// Create Virtual Account
result, err := client.VirtualAccount().CreateVA(ctx, &types.CreateVAPayload{
    PartnerServiceId:   "12345",
    CustomerNo:         "CUST001",
    VirtualAccountNo:   "8808001234567890",
    VirtualAccountName: "John Doe",
    TrxId:              "TRX-001",
    TotalAmount:        types.NewAmount(100000, "IDR"),
    ExpiredDate:        "2024-12-31T23:59:59+07:00",
    AdditionalInfo: &types.AdditionalInfo{
        DeviceId: "DEVICE001",
        Channel:  "WEB",
    },
})

// Inquiry Virtual Account
result, err := client.VirtualAccount().InquiryVA(ctx, &types.InquiryVAPayload{
    PartnerServiceId: "12345",
    CustomerNo:       "CUST001",
    VirtualAccountNo: "8808001234567890",
    TrxDateInit:      "2024-01-01",
})

// Process Payment
result, err := client.VirtualAccount().Payment(ctx, &types.PaymentPayload{
    PartnerServiceId:   "12345",
    CustomerNo:         "CUST001",
    VirtualAccountNo:   "8808001234567890",
    VirtualAccountName: "John Doe",
    TrxId:              "TRX-001",
    PaymentRequestId:   "PAY-001",
    ChannelCode:        "6014",
    SourceAccountNo:    "1234567890",
    PaidAmount:         types.NewAmount(100000, "IDR"),
    TotalAmount:        types.NewAmount(100000, "IDR"),
    TrxDateTime:        "2024-01-15T10:30:00+07:00",
    ReferenceNo:        "REF001",
    PaymentType:        "1",
    FlagAdvise:         "N",
})
```

### MPM (Merchant Payment Management)

The MPM service provides methods for merchant payment management.

```go
// Generate QR Code
qrResult, err := client.MPM().GenerateQR(ctx, &types.MPMGenerateQRPayload{
    PartnerServiceId:   "12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "QR001",
    MerchantId:         "MERCHANT001",
    Amount:             types.NewAmount(25000, "IDR"),
    MerchantName:       "Test Merchant",
    MerchantCity:       "Jakarta",
    QRType:             "DYNAMIC",
    ValidityPeriod:     "3600",
})
if err != nil {
    log.Printf("Error: %v", err)
    return
}
fmt.Printf("QR Content: %s\n", qrResult.QRContent)
fmt.Printf("Valid Until: %s\n", qrResult.ValidUntil)

// Decode QR Code
result, err := client.MPM().DecodeQR(ctx, &types.MPMDecodeQRPayload{
    PartnerServiceId:   "12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "QR002",
    QRContent:          "00020101...",
    Amount:             types.NewAmount(25000, "IDR"),
    ScanTime:           "2024-01-15T10:30:00+07:00",
})

// Process Payment
result, err := client.MPM().Payment(ctx, &types.MPMPaymentPayload{
    PartnerServiceId:   "12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "PAY001",
    MerchantId:         "MERCHANT001",
    Amount:             types.NewAmount(25000, "IDR"),
})
```

### Registration

The Registration service provides methods for user registration, card registration, and account binding.

```go
// User Registration
result, err := client.Registration().Register(ctx, &types.RegistrationPayload{
    PartnerServiceId: "REG12345",
    CustomerNo:       "CUST001",
    PhoneNo:          "081234567890",
    Email:            "user@example.com",
    Name:             "John Doe",
    IdNumber:         "1234567890123456",
    IdType:           "KTP",
    Address:          "Jl. Example No. 123, Jakarta",
    DateOfBirth:      "1990-01-01",
    PlaceOfBirth:     "Jakarta",
    Nationality:      "ID",
    AdditionalInfo: &types.AdditionalInfo{
        DeviceId: "DEVICE001",
        Channel:  "MOBILE",
    },
})

// Card Registration
result, err := client.Registration().RegisterCard(ctx, &types.CardRegistrationPayload{
    PartnerServiceId: "REG12345",
    CustomerNo:       "CUST001",
    CardNumber:       "4111111111111111",
    CardType:         "CREDIT",
    ExpiryDate:       "12/25",
    CVV:              "123",
    CardHolderName:   "John Doe",
})

// Account Binding
result, err := client.Registration().BindAccount(ctx, &types.AccountBindingPayload{
    PartnerServiceId: "REG12345",
    CustomerNo:       "CUST001",
    AccountNo:        "1234567890",
    BankCode:         "014",
    AccountName:      "John Doe",
})
```

### Balance Inquiry

The Balance Inquiry service provides methods for checking account balances.

```go
// Balance Inquiry
result, err := client.BalanceInquiry().BalanceInquiry(ctx, &types.BalanceInquiryPayload{
    PartnerServiceId: "BAL12345",
    CustomerNo:       "CUST001",
    AccountNo:        "1234567890",
    BalanceTypes:     []string{"Cash", "Coins"},
})
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Display account balances
fmt.Printf("Account Number: %s\n", result.AccountNo)
fmt.Printf("Account Name: %s\n", result.Name)
for i, info := range result.AccountInfos {
    fmt.Printf("Balance Type %d: %s\n", i+1, info.BalanceType)
    fmt.Printf("Available Balance: %s\n", info.AvailableBalance.String())
    fmt.Printf("Ledger Balance: %s\n", info.LedgerBalance.String())
}
```

### Transaction History

The Transaction History service provides methods for retrieving transaction history and bank statements.

```go
// Transaction History List
historyResult, err := client.TransactionHistory().TransactionHistoryList(ctx, &types.TransactionHistoryListPayload{
    PartnerServiceId: "HIST12345",
    CustomerNo:       "CUST001",
    FromDateTime:     "2024-01-01T00:00:00+07:00",
    ToDateTime:       "2024-01-31T23:59:59+07:00",
    PageSize:         10,
    PageNumber:       0,
})
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Display transactions
fmt.Printf("Found %d transactions\n", len(historyResult.DetailData))
for i, tx := range historyResult.DetailData {
    fmt.Printf("Transaction %d: %s - %s - %s\n", 
        i+1, tx.DateTime, tx.Amount.String(), tx.Type)
}

// Bank Statement
statementResult, err := client.TransactionHistory().BankStatement(ctx, &types.BankStatementPayload{
    PartnerServiceId: "STMT12345",
    CustomerNo:       "CUST001",
    AccountNo:        "1234567890",
    FromDateTime:     "2024-01-01T00:00:00+07:00",
    ToDateTime:       "2024-01-31T23:59:59+07:00",
})
```

### Transfer Credit

The Transfer Credit service provides methods for account inquiry, transfers, and top-ups.

```go
// Account Inquiry
inquiryResult, err := client.TransferCredit().AccountInquiry(ctx, &types.AccountInquiryPayload{
    PartnerServiceId:     "TRF12345",
    CustomerNo:           "CUST001",
    PartnerReferenceNo:   "REF001",
    BeneficiaryAccountNo: "1234567890",
    BeneficiaryBankCode:  "014",
})

// Trigger Transfer
transferResult, err := client.TransferCredit().TriggerTransfer(ctx, &types.TriggerTransferPayload{
    PartnerServiceId:       "TRF12345",
    CustomerNo:             "CUST001",
    PartnerReferenceNo:     "REF002",
    Amount:                 types.NewAmount(50000, "IDR"),
    BeneficiaryAccountNo:   "1234567890",
    BeneficiaryAccountName: "Jane Doe",
    BeneficiaryBankCode:    "014",
    BeneficiaryBankName:    "BCA",
    SourceAccountNo:        "0987654321",
    TransactionDate:        "2024-01-15T10:30:00+07:00",
    FeeType:                "OUR",
    Remark:                 "Transfer for payment",
    Currency:               "IDR",
})

// Customer Top Up
topupResult, err := client.TransferCredit().CustomerTopUp(ctx, &types.CustomerTopUpPayload{
    PartnerServiceId: "TOP12345",
    CustomerNo:       "CUST001",
    CustomerName:     "John Doe",
    Amount:           types.NewAmount(100000, "IDR"),
    TransactionDate:  "2024-01-15T10:30:00+07:00",
})
```

### Transfer Debit

The Transfer Debit service provides methods for direct debit, CPM, auth payment, and BI-FAST operations.

```go
// Direct Debit Payment
debitResult, err := client.TransferDebit().DirectDebitPayment(ctx, &types.DirectDebitPaymentPayload{
    PartnerServiceId:   "DEB12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "REF003",
    BankCardToken:      "TOKEN001",
    MerchantId:         "MERCH001",
    Amount:             types.NewAmount(75000, "IDR"),
})

// CPM Generate QR
qrResult, err := client.TransferDebit().CPMGenerateQR(ctx, &types.CPMGenerateQRPayload{
    PartnerServiceId:   "QR12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "REF004",
    UserAccessToken:    "TOKEN002",
    MerchantId:         "MERCH001",
    PartnerTrxDate:     "2024-01-15T10:30:00+07:00",
})

// Auth Payment
authResult, err := client.TransferDebit().AuthPayment(ctx, &types.AuthPaymentPayload{
    PartnerServiceId:   "AUTH12345",
    CustomerNo:         "CUST001",
    PartnerReferenceNo: "REF005",
    BankCardToken:      "TOKEN003",
    MerchantId:         "MERCH001",
    Amount:             types.NewAmount(120000, "IDR"),
})
```

## Multi-Bank Support

The SDK provides built-in support for different bank endpoints through bank presets.

```go
// BCA Bank
client, err := snap.NewClientForBank(snap.Config{
    BaseURL:        "https://sandbox.aspi-indonesia.or.id",
    ClientID:       "bca-client-id",
    ClientSecret:   "bca-client-secret",
    PrivateKeyPath: "path/to/private_key.pem",
}, "BCA")

// BNI Bank
client, err := snap.NewClientForBank(config, "BNI")

// BRI Bank
client, err := snap.NewClientForBank(config, "BRI")

// Mandiri Bank
client, err := snap.NewClientForBank(config, "MANDIRI")

// CIMB Bank
client, err := snap.NewClientForBank(config, "CIMB")

// Permata Bank
client, err := snap.NewClientForBank(config, "PERMATA")
```

## Error Handling

The SDK provides comprehensive error handling with specific error types.

```go
result, err := client.VirtualAccount().CreateVA(ctx, payload)
if err != nil {
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
```

## Best Practices

### Connection Management

```go
// Create a singleton client for your application
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

### Logging and Monitoring

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

## Examples

Check out the `examples` directory for complete examples of using the SDK:

- `examples/basic/main.go`: Basic usage examples
- `examples/multi_bank/main.go`: Multi-bank integration examples
- `examples/custom_endpoints/main.go`: Custom endpoint configuration examples
- `examples/complete/main.go`: Comprehensive examples of all features

## API Reference

For detailed API reference, please refer to the [ASPI API Documentation](https://aspi-indonesia.or.id/documentation).