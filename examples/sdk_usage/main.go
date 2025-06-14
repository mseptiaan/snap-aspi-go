package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
	// Initialize the SNAP ASPI SDK client
	client, err := snap.NewClient(snap.Config{
		BaseURL:        "https://sandbox.aspi-indonesia.or.id",
		ClientID:       "your-client-id",
		ClientSecret:   "your-client-secret",
		PrivateKeyPath: "path/to/your/private_key.pem",
		PublicKeyPath:  "path/to/your/public_key.pem",
		Environment:    "sandbox", // or "production"
		LogLevel:       "info",
	})
	if err != nil {
		log.Fatalf("Failed to initialize SNAP client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Virtual Account Operations
	fmt.Println("=== Virtual Account Examples ===")
	
	// Create Virtual Account
	createVAPayload := &types.CreateVAPayload{
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

	result, err := client.VirtualAccount().CreateVA(ctx, createVAPayload)
	if err != nil {
		log.Printf("Failed to create VA: %v", err)
	} else {
		fmt.Printf("Create VA Result: %+v\n", result)
	}

	// Inquiry Virtual Account
	inquiryPayload := &types.InquiryVAPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      "2024-01-01",
	}

	result, err = client.VirtualAccount().InquiryVA(ctx, inquiryPayload)
	if err != nil {
		log.Printf("Failed to inquiry VA: %v", err)
	} else {
		fmt.Printf("Inquiry VA Result: %+v\n", result)
	}

	// Example 2: MPM Operations
	fmt.Println("\n=== MPM Examples ===")
	
	// MPM Transfer
	transferPayload := &types.MPMTransferPayload{
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

	result, err = client.MPM().Transfer(ctx, transferPayload)
	if err != nil {
		log.Printf("Failed to transfer MPM: %v", err)
	} else {
		fmt.Printf("MPM Transfer Result: %+v\n", result)
	}

	// MPM Balance Inquiry
	balancePayload := &types.MPMBalanceInquiryPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		AccountNo:        "1234567890",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err = client.MPM().BalanceInquiry(ctx, balancePayload)
	if err != nil {
		log.Printf("Failed to inquiry balance: %v", err)
	} else {
		fmt.Printf("Balance Inquiry Result: %+v\n", result)
	}

	// Example 3: QR Operations
	fmt.Println("\n=== QR Examples ===")
	
	// Generate QR
	qrPayload := &types.MPMGenerateQRPayload{
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
		ValidityPeriod:     "3600",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "POS",
		},
	}

	qrResult, err := client.MPM().GenerateQR(ctx, qrPayload)
	if err != nil {
		log.Printf("Failed to generate QR: %v", err)
	} else {
		fmt.Printf("Generate QR Result: %+v\n", qrResult)
	}

	// Example 4: Authentication
	fmt.Println("\n=== Authentication Examples ===")
	
	// Get B2B Token
	b2bToken, err := client.Auth().GetB2BToken(ctx)
	if err != nil {
		log.Printf("Failed to get B2B token: %v", err)
	} else {
		fmt.Printf("B2B Token: %s\n", b2bToken.AccessToken)
	}

	fmt.Println("\n=== SDK Usage Examples Complete ===")
}