package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
	fmt.Println("=== SNAP ASPI SDK Custom Endpoints Examples ===")

	// Example 1: Using predefined bank configuration
	fmt.Println("\n1. Using Predefined Bank Configuration (BCA)")
	demonstrateBankPreset()

	// Example 2: Using custom endpoints
	fmt.Println("\n2. Using Custom Endpoints")
	demonstrateCustomEndpoints()

	// Example 3: Mixing bank preset with custom overrides
	fmt.Println("\n3. Mixing Bank Preset with Custom Overrides")
	demonstrateMixedConfiguration()

	fmt.Println("\n=== Custom Endpoints Examples Complete ===")
}

// demonstrateBankPreset shows how to use predefined bank configurations
func demonstrateBankPreset() {
	// Using BCA predefined configuration
	client, err := snap.NewClientForBank(snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("BCA_CLIENT_ID", "bca-client-id"),
		ClientSecret:   getEnvOrDefault("BCA_CLIENT_SECRET", "bca-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:    "sandbox",
		LogLevel:       "info",
	}, "BCA") // Bank code for BCA

	if err != nil {
		log.Printf("Failed to initialize BCA client: %v", err)
		return
	}

	ctx := context.Background()

	// Create Virtual Account using BCA endpoints
	createVAPayload := &types.CreateVAPayload{
		PartnerServiceId:   "BCA12345",
		CustomerNo:         "BCA_CUST001",
		VirtualAccountNo:   "8808001234567890",
		VirtualAccountName: "BCA Customer",
		VirtualAccountEmail: "customer@bca.co.id",
		TrxId:              "BCA-TRX-001",
		TotalAmount:        types.NewAmount(100000, "IDR"),
		ExpiredDate:        "2024-12-31T23:59:59+07:00",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE001",
			Channel:  "BCA_WEB",
		},
	}

	result, err := client.VirtualAccount().CreateVA(ctx, createVAPayload)
	if err != nil {
		log.Printf("BCA VA creation failed: %v", err)
	} else {
		fmt.Printf("BCA VA created successfully: %+v\n", result)
	}

	fmt.Printf("BCA Client Configuration: %+v\n", client.GetConfig().CustomEndpoints)
}

// demonstrateCustomEndpoints shows how to use completely custom endpoints
func demonstrateCustomEndpoints() {
	// Define custom endpoints for a specific bank or implementation
	customEndpoints := &snap.CustomEndpoints{
		B2BToken:   "/api/v2.0/auth/b2b-token",
		B2B2CToken: "/api/v2.0/auth/b2b2c-token",
		VirtualAccount: &snap.VirtualAccountEndpoints{
			CreateVA:     "/api/v2.0/custom-bank/va/create",
			UpdateVA:     "/api/v2.0/custom-bank/va/update",
			DeleteVA:     "/api/v2.0/custom-bank/va/delete",
			InquiryVA:    "/api/v2.0/custom-bank/va/inquiry-va",
			Inquiry:      "/api/v2.0/custom-bank/va/inquiry",
			Payment:      "/api/v2.0/custom-bank/va/payment",
			Status:       "/api/v2.0/custom-bank/va/status",
			Report:       "/api/v2.0/custom-bank/va/report",
			UpdateStatus: "/api/v2.0/custom-bank/va/update-status",
		},
		MPM: &snap.MPMEndpoints{
			Transfer:       "/api/v2.0/custom-bank/mpm/transfer",
			Inquiry:        "/api/v2.0/custom-bank/mpm/inquiry",
			Status:         "/api/v2.0/custom-bank/mpm/status",
			Refund:         "/api/v2.0/custom-bank/mpm/refund",
			BalanceInquiry: "/api/v2.0/custom-bank/mpm/balance",
			AccountInquiry: "/api/v2.0/custom-bank/mpm/account",
			History:        "/api/v2.0/custom-bank/mpm/history",
			GenerateQR:     "/api/v2.0/custom-bank/qr/generate",
			NotifyQR:       "/api/v2.0/custom-bank/qr/notify",
		},
	}

	client, err := snap.NewClient(snap.Config{
		BaseURL:         "https://custom-bank-api.example.com",
		ClientID:        getEnvOrDefault("CUSTOM_CLIENT_ID", "custom-client-id"),
		ClientSecret:    getEnvOrDefault("CUSTOM_CLIENT_SECRET", "custom-client-secret"),
		PrivateKeyPath:  getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:   getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:     "sandbox",
		LogLevel:        "info",
		CustomEndpoints: customEndpoints,
	})

	if err != nil {
		log.Printf("Failed to initialize custom client: %v", err)
		return
	}

	ctx := context.Background()

	// Test MPM balance inquiry with custom endpoint
	balancePayload := &types.MPMBalanceInquiryPayload{
		PartnerServiceId: "CUSTOM12345",
		CustomerNo:       "CUSTOM_CUST001",
		AccountNo:        "1234567890",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "CUSTOM_DEVICE001",
			Channel:  "CUSTOM_API",
		},
	}

	result, err := client.MPM().BalanceInquiry(ctx, balancePayload)
	if err != nil {
		log.Printf("Custom MPM balance inquiry failed: %v", err)
	} else {
		fmt.Printf("Custom MPM balance inquiry successful: %+v\n", result)
	}

	fmt.Printf("Custom Client Configuration: %+v\n", client.GetConfig().CustomEndpoints)
}

// demonstrateMixedConfiguration shows mixing bank presets with custom overrides
func demonstrateMixedConfiguration() {
	// Start with BNI preset but override some endpoints
	customEndpoints := &snap.CustomEndpoints{
		VirtualAccount: &snap.VirtualAccountEndpoints{
			CreateVA: "/api/v1.0/bni/custom-va/create", // Override only create endpoint
			Payment:  "/api/v1.0/bni/custom-va/payment", // Override only payment endpoint
			// Other endpoints will use BNI defaults
		},
		MPM: &snap.MPMEndpoints{
			Transfer: "/api/v1.0/bni/custom-mpm/transfer", // Override only transfer endpoint
			// Other endpoints will use BNI defaults
		},
	}

	client, err := snap.NewClientForBank(snap.Config{
		BaseURL:         getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:        getEnvOrDefault("BNI_CLIENT_ID", "bni-client-id"),
		ClientSecret:    getEnvOrDefault("BNI_CLIENT_SECRET", "bni-client-secret"),
		PrivateKeyPath:  getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:   getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:     "sandbox",
		LogLevel:        "info",
		CustomEndpoints: customEndpoints, // This will be merged with BNI presets
	}, "BNI")

	if err != nil {
		log.Printf("Failed to initialize mixed BNI client: %v", err)
		return
	}

	ctx := context.Background()

	// Test VA creation with mixed configuration
	createVAPayload := &types.CreateVAPayload{
		PartnerServiceId:   "BNI12345",
		CustomerNo:         "BNI_CUST001",
		VirtualAccountNo:   "8808001234567890",
		VirtualAccountName: "BNI Mixed Customer",
		TrxId:              "BNI-MIX-TRX-001",
		TotalAmount:        types.NewAmount(150000, "IDR"),
		ExpiredDate:        "2024-12-31T23:59:59+07:00",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BNI_MIX_DEVICE001",
			Channel:  "BNI_MIXED_WEB",
		},
	}

	result, err := client.VirtualAccount().CreateVA(ctx, createVAPayload)
	if err != nil {
		log.Printf("BNI mixed VA creation failed: %v", err)
	} else {
		fmt.Printf("BNI mixed VA created successfully: %+v\n", result)
	}

	fmt.Printf("Mixed BNI Client Configuration: %+v\n", client.GetConfig().CustomEndpoints)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}