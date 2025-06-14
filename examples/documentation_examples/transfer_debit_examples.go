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
	fmt.Println("=== SNAP ASPI SDK Transfer Debit Examples ===")

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

	// Example 1: Direct Debit Payment
	fmt.Println("\n1. Direct Debit Payment")
	directDebitPayment(ctx, client)

	// Example 2: Direct Debit Status
	fmt.Println("\n2. Direct Debit Status")
	directDebitStatus(ctx, client)

	// Example 3: Direct Debit Cancel
	fmt.Println("\n3. Direct Debit Cancel")
	directDebitCancel(ctx, client)

	// Example 4: Direct Debit Refund
	fmt.Println("\n4. Direct Debit Refund")
	directDebitRefund(ctx, client)

	// Example 5: CPM Generate QR
	fmt.Println("\n5. CPM Generate QR")
	cpmGenerateQR(ctx, client)

	// Example 6: CPM Payment
	fmt.Println("\n6. CPM Payment")
	cpmPayment(ctx, client)

	// Example 7: Auth Payment
	fmt.Println("\n7. Auth Payment")
	authPayment(ctx, client)

	// Example 8: Auth Capture
	fmt.Println("\n8. Auth Capture")
	authCapture(ctx, client)

	// Example 9: Auth Void
	fmt.Println("\n9. Auth Void")
	authVoid(ctx, client)

	// Example 10: Direct Debit BI-FAST
	fmt.Println("\n10. Direct Debit BI-FAST")
	directDebitBIFAST(ctx, client)

	fmt.Println("\n=== Transfer Debit Examples Finished ===")
}

func directDebitPayment(ctx context.Context, client *snap.Client) {
	// Direct Debit Payment
	debitPayload := &types.DirectDebitPaymentPayload{
		PartnerServiceId:   "DEB12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF003",
		BankCardToken:      "TOKEN001",
		ChargeToken:        "CHARGE001",
		OTP:                "123456",
		OTPTrxCode:         "01",
		MerchantId:         "MERCH001",
		TerminalId:         "TERM001",
		JourneyId:          "JOURNEY001",
		SubMerchantId:      "SUB001",
		Amount:             types.NewAmount(75000, "IDR"),
		URLParams: []types.URLParam{
			{
				URL:        "https://example.com/return",
				Type:       "PAY_RETURN",
				IsDeeplink: "N",
			},
			{
				URL:        "https://example.com/notify",
				Type:       "PAY_NOTIFY",
				IsDeeplink: "N",
			},
		},
		ExternalStoreId:    "STORE001",
		ValidUpTo:          time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		PointOfInitiation:  "Mobile App",
		FeeType:            "OUR",
		DisabledPayMethods: "CREDIT_CARD",
		PayOptionDetails: []types.PayOptionDetail{
			{
				PayMethod:     "DEBIT_CARD",
				PayOption:     "DEBIT_CARD_VISA",
				TransAmount:   types.NewAmount(75000, "IDR"),
				FeeAmount:     types.NewAmount(0, "IDR"),
				CardToken:     "TOKEN001",
				MerchantToken: "MTOKEN001",
				AdditionalInfo: &types.AdditionalInfo{
					DeviceId: "DEVICE001",
					Channel:  "MOBILE",
				},
			},
		},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().DirectDebitPayment(ctx, debitPayload)
	if err != nil {
		fmt.Printf("Direct debit payment failed: %v\n", err)
		return
	}

	fmt.Printf("Direct debit payment successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
	fmt.Printf("Approval Code: %s\n", result["approvalCode"])
}

func directDebitStatus(ctx context.Context, client *snap.Client) {
	// Direct Debit Status
	statusPayload := &types.DirectDebitStatusPayload{
		PartnerServiceId:     "DEB12345",
		CustomerNo:           "CUST001",
		OriginalPartnerRefNo: "REF003",
		OriginalReferenceNo:  "REF003",
		OriginalExternalId:   "EXT001",
		ServiceCode:          "54",
		TransactionDate:      time.Now().Format(time.RFC3339),
		Amount:               types.NewAmount(75000, "IDR"),
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		ExternalStoreId:      "STORE001",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().DirectDebitStatus(ctx, statusPayload)
	if err != nil {
		fmt.Printf("Direct debit status check failed: %v\n", err)
		return
	}

	fmt.Printf("Direct debit status check successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Latest Transaction Status: %s\n", result["latestTransactionStatus"])
	fmt.Printf("Transaction Status Description: %s\n", result["transactionStatusDesc"])
}

func directDebitCancel(ctx context.Context, client *snap.Client) {
	// Direct Debit Cancel
	cancelPayload := &types.DirectDebitCancelPayload{
		PartnerServiceId:     "DEB12345",
		CustomerNo:           "CUST001",
		OriginalPartnerRefNo: "REF003",
		OriginalReferenceNo:  "REF003",
		ApprovalCode:         "APPROVAL001",
		OriginalExternalId:   "EXT001",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		Reason:               "Customer requested cancellation",
		ExternalStoreId:      "STORE001",
		Amount:               types.NewAmount(75000, "IDR"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().DirectDebitCancel(ctx, cancelPayload)
	if err != nil {
		fmt.Printf("Direct debit cancel failed: %v\n", err)
		return
	}

	fmt.Printf("Direct debit cancel successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Cancel Time: %s\n", result["cancelTime"])
}

func directDebitRefund(ctx context.Context, client *snap.Client) {
	// Direct Debit Refund
	refundPayload := &types.DirectDebitRefundPayload{
		PartnerServiceId:     "DEB12345",
		CustomerNo:           "CUST001",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		OriginalPartnerRefNo: "REF003",
		OriginalReferenceNo:  "REF003",
		OriginalExternalId:   "EXT001",
		PartnerRefundNo:      "REFUND001",
		RefundAmount:         types.NewAmount(75000, "IDR"),
		ExternalStoreId:      "STORE001",
		Reason:               "Customer requested refund",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().DirectDebitRefund(ctx, refundPayload)
	if err != nil {
		fmt.Printf("Direct debit refund failed: %v\n", err)
		return
	}

	fmt.Printf("Direct debit refund successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Refund No: %s\n", result["refundNo"])
	fmt.Printf("Refund Time: %s\n", result["refundTime"])
}

func cpmGenerateQR(ctx context.Context, client *snap.Client) {
	// CPM Generate QR
	qrPayload := &types.CPMGenerateQRPayload{
		PartnerServiceId:   "CPM12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF004",
		UserAccessToken:    "TOKEN002",
		MerchantId:         "MERCH001",
		SubMerchantId:      "SUB001",
		PartnerTrxDate:     time.Now().Format(time.RFC3339),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().CPMGenerateQR(ctx, qrPayload)
	if err != nil {
		fmt.Printf("CPM generate QR failed: %v\n", err)
		return
	}

	fmt.Printf("CPM generate QR successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("QR Content: %s\n", result["qrContent"])
	fmt.Printf("QR URL: %s\n", result["qrUrl"])
}

func cpmPayment(ctx context.Context, client *snap.Client) {
	// CPM Payment
	paymentPayload := &types.CPMPaymentPayload{
		PartnerServiceId:  "CPM12345",
		CustomerNo:        "CUST001",
		PartnerReferenceNo: "REF005",
		QRContent:         "00020101...",
		Amount:            types.NewAmount(50000, "IDR"),
		FeeAmount:         types.NewAmount(0, "IDR"),
		MerchantId:        "MERCH001",
		SubMerchantId:     "SUB001",
		Title:             "Payment for goods",
		ExpiryTime:        time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		Items: &types.ItemInfo{
			ProductId:   "PROD001",
			ProductName: "Product 1",
			Qty:         "1",
			Desc:        "Product description",
		},
		ExternalStoreId:  "STORE001",
		MerchantName:     "Test Merchant",
		MerchantLocation: "Jakarta",
		AcquirerName:     "BCA",
		TerminalId:       "TERM001",
		ScannerInfo: &types.ScannerInfo{
			DeviceId:      "SCANNER001",
			DeviceVersion: "1.0",
			DeviceModel:   "Scanner Model X",
			DeviceIP:      "192.168.1.1",
		},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().CPMPayment(ctx, paymentPayload)
	if err != nil {
		fmt.Printf("CPM payment failed: %v\n", err)
		return
	}

	fmt.Printf("CPM payment successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
}

func authPayment(ctx context.Context, client *snap.Client) {
	// Auth Payment
	authPayload := &types.AuthPaymentPayload{
		PartnerServiceId:   "AUTH12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF006",
		BankCardToken:      "TOKEN003",
		ChargeToken:        "CHARGE002",
		OTP:                "123456",
		MerchantId:         "MERCH001",
		TerminalId:         "TERM001",
		JourneyId:          "JOURNEY001",
		SubMerchantId:      "SUB001",
		Amount:             types.NewAmount(120000, "IDR"),
		URLParams: []types.URLParam{
			{
				URL:        "https://example.com/return",
				Type:       "PAY_RETURN",
				IsDeeplink: "N",
			},
		},
		ExternalStoreId:    "STORE001",
		ValidUpTo:          time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		PointOfInitiation:  "Mobile App",
		FeeType:            "OUR",
		DisabledPayMethods: "CREDIT_CARD",
		PayOptionDetails: []types.PayOptionDetail{
			{
				PayMethod:     "DEBIT_CARD",
				PayOption:     "DEBIT_CARD_VISA",
				TransAmount:   types.NewAmount(120000, "IDR"),
				FeeAmount:     types.NewAmount(0, "IDR"),
				CardToken:     "TOKEN003",
				MerchantToken: "MTOKEN002",
				AdditionalInfo: &types.AdditionalInfo{
					DeviceId: "DEVICE001",
					Channel:  "MOBILE",
				},
			},
		},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().AuthPayment(ctx, authPayload)
	if err != nil {
		fmt.Printf("Auth payment failed: %v\n", err)
		return
	}

	fmt.Printf("Auth payment successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Reference No: %s\n", result["referenceNo"])
}

func authCapture(ctx context.Context, client *snap.Client) {
	// Auth Capture
	capturePayload := &types.AuthCapturePayload{
		PartnerServiceId:     "AUTH12345",
		CustomerNo:           "CUST001",
		OriginalReferenceNo:  "REF006",
		OriginalPartnerRefNo: "REF006",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		PartnerCaptureNo:     "CAPTURE001",
		CaptureAmount:        types.NewAmount(120000, "IDR"),
		Title:                "Capture payment",
		LastCapture:          "TRUE",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().AuthCapture(ctx, capturePayload)
	if err != nil {
		fmt.Printf("Auth capture failed: %v\n", err)
		return
	}

	fmt.Printf("Auth capture successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Capture No: %s\n", result["captureNo"])
	fmt.Printf("Capture Time: %s\n", result["captureTime"])
}

func authVoid(ctx context.Context, client *snap.Client) {
	// Auth Void
	voidPayload := &types.AuthVoidPayload{
		PartnerServiceId:     "AUTH12345",
		CustomerNo:           "CUST001",
		OriginalReferenceNo:  "REF006",
		OriginalPartnerRefNo: "REF006",
		MerchantId:           "MERCH001",
		SubMerchantId:        "SUB001",
		VoidAmount:           types.NewAmount(120000, "IDR"),
		PartnerVoidNo:        "VOID001",
		VoidRemainingAmount:  "TRUE",
		Reason:               "Customer requested void",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().AuthVoid(ctx, voidPayload)
	if err != nil {
		fmt.Printf("Auth void failed: %v\n", err)
		return
	}

	fmt.Printf("Auth void successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("Void No: %s\n", result["voidNo"])
	fmt.Printf("Void Time: %s\n", result["voidTime"])
}

func directDebitBIFAST(ctx context.Context, client *snap.Client) {
	// Direct Debit BI-FAST
	bifastPayload := &types.DirectDebitBIFASTPayload{
		PartnerServiceId:  "BIFAST12345",
		CustomerNo:        "CUST001",
		PartnerReferenceNo: "REF007",
		BankCode:          "014",
		SourceAccountNo:   "1234567890",
		SourceAccountName: "John Doe",
		MaxAmount:         types.NewAmount(1000000, "IDR"),
		BillerId:          "BILLER001",
		BillerName:        "Electricity Company",
		CustomerId:        "CUST001",
		ExpiredDatetime:   time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.TransferDebit().DirectDebitBIFAST(ctx, bifastPayload)
	if err != nil {
		fmt.Printf("Direct debit BI-FAST failed: %v\n", err)
		return
	}

	fmt.Printf("Direct debit BI-FAST successful\n")
	fmt.Printf("Response Code: %s\n", result["responseCode"])
	fmt.Printf("Response Message: %s\n", result["responseMessage"])
	fmt.Printf("E-Mandate Reference ID: %s\n", result["eMandateReffId"])
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}