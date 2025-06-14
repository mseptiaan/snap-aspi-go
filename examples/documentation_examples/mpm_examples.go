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
	fmt.Println("=== SNAP ASPI SDK MPM Examples ===")

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

	// Example 1: Generate QR Code
	fmt.Println("\n1. Generate QR Code")
	generateQRCode(ctx, client)

	// Example 2: Decode QR Code
	fmt.Println("\n2. Decode QR Code")
	decodeQRCode(ctx, client)

	// Example 3: Apply OTT
	fmt.Println("\n3. Apply OTT")
	applyOTT(ctx, client)

	// Example 4: Process Payment
	fmt.Println("\n4. Process Payment")
	processPayment(ctx, client)

	// Example 5: Query Payment
	fmt.Println("\n5. Query Payment")
	queryPayment(ctx, client)

	// Example 6: Cancel Payment
	fmt.Println("\n6. Cancel Payment")
	cancelPayment(ctx, client)

	// Example 7: Refund Payment
	fmt.Println("\n7. Refund Payment")
	refundPayment(ctx, client)

	// Example 8: MPM Transfer
	fmt.Println("\n8. MPM Transfer")
	mpmTransfer(ctx, client)

	fmt.Println("\n=== MPM Examples Finished ===")
}

func generateQRCode(ctx context.Context, client *snap.Client) {
	// Generate QR Code
	qrPayload := &types.MPMGenerateQRPayload{
		PartnerServiceId:   "MPM12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "QR001",
		Amount:             types.NewAmount(25000, "IDR"),
		FeeAmount:          types.NewAmount(0, "IDR"),
		MerchantId:         "MERCH001",
		SubMerchantId:      "SUB001",
		StoreId:            "STORE001",
		TerminalId:         "TERM001",
		ValidityPeriod:     "3600", // 1 hour
		MerchantName:       "Test Merchant",
		MerchantCity:       "Jakarta",
		MerchantPostalCode: "12345",
		TerminalLabel:      "Terminal 1",
		PurposeOfPayment:   "Payment for goods",
		QRType:             "DYNAMIC",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "POS",
		},
	}

	qrResult, err := client.MPM().GenerateQR(ctx, qrPayload)
	if err != nil {
		fmt.Printf("Generate QR failed: %v\n", err)
		return
	}

	fmt.Printf("QR Code generated successfully\n")
	fmt.Printf("QR Content: %s\n", qrResult.QRContent)
	fmt.Printf("QR URL: %s\n", qrResult.QRUrl)
	fmt.Printf("Redirect URL: %s\n", qrResult.RedirectUrl)
	fmt.Printf("Valid Until: %s\n", qrResult.ValidUntil)
}

func decodeQRCode(ctx context.Context, client *snap.Client) {
	// Decode QR Code
	decodePayload := &types.MPMDecodeQRPayload{
		PartnerServiceId:   "MPM12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "QR002",
		QRContent:          "00020101021226590012COM.MIDTRANS.WWW01189360055310215MIDTRANS.COM5204581253033605802ID5913MIDTRANS TEST6007JAKARTA610512345626304B8F9",
		Amount:             types.NewAmount(25000, "IDR"),
		MerchantId:         "MERCH001",
		SubMerchantId:      "SUB001",
		ScanTime:           time.Now().Format(time.RFC3339),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().DecodeQR(ctx, decodePayload)
	if err != nil {
		fmt.Printf("Decode QR failed: %v\n", err)
		return
	}

	fmt.Printf("QR Code decoded successfully: %+v\n", result)
}

func applyOTT(ctx context.Context, client *snap.Client) {
	// Apply OTT
	ottPayload := &types.MPMApplyOTTPayload{
		UserResources: []string{"OTT"},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().ApplyOTT(ctx, ottPayload)
	if err != nil {
		fmt.Printf("Apply OTT failed: %v\n", err)
		return
	}

	fmt.Printf("OTT applied successfully: %+v\n", result)
}

func processPayment(ctx context.Context, client *snap.Client) {
	// Process Payment
	paymentPayload := &types.MPMPaymentPayload{
		PartnerServiceId:   "MPM12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "PAY001",
		MerchantId:         "MERCH001",
		SubMerchantId:      "SUB001",
		Amount:             types.NewAmount(25000, "IDR"),
		FeeAmount:          types.NewAmount(0, "IDR"),
		OTP:                "123456",
		VerificationId:     "VERIF001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().Payment(ctx, paymentPayload)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}

	fmt.Printf("Payment processed successfully: %+v\n", result)
}

func queryPayment(ctx context.Context, client *snap.Client) {
	// Query Payment
	queryPayload := &types.MPMQueryPaymentPayload{
		PartnerServiceId:     "MPM12345",
		CustomerNo:           "CUST001",
		OriginalReferenceNo:  "REF001",
		OriginalPartnerRefNo: "PAY001",
		OriginalExternalId:   "EXT001",
		ServiceCode:          "47",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		ExternalStoreId:      "STORE001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().QueryPayment(ctx, queryPayload)
	if err != nil {
		fmt.Printf("Query payment failed: %v\n", err)
		return
	}

	fmt.Printf("Payment query successful: %+v\n", result)
}

func cancelPayment(ctx context.Context, client *snap.Client) {
	// Cancel Payment
	cancelPayload := &types.MPMCancelPaymentPayload{
		PartnerServiceId:     "MPM12345",
		CustomerNo:           "CUST001",
		OriginalPartnerRefNo: "PAY001",
		OriginalReferenceNo:  "REF001",
		OriginalExternalId:   "EXT001",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		ExternalStoreId:      "STORE001",
		Reason:               "Customer requested cancellation",
		Amount:               types.NewAmount(25000, "IDR"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().CancelPayment(ctx, cancelPayload)
	if err != nil {
		fmt.Printf("Cancel payment failed: %v\n", err)
		return
	}

	fmt.Printf("Payment cancelled successfully: %+v\n", result)
}

func refundPayment(ctx context.Context, client *snap.Client) {
	// Refund Payment
	refundPayload := &types.MPMRefundPayload{
		PartnerServiceId:    "MPM12345",
		CustomerNo:          "CUST001",
		OriginalReferenceNo: "REF001",
		PartnerRefundNo:     "REFUND001",
		RefundAmount:        types.NewAmount(25000, "IDR"),
		Reason:              "Customer requested refund",
		RefundType:          "FULL",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().Refund(ctx, refundPayload)
	if err != nil {
		fmt.Printf("Refund payment failed: %v\n", err)
		return
	}

	fmt.Printf("Payment refunded successfully: %+v\n", result)
}

func mpmTransfer(ctx context.Context, client *snap.Client) {
	// MPM Transfer
	transferPayload := &types.MPMTransferPayload{
		PartnerServiceId:      "MPM12345",
		CustomerNo:            "CUST001",
		PartnerReferenceNo:    "TRF001",
		MerchantId:            "MERCH001",
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
		TransactionDate:       time.Now().Format(time.RFC3339),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.MPM().Transfer(ctx, transferPayload)
	if err != nil {
		fmt.Printf("MPM transfer failed: %v\n", err)
		return
	}

	fmt.Printf("MPM transfer successful: %+v\n", result)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}