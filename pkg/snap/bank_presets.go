package snap

// BankPresets provides predefined configurations for different banks
type BankPresets struct{}

// GetBankConfig returns predefined configuration for specific banks
func (bp *BankPresets) GetBankConfig(bankCode string) *CustomEndpoints {
	switch bankCode {
	case "BCA":
		return bp.getBCAEndpoints()
	case "BNI":
		return bp.getBNIEndpoints()
	case "BRI":
		return bp.getBRIEndpoints()
	case "MANDIRI":
		return bp.getMandiriEndpoints()
	case "CIMB":
		return bp.getCIMBEndpoints()
	case "PERMATA":
		return bp.getPermataEndpoints()
	default:
		return nil // Use default endpoints
	}
}

// Bank-specific endpoint configurations

func (bp *BankPresets) getBCAEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bca/transfer-va/create-va",
			UpdateVA:     "/api/v1.0/bca/transfer-va/update-va",
			DeleteVA:     "/api/v1.0/bca/transfer-va/delete-va",
			InquiryVA:    "/api/v1.0/bca/transfer-va/inquiry-va",
			Inquiry:      "/api/v1.0/bca/transfer-va/inquiry",
			Payment:      "/api/v1.0/bca/transfer-va/payment",
			Status:       "/api/v1.0/bca/transfer-va/status",
			Report:       "/api/v1.0/bca/transfer-va/report",
			UpdateStatus: "/api/v1.0/bca/transfer-va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bca/transfer-kredit/mpm",
			Inquiry:        "/api/v1.0/bca/transfer-kredit/mpm/inquiry",
			Status:         "/api/v1.0/bca/transfer-kredit/mpm/status",
			Refund:         "/api/v1.0/bca/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/bca/transfer-kredit/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/bca/transfer-kredit/mpm/account-inquiry",
			History:        "/api/v1.0/bca/transfer-kredit/mpm/history",
			GenerateQR:     "/api/v1.0/bca/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/bca/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/bca/qr/apply-ott",
			Payment:        "/api/v1.0/bca/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/bca/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/bca/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/bca/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bca/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bca/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bca/transaction-history-detail",
			BankStatement:            "/api/v1.0/bca/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bca/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bca/trigger-transfer",
			TransferStatus:  "/api/v1.0/bca/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bca/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bca/debit/status",
			DirectDebitCancel:  "/api/v1.0/bca/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bca/debit/refund",
			CPMGenerateQR:      "/api/v1.0/bca/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/bca/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/bca/auth/payment",
			AuthCapture:        "/api/v1.0/bca/auth/capture",
			AuthVoid:           "/api/v1.0/bca/auth/void",
		},
	}
}

func (bp *BankPresets) getBNIEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bni/virtual-account/create",
			UpdateVA:     "/api/v1.0/bni/virtual-account/update",
			DeleteVA:     "/api/v1.0/bni/virtual-account/delete",
			InquiryVA:    "/api/v1.0/bni/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/bni/virtual-account/inquiry",
			Payment:      "/api/v1.0/bni/virtual-account/payment",
			Status:       "/api/v1.0/bni/virtual-account/status",
			Report:       "/api/v1.0/bni/virtual-account/report",
			UpdateStatus: "/api/v1.0/bni/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bni/mpm/transfer",
			Inquiry:        "/api/v1.0/bni/mpm/inquiry",
			Status:         "/api/v1.0/bni/mpm/status",
			Refund:         "/api/v1.0/bni/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/bni/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/bni/mpm/account-inquiry",
			History:        "/api/v1.0/bni/mpm/history",
			GenerateQR:     "/api/v1.0/bni/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/bni/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/bni/qr/apply-ott",
			Payment:        "/api/v1.0/bni/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/bni/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/bni/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/bni/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bni/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bni/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bni/transaction-history-detail",
			BankStatement:            "/api/v1.0/bni/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bni/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bni/trigger-transfer",
			TransferStatus:  "/api/v1.0/bni/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bni/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bni/debit/status",
			DirectDebitCancel:  "/api/v1.0/bni/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bni/debit/refund",
			CPMGenerateQR:      "/api/v1.0/bni/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/bni/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/bni/auth/payment",
			AuthCapture:        "/api/v1.0/bni/auth/capture",
			AuthVoid:           "/api/v1.0/bni/auth/void",
		},
	}
}

func (bp *BankPresets) getBRIEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/bri/va/create",
			UpdateVA:     "/api/v1.0/bri/va/update",
			DeleteVA:     "/api/v1.0/bri/va/delete",
			InquiryVA:    "/api/v1.0/bri/va/inquiry-va",
			Inquiry:      "/api/v1.0/bri/va/inquiry",
			Payment:      "/api/v1.0/bri/va/payment",
			Status:       "/api/v1.0/bri/va/status",
			Report:       "/api/v1.0/bri/va/report",
			UpdateStatus: "/api/v1.0/bri/va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/bri/mpm-transfer/transfer",
			Inquiry:        "/api/v1.0/bri/mpm-transfer/inquiry",
			Status:         "/api/v1.0/bri/mpm-transfer/status",
			Refund:         "/api/v1.0/bri/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/bri/mpm-transfer/balance",
			AccountInquiry: "/api/v1.0/bri/mpm-transfer/account",
			History:        "/api/v1.0/bri/mpm-transfer/history",
			GenerateQR:     "/api/v1.0/bri/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/bri/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/bri/qr/apply-ott",
			Payment:        "/api/v1.0/bri/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/bri/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/bri/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/bri/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/bri/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/bri/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/bri/transaction-history-detail",
			BankStatement:            "/api/v1.0/bri/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/bri/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bri/trigger-transfer",
			TransferStatus:  "/api/v1.0/bri/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/bri/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/bri/debit/status",
			DirectDebitCancel:  "/api/v1.0/bri/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/bri/debit/refund",
			CPMGenerateQR:      "/api/v1.0/bri/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/bri/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/bri/auth/payment",
			AuthCapture:        "/api/v1.0/bri/auth/capture",
			AuthVoid:           "/api/v1.0/bri/auth/void",
		},
	}
}

func (bp *BankPresets) getMandiriEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/mandiri/virtual-account/create-va",
			UpdateVA:     "/api/v1.0/mandiri/virtual-account/update-va",
			DeleteVA:     "/api/v1.0/mandiri/virtual-account/delete-va",
			InquiryVA:    "/api/v1.0/mandiri/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/mandiri/virtual-account/inquiry",
			Payment:      "/api/v1.0/mandiri/virtual-account/payment",
			Status:       "/api/v1.0/mandiri/virtual-account/status",
			Report:       "/api/v1.0/mandiri/virtual-account/report",
			UpdateStatus: "/api/v1.0/mandiri/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/mandiri/transfer-kredit/mpm",
			Inquiry:        "/api/v1.0/mandiri/transfer-kredit/mpm/inquiry",
			Status:         "/api/v1.0/mandiri/transfer-kredit/mpm/status",
			Refund:         "/api/v1.0/mandiri/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/mandiri/transfer-kredit/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/mandiri/transfer-kredit/mpm/account-inquiry",
			History:        "/api/v1.0/mandiri/transfer-kredit/mpm/history",
			GenerateQR:     "/api/v1.0/mandiri/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/mandiri/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/mandiri/qr/apply-ott",
			Payment:        "/api/v1.0/mandiri/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/mandiri/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/mandiri/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/mandiri/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/mandiri/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/mandiri/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/mandiri/transaction-history-detail",
			BankStatement:            "/api/v1.0/mandiri/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/mandiri/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/mandiri/trigger-transfer",
			TransferStatus:  "/api/v1.0/mandiri/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/mandiri/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/mandiri/debit/status",
			DirectDebitCancel:  "/api/v1.0/mandiri/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/mandiri/debit/refund",
			CPMGenerateQR:      "/api/v1.0/mandiri/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/mandiri/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/mandiri/auth/payment",
			AuthCapture:        "/api/v1.0/mandiri/auth/capture",
			AuthVoid:           "/api/v1.0/mandiri/auth/void",
		},
	}
}

func (bp *BankPresets) getCIMBEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/cimb/va/create",
			UpdateVA:     "/api/v1.0/cimb/va/update",
			DeleteVA:     "/api/v1.0/cimb/va/delete",
			InquiryVA:    "/api/v1.0/cimb/va/inquiry-va",
			Inquiry:      "/api/v1.0/cimb/va/inquiry",
			Payment:      "/api/v1.0/cimb/va/payment",
			Status:       "/api/v1.0/cimb/va/status",
			Report:       "/api/v1.0/cimb/va/report",
			UpdateStatus: "/api/v1.0/cimb/va/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/cimb/mpm/transfer",
			Inquiry:        "/api/v1.0/cimb/mpm/inquiry",
			Status:         "/api/v1.0/cimb/mpm/status",
			Refund:         "/api/v1.0/cimb/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/cimb/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/cimb/mpm/account-inquiry",
			History:        "/api/v1.0/cimb/mpm/history",
			GenerateQR:     "/api/v1.0/cimb/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/cimb/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/cimb/qr/apply-ott",
			Payment:        "/api/v1.0/cimb/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/cimb/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/cimb/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/cimb/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/cimb/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/cimb/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/cimb/transaction-history-detail",
			BankStatement:            "/api/v1.0/cimb/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/cimb/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/cimb/trigger-transfer",
			TransferStatus:  "/api/v1.0/cimb/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/cimb/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/cimb/debit/status",
			DirectDebitCancel:  "/api/v1.0/cimb/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/cimb/debit/refund",
			CPMGenerateQR:      "/api/v1.0/cimb/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/cimb/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/cimb/auth/payment",
			AuthCapture:        "/api/v1.0/cimb/auth/capture",
			AuthVoid:           "/api/v1.0/cimb/auth/void",
		},
	}
}

func (bp *BankPresets) getPermataEndpoints() *CustomEndpoints {
	return &CustomEndpoints{
		B2BToken:   "/api/v1.0/access-token/b2b",
		B2B2CToken: "/api/v1.0/access-token/b2b2c",
		VirtualAccount: &VirtualAccountEndpoints{
			CreateVA:     "/api/v1.0/permata/virtual-account/create",
			UpdateVA:     "/api/v1.0/permata/virtual-account/update",
			DeleteVA:     "/api/v1.0/permata/virtual-account/delete",
			InquiryVA:    "/api/v1.0/permata/virtual-account/inquiry-va",
			Inquiry:      "/api/v1.0/permata/virtual-account/inquiry",
			Payment:      "/api/v1.0/permata/virtual-account/payment",
			Status:       "/api/v1.0/permata/virtual-account/status",
			Report:       "/api/v1.0/permata/virtual-account/report",
			UpdateStatus: "/api/v1.0/permata/virtual-account/update-status",
		},
		MPM: &MPMEndpoints{
			Transfer:       "/api/v1.0/permata/mpm/transfer",
			Inquiry:        "/api/v1.0/permata/mpm/inquiry",
			Status:         "/api/v1.0/permata/mpm/status",
			Refund:         "/api/v1.0/permata/qr/qr-mpm-refund",
			BalanceInquiry: "/api/v1.0/permata/mpm/balance-inquiry",
			AccountInquiry: "/api/v1.0/permata/mpm/account-inquiry",
			History:        "/api/v1.0/permata/mpm/history",
			GenerateQR:     "/api/v1.0/permata/qr/qr-mpm-generate",
			DecodeQR:       "/api/v1.0/permata/qr/qr-mpm-decode",
			ApplyOTT:       "/api/v1.0/permata/qr/apply-ott",
			Payment:        "/api/v1.0/permata/qr/qr-mpm-payment",
			QueryPayment:   "/api/v1.0/permata/qr/qr-mpm-query",
			CancelPayment:  "/api/v1.0/permata/qr/qr-mpm-cancel",
			NotifyQR:       "/api/v1.0/permata/qr/qr-mpm-notify",
		},
		BalanceInquiry: &BalanceInquiryEndpoints{
			BalanceInquiry: "/api/v1.0/permata/balance-inquiry",
		},
		TransactionHistory: &TransactionHistoryEndpoints{
			TransactionHistoryList:   "/api/v1.0/permata/transaction-history-list",
			TransactionHistoryDetail: "/api/v1.0/permata/transaction-history-detail",
			BankStatement:            "/api/v1.0/permata/bank-statement",
		},
		TransferCredit: &TransferCreditEndpoints{
			AccountInquiry:  "/api/v1.0/permata/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/permata/trigger-transfer",
			TransferStatus:  "/api/v1.0/permata/transfer/status",
		},
		TransferDebit: &TransferDebitEndpoints{
			DirectDebitPayment: "/api/v1.0/permata/debit/payment-host-to-host",
			DirectDebitStatus:  "/api/v1.0/permata/debit/status",
			DirectDebitCancel:  "/api/v1.0/permata/debit/cancel",
			DirectDebitRefund:  "/api/v1.0/permata/debit/refund",
			CPMGenerateQR:      "/api/v1.0/permata/qr/qr-cpm-generate",
			CPMPayment:         "/api/v1.0/permata/qr/qr-cpm-payment",
			AuthPayment:        "/api/v1.0/permata/auth/payment",
			AuthCapture:        "/api/v1.0/permata/auth/capture",
			AuthVoid:           "/api/v1.0/permata/auth/void",
		},
	}
}