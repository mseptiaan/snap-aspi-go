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
	fmt.Println("=== SNAP ASPI SDK Complete Example ===")

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

	// Example 1: Registration
	fmt.Println("\n1. User Registration")
	demonstrateRegistration(ctx, client)

	// Example 2: Balance Inquiry
	fmt.Println("\n2. Balance Inquiry")
	demonstrateBalanceInquiry(ctx, client)

	// Example 3: Transaction History
	fmt.Println("\n3. Transaction History")
	demonstrateTransactionHistory(ctx, client)

	// Example 4: Transfer Credit
	fmt.Println("\n4. Transfer Credit")
	demonstrateTransferCredit(ctx, client)

	// Example 5: Transfer Debit
	fmt.Println("\n5. Transfer Debit")
	demonstrateTransferDebit(ctx, client)

	// Example 6: Virtual Account
	fmt.Println("\n6. Virtual Account")
	demonstrateVirtualAccount(ctx, client)

	// Example 7: MPM
	fmt.Println("\n7. MPM")
	demonstrateMPM(ctx, client)

	fmt.Println("\n=== Complete Example Finished ===")
}

func demonstrateRegistration(ctx context.Context, client *snap.Client) {
	// User registration
	regPayload := &types.RegistrationPayload{
		PartnerServiceId: "REG12345",
		CustomerNo:       "CUST001",
		PhoneNo:          "081234567890",
		Email:            "user@example.com",
		Name:             "John Doe",
		IdNumber:         "1234567890123456",
		IdType:           "KTP",
		Address:          "Jl. Example No. 123, Jakarta",
		DateOfBirth:      "1990-01-01",
		PlaceOfBirth:     "Jakarta",
		Nationality:      "ID",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.Registration().Register(ctx, regPayload)
	if err != nil {
		fmt.Printf("Registration failed: %v\n", err)
	} else {
		fmt.Printf("Registration successful: %+v\n", result)
	}

	// Card registration
	cardPayload := &types.CardRegistrationPayload{
		PartnerServiceId: "REG12345",
		CustomerNo:       "CUST001",
		CardNumber:       "4111111111111111",
		CardType:         "CREDIT",
		ExpiryDate:       "12/25",
		CVV:              "123",
		CardHolderName:   "John Doe",
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	cardResult, err := client.Registration().RegisterCard(ctx, cardPayload)
	if err != nil {
		fmt.Printf("Card registration failed: %v\n", err)
	} else {
		fmt.Printf("Card registration successful: %+v\n", cardResult)
	}
}

func demonstrateBalanceInquiry(ctx context.Context, client *snap.Client) {
	balancePayload := &types.BalanceInquiryPayload{
		PartnerServiceId: "BAL12345",
		CustomerNo:       "CUST001",
		AccountNo:        "1234567890",
		BalanceTypes:     []string{"Cash", "Coins"},
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	result, err := client.BalanceInquiry().BalanceInquiry(ctx, balancePayload)
	if err != nil {
		fmt.Printf("Balance inquiry failed: %v\n", err)
	} else {
		fmt.Printf("Balance inquiry successful: %+v\n", result)
		
		// Display account balances
		if len(result.AccountInfos) > 0 {
			fmt.Println("Account Balances:")
			for i, info := range result.AccountInfos {
				fmt.Printf("  %d. Type: %s, Available: %s, Ledger: %s\n", 
					i+1, 
					info.BalanceType, 
					info.AvailableBalance.String(), 
					info.LedgerBalance.String())
			}
		}
	}
}

func demonstrateTransactionHistory(ctx context.Context, client *snap.Client) {
	// Transaction history list
	historyPayload := &types.TransactionHistoryListPayload{
		PartnerServiceId: "HIST12345",
		CustomerNo:       "CUST001",
		FromDateTime:     time.Now().AddDate(0, -1, 0).Format(time.RFC3339), // 1 month ago
		ToDateTime:       time.Now().Format(time.RFC3339),                   // now
		PageSize:         10,
		PageNumber:       0,
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	historyResult, err := client.TransactionHistory().TransactionHistoryList(ctx, historyPayload)
	if err != nil {
		fmt.Printf("Transaction history list failed: %v\n", err)
	} else {
		fmt.Printf("Transaction history list successful: %d transactions found\n", len(historyResult.DetailData))
		
		// Display transaction details
		if len(historyResult.DetailData) > 0 {
			fmt.Println("Recent Transactions:")
			for i, tx := range historyResult.DetailData {
				if i >= 3 { // Show only first 3 transactions
					break
				}
				fmt.Printf("  %d. Date: %s, Amount: %s, Type: %s, Status: %s\n", 
					i+1, 
					tx.DateTime, 
					tx.Amount.String(), 
					tx.Type, 
					tx.Status)
			}
		}
	}

	// Bank statement
	statementPayload := &types.BankStatementPayload{
		PartnerServiceId: "STMT12345",
		CustomerNo:       "CUST001",
		AccountNo:        "1234567890",
		FromDateTime:     time.Now().AddDate(0, -1, 0).Format(time.RFC3339), // 1 month ago
		ToDateTime:       time.Now().Format(time.RFC3339),                   // now
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	statementResult, err := client.TransactionHistory().BankStatement(ctx, statementPayload)
	if err != nil {
		fmt.Printf("Bank statement failed: %v\n", err)
	} else {
		fmt.Printf("Bank statement successful: %d transactions found\n", len(statementResult.DetailData))
	}
}

func demonstrateTransferCredit(ctx context.Context, client *snap.Client) {
	// Account inquiry
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

	inquiryResult, err := client.TransferCredit().AccountInquiry(ctx, inquiryPayload)
	if err != nil {
		fmt.Printf("Account inquiry failed: %v\n", err)
	} else {
		fmt.Printf("Account inquiry successful: %+v\n", inquiryResult)
	}

	// Trigger transfer
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

	transferResult, err := client.TransferCredit().TriggerTransfer(ctx, transferPayload)
	if err != nil {
		fmt.Printf("Trigger transfer failed: %v\n", err)
	} else {
		fmt.Printf("Trigger transfer successful: %+v\n", transferResult)
	}

	// Customer top up
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

	topupResult, err := client.TransferCredit().CustomerTopUp(ctx, topupPayload)
	if err != nil {
		fmt.Printf("Customer top up failed: %v\n", err)
	} else {
		fmt.Printf("Customer top up successful: %+v\n", topupResult)
	}
}

func demonstrateTransferDebit(ctx context.Context, client *snap.Client) {
	// Direct debit payment
	debitPayload := &types.DirectDebitPaymentPayload{
		PartnerServiceId:   "DEB12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF003",
		BankCardToken:      "TOKEN001",
		MerchantId:         "MERCH001",
		Amount:             types.NewAmount(75000, "IDR"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	debitResult, err := client.TransferDebit().DirectDebitPayment(ctx, debitPayload)
	if err != nil {
		fmt.Printf("Direct debit payment failed: %v\n", err)
	} else {
		fmt.Printf("Direct debit payment successful: %+v\n", debitResult)
	}

	// CPM generate QR
	qrPayload := &types.CPMGenerateQRPayload{
		PartnerServiceId:   "QR12345",
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

	qrResult, err := client.TransferDebit().CPMGenerateQR(ctx, qrPayload)
	if err != nil {
		fmt.Printf("CPM generate QR failed: %v\n", err)
	} else {
		fmt.Printf("CPM generate QR successful: %+v\n", qrResult)
	}

	// Auth payment
	authPayload := &types.AuthPaymentPayload{
		PartnerServiceId:   "AUTH12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF005",
		BankCardToken:      "TOKEN003",
		MerchantId:         "MERCH001",
		Amount:             types.NewAmount(120000, "IDR"),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	authResult, err := client.TransferDebit().AuthPayment(ctx, authPayload)
	if err != nil {
		fmt.Printf("Auth payment failed: %v\n", err)
	} else {
		fmt.Printf("Auth payment successful: %+v\n", authResult)
	}
}

func demonstrateVirtualAccount(ctx context.Context, client *snap.Client) {
	// Create VA
	vaPayload := &types.CreateVAPayload{
		PartnerServiceId:    "VA12345",
		CustomerNo:          "CUST001",
		VirtualAccountNo:    "8808001234567890",
		VirtualAccountName:  "John Doe",
		VirtualAccountEmail: "john@example.com",
		VirtualAccountPhone: "081234567890",
		TrxId:               "TRX001",
		TotalAmount:         types.NewAmount(150000, "IDR"),
		ExpiredDate:         time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // 1 month from now
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	vaResult, err := client.VirtualAccount().CreateVA(ctx, vaPayload)
	if err != nil {
		fmt.Printf("Create VA failed: %v\n", err)
	} else {
		fmt.Printf("Create VA successful: %+v\n", vaResult)
	}

	// Inquiry VA
	inquiryPayload := &types.InquiryVAPayload{
		PartnerServiceId: "VA12345",
		CustomerNo:       "CUST001",
		VirtualAccountNo: "8808001234567890",
		TrxDateInit:      time.Now().Format("2006-01-02"),
	}

	inquiryResult, err := client.VirtualAccount().InquiryVA(ctx, inquiryPayload)
	if err != nil {
		fmt.Printf("Inquiry VA failed: %v\n", err)
	} else {
		fmt.Printf("Inquiry VA successful: %+v\n", inquiryResult)
	}
}

func demonstrateMPM(ctx context.Context, client *snap.Client) {
	// MPM transfer
	mpmPayload := &types.MPMTransferPayload{
		PartnerServiceId:      "MPM12345",
		CustomerNo:            "CUST001",
		PartnerReferenceNo:    "REF006",
		MerchantId:            "MERCH001",
		SubMerchantId:         "SUB001",
		ExternalStoreId:       "STORE001",
		Amount:                types.NewAmount(200000, "IDR"),
		BeneficiaryAccountNo:  "1234567890",
		BeneficiaryName:       "Jane Doe",
		BeneficiaryCustomerNo: "CUST002",
		BeneficiaryEmail:      "jane@example.com",
		Currency:              "IDR",
		CustomerReference:     "CUSTREF002",
		FeeType:               "OUR",
		Remark:                "MPM transfer",
		SourceAccountNo:       "0987654321",
		TransactionDate:       time.Now().Format(time.RFC3339),
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "MOBILE",
		},
	}

	mpmResult, err := client.MPM().Transfer(ctx, mpmPayload)
	if err != nil {
		fmt.Printf("MPM transfer failed: %v\n", err)
	} else {
		fmt.Printf("MPM transfer successful: %+v\n", mpmResult)
	}

	// Generate QR
	qrPayload := &types.MPMGenerateQRPayload{
		PartnerServiceId:   "MPM12345",
		CustomerNo:         "CUST001",
		PartnerReferenceNo: "REF007",
		MerchantId:         "MERCH001",
		SubMerchantId:      "SUB001",
		ExternalStoreId:    "STORE001",
		Amount:             types.NewAmount(50000, "IDR"),
		MerchantName:       "Test Merchant",
		MerchantCity:       "Jakarta",
		MerchantPostalCode: "12345",
		TerminalLabel:      "Terminal 1",
		PurposeOfPayment:   "Payment for goods",
		QRType:             "DYNAMIC",
		ValidityPeriod:     "3600", // 1 hour
		AdditionalInfo: &types.AdditionalInfo{
			DeviceId: "DEVICE001",
			Channel:  "POS",
		},
	}

	qrResult, err := client.MPM().GenerateQR(ctx, qrPayload)
	if err != nil {
		fmt.Printf("Generate QR failed: %v\n", err)
	} else {
		fmt.Printf("Generate QR successful: %s\n", qrResult.QRContent)
		fmt.Printf("Valid Until: %s\n", qrResult.ValidUntil)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}