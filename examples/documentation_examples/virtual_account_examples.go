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
	fmt.Println("=== SNAP ASPI SDK Virtual Account Examples ===")

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

	// Example 1: Create Virtual Account
	fmt.Println("\n1. Create Virtual Account")
	createVirtualAccount(ctx, client)

	// Example 2: Update Virtual Account
	fmt.Println("\n2. Update Virtual Account")
	updateVirtualAccount(ctx, client)

	// Example 3: Inquiry Virtual Account
	fmt.Println("\n3. Inquiry Virtual Account")
	inquiryVirtualAccount(ctx, client)

	// Example 4: Process Payment
	fmt.Println("\n4. Process Payment")
	processPayment(ctx, client)

	// Example 5: Check Status
	fmt.Println("\n5. Check Status")
	checkStatus(ctx, client)

	// Example 6: Generate Report
	fmt.Println("\n6. Generate Report")
	generateReport(ctx, client)

	// Example 7: Delete Virtual Account
	fmt.Println("\n7. Delete Virtual Account")
	deleteVirtualAccount(ctx, client)

	fmt.Println("\n=== Virtual Account Examples Finished ===")
}

func createVirtualAccount(ctx context.Context, client *snap.Client) {
	// Create Virtual Account
	vaPayload := &types.CreateVAPayload{
		PartnerServiceId:    "VA12345",
		CustomerNo:          "CUST001",
		VirtualAccountNo:    "8808001234567890",
		VirtualAccountName:  "John Doe",
		VirtualAccountEmail: "john@example.com",
		VirtualAccountPhone: "081234567890",
		TrxId:               "TRX001",
		TotalAmount:         types.NewAmount(150000, "IDR"),
		BillDetails: []types.BillDetail{
			{
				BillCode:        "BILL001",
				BillNo:          "INV001",
				BillName:        "Invoice January 2024",
				BillShortName:   "INV-JAN",
				BillDescription: map[string]string{"en": "January Invoice", "id": "Faktur Januari"},
				BillSubCompany:  "Finance",
				BillAmount:      types.NewAmount(150000, "IDR"),
				AdditionalInfo:  map[string]any{"priority": "high"},
			},
		},
		FreeTexts: []types.FreeText{
			{
				English:   "Thank you for your payment",
				Indonesia: "Terima kasih atas pembayaran Anda",
			},
		},
		VirtualAccountTrxType: "1",
		FeeAmount:             types.NewAmount(0, "IDR"),
		ExpiredDate:           time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // 1 month from now
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "WEB",
		},
	}

	result, err := client.VirtualAccount().CreateVA(ctx, vaPayload)
	if err != nil {
		fmt.Printf("Create VA failed: %v\n", err)
		return
	}

	fmt.Printf("Virtual Account created successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
	fmt.Printf("Virtual Account No: %s\n", result["virtualAccountNo"])
}

func updateVirtualAccount(ctx context.Context, client *snap.Client) {
	// Update Virtual Account
	vaPayload := &types.CreateVAPayload{
		PartnerServiceId:    "VA12345",
		CustomerNo:          "CUST001",
		VirtualAccountNo:    "8808001234567890",
		VirtualAccountName:  "John Doe Updated",
		VirtualAccountEmail: "john.updated@example.com",
		VirtualAccountPhone: "081234567890",
		TrxId:               "TRX001",
		TotalAmount:         types.NewAmount(200000, "IDR"),
		ExpiredDate:         time.Now().AddDate(0, 2, 0).Format(time.RFC3339), // 2 months from now
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "WEB",
		},
	}

	result, err := client.VirtualAccount().UpdateVA(ctx, vaPayload)
	if err != nil {
		fmt.Printf("Update VA failed: %v\n", err)
		return
	}

	fmt.Printf("Virtual Account updated successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
}

func inquiryVirtualAccount(ctx context.Context, client *snap.Client) {
	// Inquiry Virtual Account
	inquiryPayload := &types.InquiryVAPayload{
		PartnerServiceId: "VA12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      time.Now().Format("2006-01-02"),
	}

	result, err := client.VirtualAccount().InquiryVA(ctx, inquiryPayload)
	if err != nil {
		fmt.Printf("Inquiry VA failed: %v\n", err)
		return
	}

	fmt.Printf("Virtual Account inquiry successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Virtual Account Name: %s\n", result["virtualAccountName"])
	fmt.Printf("Total Amount: %s\n", result["totalAmount"])
	fmt.Printf("Expired Date: %s\n", result["expiredDate"])
}

func processPayment(ctx context.Context, client *snap.Client) {
	// Process Payment
	paymentPayload := &types.PaymentPayload{
		PartnerServiceId:        "VA12345",
		CustomerNo:              "CUST001",
		VirtualAccountNo:        "8808001234567890",
		VirtualAccountName:      "John Doe",
		VirtualAccountEmail:     "john@example.com",
		VirtualAccountPhone:     "081234567890",
		TrxId:                   "TRX001",
		PaymentRequestId:        "PAY001",
		ChannelCode:             "6014",
		SourceAccountNo:         "1234567890",
		PaidAmount:              types.NewAmount(150000, "IDR"),
		CumulativePaymentAmount: types.NewAmount(150000, "IDR"),
		PaidBills:               "BILL001",
		TotalAmount:             types.NewAmount(150000, "IDR"),
		TrxDateTime:             time.Now().Format(time.RFC3339),
		ReferenceNo:             "REF001",
		JournalNum:              "JRN001",
		PaymentType:             "1",
		FlagAdvise:              "N",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "ATM",
		},
	}

	result, err := client.VirtualAccount().Payment(ctx, paymentPayload)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}

	fmt.Printf("Payment processed successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
}

func checkStatus(ctx context.Context, client *snap.Client) {
	// Check Status
	statusPayload := &types.StatusPayload{
		PartnerServiceId: "VA12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		InquiryRequestId: "INQ001",
		PaymentRequestId: "PAY001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "WEB",
		},
	}

	result, err := client.VirtualAccount().Status(ctx, statusPayload)
	if err != nil {
		fmt.Printf("Status check failed: %v\n", err)
		return
	}

	fmt.Printf("Status check successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Transaction Status: %s\n", result["transactionStatus"])
}

func generateReport(ctx context.Context, client *snap.Client) {
	// Generate Report
	reportPayload := &types.ReportPayload{
		PartnerServiceId: "VA12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		TrxDateEnd:       time.Now().Format("2006-01-02"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "WEB",
		},
	}

	result, err := client.VirtualAccount().Report(ctx, reportPayload)
	if err != nil {
		fmt.Printf("Report generation failed: %v\n", err)
		return
	}

	fmt.Printf("Report generated successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Total Transactions: %v\n", result["totalTransactions"])
}

func deleteVirtualAccount(ctx context.Context, client *snap.Client) {
	// Delete Virtual Account
	deletePayload := &types.InquiryVAPayload{
		PartnerServiceId: "VA12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      time.Now().Format("2006-01-02"),
	}

	result, err := client.VirtualAccount().DeleteVA(ctx, deletePayload)
	if err != nil {
		fmt.Printf("Delete VA failed: %v\n", err)
		return
	}

	fmt.Printf("Virtual Account deleted successfully\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}