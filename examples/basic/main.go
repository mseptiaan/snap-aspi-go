package main

import (
	"context"
	"log"
	"os"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
	// Initialize the SDK
	client, err := snap.NewClient(snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("ASPI_CLIENT_ID", "your-client-id"),
		ClientSecret:   getEnvOrDefault("ASPI_CLIENT_SECRET", "your-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:    "sandbox",
		LogLevel:       "info",
	})
	if err != nil {
		log.Fatalf("Failed to initialize SNAP client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Create Virtual Account
	log.Println("=== Creating Virtual Account ===")
	createVirtualAccount(ctx, client)

	// Example 2: MPM Transfer
	log.Println("\n=== MPM Transfer ===")
	mpmTransfer(ctx, client)

	// Example 3: Generate QR Code
	log.Println("\n=== Generate QR Code ===")
	generateQRCode(ctx, client)

	// Example 4: Authentication
	log.Println("\n=== Authentication ===")
	testAuthentication(ctx, client)
}

func createVirtualAccount(ctx context.Context, client *snap.Client) {
	result, err := client.VirtualAccount().CreateVA(ctx, &types.CreateVAPayload{
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
	})
	if err != nil {
		log.Printf("Create VA Error: %v", err)
		return
	}
	log.Printf("VA Created: %+v", result)
}

func mpmTransfer(ctx context.Context, client *snap.Client) {
	result, err := client.MPM().Transfer(ctx, &types.MPMTransferPayload{
		PartnerServiceId:      "12345",
		CustomerNo:            "CUST001",
		PartnerReferenceNo:    "REF001",
		MerchantId:            "MERCHANT001",
		Amount:                types.NewAmount(50000, "IDR"),
		BeneficiaryAccountNo:  "1234567890",
		BeneficiaryName:       "Jane Doe",
		Currency:              "IDR",
		TransactionDate:       "2024-01-15T10:30:00+07:00",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	})
	if err != nil {
		log.Printf("MPM Transfer Error: %v", err)
		return
	}
	log.Printf("Transfer Result: %+v", result)
}

func generateQRCode(ctx context.Context, client *snap.Client) {
	result, err := client.MPM().GenerateQR(ctx, &types.MPMGenerateQRPayload{
		PartnerServiceId:   "12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "QR001",
		MerchantId:         "MERCHANT001",
		Amount:             types.NewAmount(25000, "IDR"),
		MerchantName:       "Test Merchant",
		MerchantCity:       "Jakarta",
		QRType:             "DYNAMIC",
		ValidityPeriod:     "3600",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "POS",
		},
	})
	if err != nil {
		log.Printf("Generate QR Error: %v", err)
		return
	}
	log.Printf("QR Generated: %s", result.QRContent)
}

func testAuthentication(ctx context.Context, client *snap.Client) {
	token, err := client.Auth().GetB2BToken(ctx)
	if err != nil {
		log.Printf("B2B Token Error: %v", err)
		return
	}
	log.Printf("B2B Token: %s", token.AccessToken)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}