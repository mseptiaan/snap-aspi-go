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
	fmt.Println("=== SNAP ASPI SDK Transfer Credit Examples ===")

	// Initialize the SDK client
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

	// Example 1: Account Inquiry
	fmt.Println("\n1. Account Inquiry")
	accountInquiry(ctx, client)

	// Example 2: Trigger Transfer
	fmt.Println("\n2. Trigger Transfer")
	triggerTransfer(ctx, client)

	// Example 3: Transfer Status Inquiry
	fmt.Println("\n3. Transfer Status Inquiry")
	transferStatusInquiry(ctx, client)

	// Example 4: Customer Top Up
	fmt.Println("\n4. Customer Top Up")
	customerTopUp(ctx, client)

	// Example 5: Bulk Cashin
	fmt.Println("\n5. Bulk Cashin")
	bulkCashin(ctx, client)

	// Example 6: Transfer To Bank
	fmt.Println("\n6. Transfer To Bank")
	transferToBank(ctx, client)

	// Example 7: Transfer To OTC
	fmt.Println("\n7. Transfer To OTC")
	transferToOTC(ctx, client)

	fmt.Println("\n=== Transfer Credit Examples Finished ===")
}

func accountInquiry(ctx context.Context, client *snap.Client) {
	// Account Inquiry
	inquiryPayload := &types.AccountInquiryPayload{
		PartnerServiceId:     "TRF12345",
		CustomerNo:           "CUST001",
		PartnerReferenceNo:   "REF001",
		BeneficiaryAccountNo: "1234567890",
		BeneficiaryBankCode:  "014",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().AccountInquiry(ctx, inquiryPayload)
	if err != nil {
		fmt.Printf("Account inquiry failed: %v\n", err)
		return
	}

	fmt.Printf("Account inquiry successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Beneficiary Account Name: %s\n", result["beneficiaryAccountName"])
	fmt.Printf("Beneficiary Account No: %s\n", result["beneficiaryAccountNo"])
	fmt.Printf("Beneficiary Bank Name: %s\n", result["beneficiaryBankName"])
}

func triggerTransfer(ctx context.Context, client *snap.Client) {
	// Trigger Transfer
	transferPayload := &types.TriggerTransferPayload{
		PartnerServiceId:       "TRF12345",
		CustomerNo:             "CUST001",
		PartnerReferenceNo:     "REF002",
		Amount:                 types.NewAmount(50000, "IDR"),
		BeneficiaryAccountNo:   "1234567890",
		BeneficiaryAccountName: "Jane Doe",
		BeneficiaryBankCode:    "014",
		BeneficiaryBankName:    "BCA",
		SourceAccountNo:        "0987654321",
		TransactionDate:        time.Now().Format(time.RFC3339),
		FeeType:                "OUR",
		Remark:                 "Transfer for payment",
		Currency:               "IDR",
		CustomerReference:      "CUSTREF001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().TriggerTransfer(ctx, transferPayload)
	if err != nil {
		fmt.Printf("Trigger transfer failed: %v\n", err)
		return
	}

	fmt.Printf("Transfer triggered successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
}

func transferStatusInquiry(ctx context.Context, client *snap.Client) {
	// Transfer Status Inquiry
	statusPayload := &types.TransferStatusInquiryPayload{
		PartnerServiceId:     "TRF12345",
		CustomerNo:           "CUST001",
		OriginalPartnerRefNo: "REF002",
		OriginalReferenceNo:  "REF002",
		OriginalExternalId:   "EXT001",
		ServiceCode:          "17",
		TransactionDate:      time.Now().Format(time.RFC3339),
		Amount:               types.NewAmount(50000, "IDR"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().TransferStatusInquiry(ctx, statusPayload)
	if err != nil {
		fmt.Printf("Transfer status inquiry failed: %v\n", err)
		return
	}

	fmt.Printf("Transfer status inquiry successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Latest Transaction Status: %s\n", result["latestTransactionStatus"])
	fmt.Printf("Transaction Status Description: %s\n", result["transactionStatusDesc"])
}

func customerTopUp(ctx context.Context, client *snap.Client) {
	// Customer Top Up
	topupPayload := &types.CustomerTopUpPayload{
		PartnerServiceId: "TOP12345",
		CustomerNo:       "CUST001",
		CustomerName:     "John Doe",
		Amount:           types.NewAmount(100000, "IDR"),
		FeeAmount:        types.NewAmount(0, "IDR"),
		TransactionDate:  time.Now().Format(time.RFC3339),
		SessionId:        "SESSION001",
		CategoryId:       1,
		Notes:            "Top up e-wallet",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().CustomerTopUp(ctx, topupPayload)
	if err != nil {
		fmt.Printf("Customer top up failed: %v\n", err)
		return
	}

	fmt.Printf("Customer top up successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
	fmt.Printf("Reference Number: %s\n", result["referenceNumber"])
}

func bulkCashin(ctx context.Context, client *snap.Client) {
	// Bulk Cashin
	bulkPayload := &types.BulkCashinPayload{
		PartnerServiceId: "BULK12345",
		CustomerNo:       "CUST001",
		PartnerBulkId:    "BULK001",
		TransactionDate:  time.Now().Format(time.RFC3339),
		Currency:         "IDR",
		BulkObject: []types.BulkItem{
			{
				AccountNumber:      "6281234567890",
				AccountName:        "User 1",
				Amount:             types.NewAmount(50000, "IDR"),
				PartnerReferenceNo: "BULK001-1",
				AdditionalInfo: &types.AdditionalInfo{
					DeviceId: "DEVICE001",
					Channel:  "BATCH",
				},
			},
			{
				AccountNumber:      "6289876543210",
				AccountName:        "User 2",
				Amount:             types.NewAmount(75000, "IDR"),
				PartnerReferenceNo: "BULK001-2",
				AdditionalInfo: &types.AdditionalInfo{
					DeviceId: "DEVICE001",
					Channel:  "BATCH",
				},
			},
		},
		FeeType: "OUR",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "BATCH",
		},
	}

	result, err := client.TransferCredit().BulkCashin(ctx, bulkPayload)
	if err != nil {
		fmt.Printf("Bulk cashin failed: %v\n", err)
		return
	}

	fmt.Printf("Bulk cashin successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Bulk ID: %s\n", result["bulkId"])
}

func transferToBank(ctx context.Context, client *snap.Client) {
	// Transfer To Bank
	bankPayload := &types.TransferToBankPayload{
		PartnerServiceId:         "BANK12345",
		CustomerNo:               "CUST001",
		AccountType:              "SAVINGS",
		BeneficiaryAccountNumber: "1234567890",
		BeneficiaryBankCode:      "014",
		Amount:                   types.NewAmount(200000, "IDR"),
		SessionId:                "SESSION002",
		FeeType:                  "OUR",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().TransferToBank(ctx, bankPayload)
	if err != nil {
		fmt.Printf("Transfer to bank failed: %v\n", err)
		return
	}

	fmt.Printf("Transfer to bank successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
	fmt.Printf("Transaction Date: %s\n", result["transactionDate"])
}

func transferToOTC(ctx context.Context, client *snap.Client) {
	// Transfer To OTC
	otcPayload := &types.TransferToOTCPayload{
		PartnerServiceId: "OTC12345",
		CustomerNo:       "CUST001",
		OTP:              "123456",
		Amount:           types.NewAmount(300000, "IDR"),
		FeeType:          "OUR",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferCredit().TransferToOTC(ctx, otcPayload)
	if err != nil {
		fmt.Printf("Transfer to OTC failed: %v\n", err)
		return
	}

	fmt.Printf("Transfer to OTC successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
	fmt.Printf("Transaction Date: %s\n", result["transactionDate"])
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}