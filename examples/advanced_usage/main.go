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
	// Advanced SDK usage examples
	fmt.Println("=== SNAP ASPI SDK Advanced Usage Examples ===")

	// Initialize client with environment variables
	client := initializeClientFromEnv()

	// Example 1: Error handling and retry logic
	fmt.Println("\n1. Error Handling and Retry Logic")
	demonstrateErrorHandling(client)

	// Example 2: Concurrent operations
	fmt.Println("\n2. Concurrent Operations")
	demonstrateConcurrentOperations(client)

	// Example 3: Context management and timeouts
	fmt.Println("\n3. Context Management and Timeouts")
	demonstrateContextManagement(client)

	// Example 4: Complete Virtual Account workflow
	fmt.Println("\n4. Complete Virtual Account Workflow")
	demonstrateVAWorkflow(client)

	// Example 5: MPM operations with error recovery
	fmt.Println("\n5. MPM Operations with Error Recovery")
	demonstrateMPMWorkflow(client)

	// Example 6: QR Code operations
	fmt.Println("\n6. QR Code Operations")
	demonstrateQRWorkflow(client)

	fmt.Println("\n=== Advanced Usage Examples Complete ===")
}

// initializeClientFromEnv initializes the client using environment variables
func initializeClientFromEnv() *snap.Client {
	config := snap.Config{
		BaseURL:        getEnvOrDefault("ASPI_BASE_URL", "https://sandbox.aspi-indonesia.or.id"),
		ClientID:       getEnvOrDefault("ASPI_CLIENT_ID", "test-client-id"),
		ClientSecret:   getEnvOrDefault("ASPI_CLIENT_SECRET", "test-client-secret"),
		PrivateKeyPath: getEnvOrDefault("ASPI_PRIVATE_KEY_PATH", "../../keys/private_key.pem"),
		PublicKeyPath:  getEnvOrDefault("ASPI_PUBLIC_KEY_PATH", "../../keys/public_key.pem"),
		Environment:    getEnvOrDefault("ASPI_ENVIRONMENT", "sandbox"),
		ConnectTimeout: 10,
		RequestTimeout: 30,
		MaxRetries:     3,
		LogLevel:       "info",
	}

	client, err := snap.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to initialize SNAP client: %v", err)
	}

	return client
}

// demonstrateErrorHandling shows comprehensive error handling
func demonstrateErrorHandling(client *snap.Client) {
	ctx := context.Background()

	// Intentionally create an invalid payload to demonstrate error handling
	invalidPayload := &types.CreateVAPayload{
		// Missing required fields to trigger validation error
		PartnerServiceId: "",
		CustomerNo:       "",
	}

	result, err := client.VirtualAccount().CreateVA(ctx, invalidPayload)
	if err != nil {
		fmt.Printf("Expected error caught: %v\n", err)
		
		// Handle different types of errors
		switch {
		case containsString(err.Error(), "validation"):
			fmt.Println("→ Handling validation error: Check required fields")
		case containsString(err.Error(), "authentication"):
			fmt.Println("→ Handling auth error: Token might be expired")
		case containsString(err.Error(), "network"):
			fmt.Println("→ Handling network error: Retry with backoff")
		default:
			fmt.Printf("→ Handling unknown error: %v\n", err)
		}
	} else {
		fmt.Printf("Unexpected success: %+v\n", result)
	}

	// Demonstrate retry logic
	fmt.Println("Implementing custom retry logic...")
	err = withRetry(func() error {
		// This would normally be a real operation
		_, err := client.Auth().GetB2BToken(ctx)
		return err
	}, 3, time.Second)

	if err != nil {
		fmt.Printf("Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("Operation succeeded with retry logic")
	}
}

// demonstrateConcurrentOperations shows how to handle concurrent API calls
func demonstrateConcurrentOperations(client *snap.Client) {
	ctx := context.Background()

	// Create multiple goroutines for concurrent operations
	type result struct {
		operation string
		success   bool
		error     error
	}

	results := make(chan result, 3)

	// Concurrent B2B token requests
	go func() {
		_, err := client.Auth().GetB2BToken(ctx)
		results <- result{"B2B Token", err == nil, err}
	}()

	// Concurrent VA inquiry (will likely fail without proper setup)
	go func() {
		payload := &types.InquiryVAPayload{
			PartnerServiceId: "12345",
			CustomerNo:       "CUST001",
			VirtualAccountNo: "8808001234567890",
			TrxDateInit:      "2024-01-01",
		}
		_, err := client.VirtualAccount().InquiryVA(ctx, payload)
		results <- result{"VA Inquiry", err == nil, err}
	}()

	// Concurrent MPM balance inquiry (will likely fail without proper setup)
	go func() {
		payload := &types.MPMBalanceInquiryPayload{
			PartnerServiceId: "12345",
			CustomerNo:       "CUST001",
			AccountNo:        "1234567890",
		}
		_, err := client.MPM().BalanceInquiry(ctx, payload)
		results <- result{"MPM Balance", err == nil, err}
	}()

	// Collect results
	for i := 0; i < 3; i++ {
		res := <-results
		status := "✓"
		if !res.success {
			status = "✗"
		}
		fmt.Printf("%s %s: %v\n", status, res.operation, res.error)
	}
}

// demonstrateContextManagement shows proper context usage
func demonstrateContextManagement(client *snap.Client) {
	// Context with timeout
	fmt.Println("Testing with 5-second timeout...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	_, err := client.Auth().GetB2BToken(ctx)
	duration := time.Since(start)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("Operation timed out after %v\n", duration)
		} else {
			fmt.Printf("Operation failed: %v (took %v)\n", err, duration)
		}
	} else {
		fmt.Printf("Operation succeeded in %v\n", duration)
	}

	// Context with cancellation
	fmt.Println("Testing with cancellation...")
	ctx2, cancel2 := context.WithCancel(context.Background())

	// Cancel after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		cancel2()
	}()

	start = time.Now()
	_, err = client.Auth().GetB2BToken(ctx2)
	duration = time.Since(start)

	if err != nil {
		if ctx2.Err() == context.Canceled {
			fmt.Printf("Operation cancelled after %v\n", duration)
		} else {
			fmt.Printf("Operation failed: %v (took %v)\n", err, duration)
		}
	} else {
		fmt.Printf("Operation completed before cancellation in %v\n", duration)
	}
}

// demonstrateVAWorkflow shows a complete Virtual Account workflow
func demonstrateVAWorkflow(client *snap.Client) {
	ctx := context.Background()

	// Step 1: Create Virtual Account
	fmt.Println("Step 1: Creating Virtual Account...")
	createPayload := &types.CreateVAPayload{
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

	createResult, err := client.VirtualAccount().CreateVA(ctx, createPayload)
	if err != nil {
		fmt.Printf("✗ Failed to create VA: %v\n", err)
		return
	}
	fmt.Printf("✓ VA Created: %+v\n", createResult)

	// Step 2: Inquiry Virtual Account
	fmt.Println("Step 2: Inquiring Virtual Account...")
	inquiryPayload := &types.InquiryVAPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      "2024-01-01",
	}

	inquiryResult, err := client.VirtualAccount().InquiryVA(ctx, inquiryPayload)
	if err != nil {
		fmt.Printf("✗ Failed to inquiry VA: %v\n", err)
	} else {
		fmt.Printf("✓ VA Inquiry: %+v\n", inquiryResult)
	}

	// Step 3: Check Status
	fmt.Println("Step 3: Checking Status...")
	statusPayload := &types.StatusPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		InquiryRequestId: "INQ-001",
	}

	statusResult, err := client.VirtualAccount().Status(ctx, statusPayload)
	if err != nil {
		fmt.Printf("✗ Failed to check status: %v\n", err)
	} else {
		fmt.Printf("✓ Status: %+v\n", statusResult)
	}
}

// demonstrateMPMWorkflow shows MPM operations with error recovery
func demonstrateMPMWorkflow(client *snap.Client) {
	ctx := context.Background()

	// Step 1: Check Balance
	fmt.Println("Step 1: Checking MPM Balance...")
	balancePayload := &types.MPMBalanceInquiryPayload{
		PartnerServiceId: "12345",
		CustomerNo:       "CUST001",
		AccountNo:        "1234567890",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	balanceResult, err := client.MPM().BalanceInquiry(ctx, balancePayload)
	if err != nil {
		fmt.Printf("✗ Failed to check balance: %v\n", err)
	} else {
		fmt.Printf("✓ Balance: %+v\n", balanceResult)
	}

	// Step 2: Transfer Credit
	fmt.Println("Step 2: Transferring Credit...")
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

	transferResult, err := client.MPM().Transfer(ctx, transferPayload)
	if err != nil {
		fmt.Printf("✗ Failed to transfer: %v\n", err)
		
		// Implement error recovery
		fmt.Println("Implementing error recovery...")
		if containsString(err.Error(), "insufficient") {
			fmt.Println("→ Insufficient balance, checking account status")
		} else if containsString(err.Error(), "network") {
			fmt.Println("→ Network error, will retry later")
		}
	} else {
		fmt.Printf("✓ Transfer: %+v\n", transferResult)
	}

	// Step 3: Check Transfer Status
	fmt.Println("Step 3: Checking Transfer Status...")
	statusPayload := &types.MPMStatusPayload{
		PartnerServiceId:   "12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF001",
		ServiceCode:        "MPM_TRANSFER",
		TransactionDate:    "2024-01-15",
		ExternalStoreId:    "STORE001",
	}

	statusResult, err := client.MPM().Status(ctx, statusPayload)
	if err != nil {
		fmt.Printf("✗ Failed to check transfer status: %v\n", err)
	} else {
		fmt.Printf("✓ Transfer Status: %+v\n", statusResult)
	}
}

// demonstrateQRWorkflow shows QR code operations
func demonstrateQRWorkflow(client *snap.Client) {
	ctx := context.Background()

	// Step 1: Generate QR Code
	fmt.Println("Step 1: Generating QR Code...")
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
		fmt.Printf("✗ Failed to generate QR: %v\n", err)
		return
	}
	fmt.Printf("✓ QR Generated: %s\n", qrResult.QRContent)
	fmt.Printf("  Valid Until: %s\n", qrResult.ValidUntil)

	// Step 2: Simulate QR Payment Notification
	fmt.Println("Step 2: Processing QR Payment Notification...")
	notifyPayload := &types.MPMNotifyQRPayload{
		PartnerServiceId:        "12345",
		CustomerNo:              "CUST001",
		PartnerReferenceNo:      "NOTIFY001",
		MerchantId:              "MERCHANT001",
		SubMerchantId:           "SUB001",
		ExternalStoreId:         "STORE001",
		OriginalReferenceNo:     "QR001",
		OriginalPartnerRefNo:    "PARTNER001",
		ServiceCode:             "QR_PAYMENT",
		TransactionDate:         "2024-01-01T10:00:00+07:00",
		Amount:                  types.NewAmount(25000, "IDR"),
		FeeAmount:               types.NewAmount(500, "IDR"),
		PaidAmount:              types.NewAmount(25500, "IDR"),
		CurrencyCode:            "IDR",
		TransactionStatus:       "SUCCESS",
		TransactionStatusDesc:   "Transaction successful",
		LatestTransactionStatus: "COMPLETED",
		PaymentMethod:           "QR_CODE",
		GatewayReferenceNo:      "GW001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "POS",
		},
	}

	notifyResult, err := client.MPM().NotifyQR(ctx, notifyPayload)
	if err != nil {
		fmt.Printf("✗ Failed to process QR notification: %v\n", err)
	} else {
		fmt.Printf("✓ QR Notification: %+v\n", notifyResult)
	}
}

// Helper functions

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			 s[len(s)-len(substr):] == substr ||
			 findSubstring(s, substr))))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func withRetry(operation func() error, maxRetries int, delay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		
		// Check if error is retryable
		if containsString(err.Error(), "network") || 
		   containsString(err.Error(), "timeout") ||
		   containsString(err.Error(), "connection") {
			if i < maxRetries-1 {
				fmt.Printf("Retry %d/%d after %v: %v\n", i+1, maxRetries, delay, err)
				time.Sleep(delay * time.Duration(i+1)) // Exponential backoff
				continue
			}
		}
		
		// Non-retryable error or max retries reached
		break
	}
	return err
}