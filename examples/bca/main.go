package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

func main() {
	fmt.Println("=== BCA SNAP Integration Example ===")

	// Initialize the BCA client
	client, err := snap.NewBCAClient(snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("BCA_CLIENT_ID", "your-bca-client-id"),
		ClientSecret:   getEnvOrDefault("BCA_CLIENT_SECRET", "your-bca-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "keys/public_key.pem"),
		Environment:    "sandbox",
		LogLevel:       "info",
	})
	if err != nil {
		log.Fatalf("Failed to initialize BCA client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Create BCA Virtual Account
	fmt.Println("\n1. Create BCA Virtual Account")
	createBCAVirtualAccount(ctx, client)

	// Example 2: Inquiry BCA Virtual Account
	fmt.Println("\n2. Inquiry BCA Virtual Account")
	inquiryBCAVirtualAccount(ctx, client)

	// Example 3: BCA Balance Inquiry
	fmt.Println("\n3. BCA Balance Inquiry")
	bcaBalanceInquiry(ctx, client)

	// Example 4: Generate BCA QR Code
	fmt.Println("\n4. Generate BCA QR Code")
	generateBCAQR(ctx, client)

	// Example 5: BCA Transfer
	fmt.Println("\n5. BCA Transfer")
	bcaTransfer(ctx, client)

	fmt.Println("\n=== BCA Integration Examples Complete ===")
}

func createBCAVirtualAccount(ctx context.Context, client *snap.BCAClient) {
	// Create BCA Virtual Account
	payload := &types.BCACreateVAPayload{
		PartnerServiceId:    "BCA12345",
		CustomerNo:          "BCA_CUSTOMER_001",
		VirtualAccountNo:    "8808001234567890",
		VirtualAccountName:  "BCA Customer",
		VirtualAccountEmail: "customer@example.com",
		VirtualAccountPhone: "081234567890",
		TrxId:               "BCA_TRX_001",
		TotalAmount:         types.NewAmount(150000, "IDR"),
		ExpiredDate:         time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // 1 month from now
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_WEB",
		},
		// BCA-specific fields
		BcaCustomerNo:      "BCA123456789",
		BcaSubCompany:      "BCA Finance",
		BcaTransactionType: "1",
	}

	result, err := client.CreateBCAVirtualAccount(ctx, payload)
	if err != nil {
		fmt.Printf("Create BCA VA failed: %v\n", err)
		return
	}

	fmt.Printf("BCA Virtual Account created successfully\n")
	fmt.Printf("Response Code: %s\n", result.ResponseCode)
	fmt.Printf("Response Message: %s\n", result.ResponseMessage)
	fmt.Printf("Reference No: %s\n", result.ReferenceNo)
	fmt.Printf("Virtual Account No: %s\n", result.VirtualAccountNo)
	if result.BcaReferenceNo != "" {
		fmt.Printf("BCA Reference No: %s\n", result.BcaReferenceNo)
	}
}

func inquiryBCAVirtualAccount(ctx context.Context, client *snap.BCAClient) {
	// Inquiry BCA Virtual Account
	payload := &types.BCAInquiryVAPayload{
		PartnerServiceId: "BCA12345",
		CustomerNo:       "BCA_CUSTOMER_001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      time.Now().Format("2006-01-02"),
		// BCA-specific fields
		BcaCustomerNo: "BCA123456789",
	}

	result, err := client.InquiryBCAVirtualAccount(ctx, payload)
	if err != nil {
		fmt.Printf("Inquiry BCA VA failed: %v\n", err)
		return
	}

	fmt.Printf("BCA Virtual Account inquiry successful\n")
	fmt.Printf("Response Code: %s\n", result.ResponseCode)
	fmt.Printf("Response Message: %s\n", result.ResponseMessage)
	fmt.Printf("Virtual Account No: %s\n", result.VirtualAccountNo)
	fmt.Printf("Virtual Account Name: %s\n", result.VirtualAccountName)
	fmt.Printf("Expired Date: %s\n", result.ExpiredDate)
	if result.BcaReferenceNo != "" {
		fmt.Printf("BCA Reference No: %s\n", result.BcaReferenceNo)
	}
	if result.BcaCustomerNo != "" {
		fmt.Printf("BCA Customer No: %s\n", result.BcaCustomerNo)
	}
	if result.BcaSubCompany != "" {
		fmt.Printf("BCA Sub Company: %s\n", result.BcaSubCompany)
	}
	if result.BcaTransactionType != "" {
		fmt.Printf("BCA Transaction Type: %s\n", result.BcaTransactionType)
	}
}

func bcaBalanceInquiry(ctx context.Context, client *snap.BCAClient) {
	// BCA Balance Inquiry
	payload := &types.BCABalanceInquiryPayload{
		PartnerServiceId: "BCA12345",
		CustomerNo:       "BCA_CUSTOMER_001",
		AccountNo:        "1234567890",
		BalanceTypes:     []string{"Cash", "Coins"},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_MOBILE",
		},
		// BCA-specific fields
		BcaCustomerNo: "BCA123456789",
	}

	result, err := client.BCABalanceInquiry(ctx, payload)
	if err != nil {
		fmt.Printf("BCA balance inquiry failed: %v\n", err)
		return
	}

	fmt.Printf("BCA balance inquiry successful\n")
	fmt.Printf("Response Code: %s\n", result.ResponseCode)
	fmt.Printf("Response Message: %s\n", result.ResponseMessage)
	fmt.Printf("Account No: %s\n", result.AccountNo)
	fmt.Printf("Account Name: %s\n", result.Name)
	
	// Display account balances
	if len(result.AccountInfos) > 0 {
		fmt.Println("Account Balances:")
		for i, info := range result.AccountInfos {
			fmt.Printf("  %d. Type: %s\n", i+1, info.BalanceType)
			fmt.Printf("     Available Balance: %s\n", info.AvailableBalance.String())
			fmt.Printf("     Ledger Balance: %s\n", info.LedgerBalance.String())
		}
	}
	
	if result.BcaAccountType != "" {
		fmt.Printf("BCA Account Type: %s\n", result.BcaAccountType)
	}
}

func generateBCAQR(ctx context.Context, client *snap.BCAClient) {
	// Generate BCA QR Code
	payload := &types.BCAGenerateQRPayload{
		PartnerServiceId:   "BCA12345",
		CustomerNo:         "BCA_CUSTOMER_001",
		PartnerReferenceNo: "BCA_QR_001",
		Amount:             types.NewAmount(50000, "IDR"),
		MerchantId:         "BCA_MERCHANT_001",
		MerchantName:       "BCA Test Merchant",
		ValidityPeriod:     "3600", // 1 hour
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_POS",
		},
		// BCA-specific fields
		BcaQRType:        "DYNAMIC",
		BcaMerchantCity:  "Jakarta",
		BcaTerminalLabel: "BCA Terminal 1",
	}

	result, err := client.GenerateBCAQR(ctx, payload)
	if err != nil {
		fmt.Printf("Generate BCA QR failed: %v\n", err)
		return
	}

	fmt.Printf("BCA QR Code generated successfully\n")
	fmt.Printf("Response Code: %s\n", result.ResponseCode)
	fmt.Printf("Response Message: %s\n", result.ResponseMessage)
	fmt.Printf("Reference No: %s\n", result.ReferenceNo)
	fmt.Printf("QR Content: %s\n", result.QRContent)
	fmt.Printf("QR URL: %s\n", result.QRUrl)
	fmt.Printf("Valid Until: %s\n", result.ValidUntil)
	if result.BcaQRId != "" {
		fmt.Printf("BCA QR ID: %s\n", result.BcaQRId)
	}
}

func bcaTransfer(ctx context.Context, client *snap.BCAClient) {
	// BCA Transfer
	payload := &types.BCATransferPayload{
		PartnerServiceId:       "BCA12345",
		CustomerNo:             "BCA_CUSTOMER_001",
		PartnerReferenceNo:     "BCA_TRF_001",
		Amount:                 types.NewAmount(75000, "IDR"),
		BeneficiaryAccountNo:   "1234567890",
		BeneficiaryAccountName: "Jane Doe",
		BeneficiaryBankCode:    "014",
		SourceAccountNo:        "0987654321",
		TransactionDate:        time.Now().Format(time.RFC3339),
		FeeType:                "OUR",
		Remark:                 "BCA Transfer Test",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "BCA_DEVICE_001",
			Channel:  "BCA_MOBILE",
		},
		// BCA-specific fields
		BcaTransferType: "RTGS",
		BcaPriority:     "H",
	}

	result, err := client.BCATransfer(ctx, payload)
	if err != nil {
		fmt.Printf("BCA transfer failed: %v\n", err)
		return
	}

	fmt.Printf("BCA transfer successful\n")
	fmt.Printf("Response Code: %s\n", result.ResponseCode)
	fmt.Printf("Response Message: %s\n", result.ResponseMessage)
	fmt.Printf("Reference No: %s\n", result.ReferenceNo)
	if result.BcaReferenceNo != "" {
		fmt.Printf("BCA Reference No: %s\n", result.BcaReferenceNo)
	}
	if result.TransactionDate != "" {
		fmt.Printf("Transaction Date: %s\n", result.TransactionDate)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}