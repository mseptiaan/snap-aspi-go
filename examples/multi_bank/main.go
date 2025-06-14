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
	log.Println("=== Multi-Bank Integration Example ===")

	// Example 1: BCA Integration
	log.Println("\n1. BCA Bank Integration")
	demonstrateBCAIntegration()

	// Example 2: BNI Integration
	log.Println("\n2. BNI Bank Integration")
	demonstrateBNIIntegration()

	// Example 3: Multi-bank support
	log.Println("\n3. Multi-Bank Support")
	demonstrateMultiBankSupport()

	log.Println("\n=== Multi-Bank Examples Complete ===")
}

func demonstrateBCAIntegration() {
	// BCA-specific configuration
	client, err := snap.NewClientForBank(snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("BCA_CLIENT_ID", "bca-client-id"),
		ClientSecret:   getEnvOrDefault("BCA_CLIENT_SECRET", "bca-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:    "sandbox",
		LogLevel:       "info",
	}, "BCA")

	if err != nil {
		log.Printf("Failed to initialize BCA client: %v", err)
		return
	}

	ctx := context.Background()

	// BCA Virtual Account creation
	vaPayload := &types.CreateVAPayload{
		PartnerServiceId:    "BCA001",
		CustomerNo:          "BCA_CUSTOMER_001",
		VirtualAccountNo:    "8808001234567890",
		VirtualAccountName:  "BCA Customer Name",
		VirtualAccountEmail: "customer@bca.co.id",
		VirtualAccountPhone: "081234567890",
		TrxId:               "BCA_TRX_001",
		TotalAmount:         types.NewAmount(500000, "IDR"),
		ExpiredDate:         "2024-12-31T23:59:59+07:00",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_INTERNET_BANKING",
		},
	}

	result, err := client.VirtualAccount().CreateVA(ctx, vaPayload)
	if err != nil {
		log.Printf("BCA VA creation failed: %v", err)
	} else {
		fmt.Printf("✓ BCA VA created: %+v\n", result)
	}

	// BCA MPM Transfer
	mpmPayload := &types.MPMTransferPayload{
		PartnerServiceId:      "BCA001",
		CustomerNo:            "BCA_CUSTOMER_001",
		PartnerReferenceNo:    "BCA_REF_001",
		MerchantId:            "BCA_MERCHANT_001",
		SubMerchantId:         "BCA_SUB_001",
		ExternalStoreId:       "BCA_STORE_001",
		Amount:                types.NewAmount(250000, "IDR"),
		BeneficiaryAccountNo:  "1234567890",
		BeneficiaryName:       "BCA Beneficiary",
		BeneficiaryCustomerNo: "BCA_BENEFICIARY_001",
		BeneficiaryEmail:      "beneficiary@bca.co.id",
		Currency:              "IDR",
		CustomerReference:     "BCA_CUST_REF_001",
		FeeType:               "OUR",
		Remark:                "BCA Transfer",
		SourceAccountNo:       "0987654321",
		TransactionDate:       "2024-01-15T10:30:00+07:00",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_MOBILE",
		},
	}

	transferResult, err := client.MPM().Transfer(ctx, mpmPayload)
	if err != nil {
		log.Printf("BCA MPM transfer failed: %v", err)
	} else {
		fmt.Printf("✓ BCA MPM transfer: %+v\n", transferResult)
	}
}

func demonstrateBNIIntegration() {
	// BNI-specific configuration
	client, err := snap.NewClientForBank(snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("BNI_CLIENT_ID", "bni-client-id"),
		ClientSecret:   getEnvOrDefault("BNI_CLIENT_SECRET", "bni-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:    "sandbox",
		LogLevel:       "info",
	}, "BNI")

	if err != nil {
		log.Printf("Failed to initialize BNI client: %v", err)
		return
	}

	ctx := context.Background()

	// BNI QR Code generation
	qrPayload := &types.MPMGenerateQRPayload{
		PartnerServiceId:   "BNI001",
		CustomerNo:         "BNI_CUSTOMER_001",
		PartnerReferenceNo: "BNI_QR_001",
		MerchantId:         "BNI_MERCHANT_001",
		SubMerchantId:      "BNI_SUB_001",
		StoreId:            "BNI_STORE_001",
		Amount:             types.NewAmount(100000, "IDR"),
		MerchantName:       "BNI Test Merchant",
		MerchantCity:       "Jakarta",
		MerchantPostalCode: "12345",
		TerminalLabel:      "BNI Terminal 1",
		PurposeOfPayment:   "BNI Payment",
		QRType:             "DYNAMIC",
		ValidityPeriod:     "3600",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BNI_DEVICE_001",
			Channel:  "BNI_POS",
		},
	}

	qrResult, err := client.MPM().GenerateQR(ctx, qrPayload)
	if err != nil {
		log.Printf("BNI QR generation failed: %v", err)
	} else {
		fmt.Printf("✓ BNI QR generated: %s\n", qrResult.QRContent)
		fmt.Printf("  Valid until: %s\n", qrResult.ValidUntil)
	}

	// BNI Balance inquiry
	balancePayload := &types.MPMBalanceInquiryPayload{
		PartnerServiceId: "BNI001",
		CustomerNo:       "BNI_CUSTOMER_001",
		AccountNo:        "1234567890",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BNI_DEVICE_001",
			Channel:  "BNI_MOBILE",
		},
	}

	balanceResult, err := client.MPM().BalanceInquiry(ctx, balancePayload)
	if err != nil {
		log.Printf("BNI balance inquiry failed: %v", err)
	} else {
		fmt.Printf("✓ BNI balance inquiry: %+v\n", balanceResult)
	}
}

func demonstrateMultiBankSupport() {
	// Multi-bank application supporting different banks
	banks := []BankConfig{
		{
			Code:         "BCA",
			Name:         "Bank Central Asia",
			ClientID:     getEnvOrDefault("BCA_CLIENT_ID", "bca-client-id"),
			ClientSecret: getEnvOrDefault("BCA_CLIENT_SECRET", "bca-client-secret"),
		},
		{
			Code:         "BNI",
			Name:         "Bank Negara Indonesia",
			ClientID:     getEnvOrDefault("BNI_CLIENT_ID", "bni-client-id"),
			ClientSecret: getEnvOrDefault("BNI_CLIENT_SECRET", "bni-client-secret"),
		},
		{
			Code:         "BRI",
			Name:         "Bank Rakyat Indonesia",
			ClientID:     getEnvOrDefault("BRI_CLIENT_ID", "bri-client-id"),
			ClientSecret: getEnvOrDefault("BRI_CLIENT_SECRET", "bri-client-secret"),
		},
	}

	// Create clients for each bank
	clients := make(map[string]*snap.Client)

	for _, bank := range banks {
		client, err := snap.NewClientForBank(snap.Config{
			BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
			ClientID:       bank.ClientID,
			ClientSecret:   bank.ClientSecret,
			PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
			PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
			Environment:    "sandbox",
			LogLevel:       "info",
		}, bank.Code)

		if err != nil {
			log.Printf("Failed to initialize %s client: %v", bank.Name, err)
			continue
		}

		clients[bank.Code] = client
		fmt.Printf("✓ %s client initialized\n", bank.Name)
	}

	// Test each bank's authentication
	ctx := context.Background()
	for bankCode, client := range clients {
		_, err := client.Auth().GetB2BToken(ctx)
		if err != nil {
			log.Printf("✗ %s authentication failed: %v", bankCode, err)
		} else {
			fmt.Printf("✓ %s authentication successful\n", bankCode)
		}
	}

	// Example: Route transaction to specific bank based on account number
	accountNumber := "8808001234567890"
	selectedBank := routeToBank(accountNumber)

	if client, exists := clients[selectedBank]; exists {
		fmt.Printf("Routing transaction to %s bank\n", selectedBank)

		// Process transaction with selected bank
		vaPayload := &types.CreateVAPayload{
			PartnerServiceId:   selectedBank + "001",
			CustomerNo:         selectedBank + "_CUSTOMER_001",
			VirtualAccountNo:   accountNumber,
			VirtualAccountName: "Multi-bank Customer",
			TrxId:              selectedBank + "_TRX_001",
			TotalAmount:        types.NewAmount(300000, "IDR"),
			ExpiredDate:        "2024-12-31T23:59:59+07:00",
			AdditionalInfo: &types.AdditionalInfo{
				DeviceId: selectedBank + "_DEVICE_001",
				Channel:  "MULTI_BANK_APP",
			},
		}

		result, err := client.VirtualAccount().CreateVA(ctx, vaPayload)
		if err != nil {
			log.Printf("Multi-bank VA creation failed: %v", err)
		} else {
			fmt.Printf("✓ Multi-bank VA created via %s: %+v\n", selectedBank, result)
		}
	}
}

// BankConfig represents bank configuration
type BankConfig struct {
	Code         string
	Name         string
	ClientID     string
	ClientSecret string
}

// routeToBank determines which bank to use based on account number
func routeToBank(accountNumber string) string {
	// Simple routing logic based on account number prefix
	// In real implementation, this would be more sophisticated
	switch {
	case len(accountNumber) >= 4 && accountNumber[:4] == "8808":
		return "BCA"
	case len(accountNumber) >= 4 && accountNumber[:4] == "8809":
		return "BNI"
	case len(accountNumber) >= 4 && accountNumber[:4] == "8810":
		return "BRI"
	default:
		return "BCA" // Default to BCA
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
