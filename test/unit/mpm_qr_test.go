package unit

import (
	"context"
	"testing"

	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// TestMPMGenerateQRPayloadValidation tests MPM Generate QR payload validation
func TestMPMGenerateQRPayloadValidation(t *testing.T) {
	tests := []struct {
		name    string
		payload types.MPMGenerateQRPayload
		isValid bool
	}{
		{
			name: "valid generate QR payload",
			payload: types.MPMGenerateQRPayload{
				PartnerServiceId:   "12345",
				CustomerNo:         "CUST001",
				PartnerReferenceNo: "QR001",
				MerchantId:         "MERCHANT001",
				SubMerchantId:      "SUB001",
				ExternalStoreId:    "STORE001",
				Amount:             &types.Amount{Value: "100000", Currency: "IDR"},
				MerchantName:       "Test Merchant",
				MerchantCity:       "Jakarta",
				MerchantPostalCode: "12345",
				TerminalLabel:      "Terminal 1",
				PurposeOfPayment:   "Payment for goods",
				QRType:             "DYNAMIC",
				ValidityPeriod:     "3600",
			},
			isValid: true,
		},
		{
			name: "missing partner service ID",
			payload: types.MPMGenerateQRPayload{
				CustomerNo:         "CUST001",
				PartnerReferenceNo: "QR001",
				MerchantId:         "MERCHANT001",
				Amount:             &types.Amount{Value: "100000", Currency: "IDR"},
				MerchantName:       "Test Merchant",
				QRType:             "DYNAMIC",
			},
			isValid: false,
		},
		{
			name: "missing amount",
			payload: types.MPMGenerateQRPayload{
				PartnerServiceId:   "12345",
				CustomerNo:         "CUST001",
				PartnerReferenceNo: "QR001",
				MerchantId:         "MERCHANT001",
				MerchantName:       "Test Merchant",
				QRType:             "DYNAMIC",
			},
			isValid: false,
		},
		{
			name: "missing QR type",
			payload: types.MPMGenerateQRPayload{
				PartnerServiceId:   "12345",
				CustomerNo:         "CUST001",
				PartnerReferenceNo: "QR001",
				MerchantId:         "MERCHANT001",
				Amount:             &types.Amount{Value: "100000", Currency: "IDR"},
				MerchantName:       "Test Merchant",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation - check if required fields are present
			isValid := validateGenerateQRPayload(tt.payload)
			if isValid != tt.isValid {
				t.Errorf("validateGenerateQRPayload() = %v, want %v", isValid, tt.isValid)
			}
		})
	}
}

// TestMPMNotifyQRPayloadValidation tests MPM Notify QR payload validation
func TestMPMNotifyQRPayloadValidation(t *testing.T) {
	tests := []struct {
		name    string
		payload types.MPMNotifyQRPayload
		isValid bool
	}{
		{
			name: "valid notify QR payload",
			payload: types.MPMNotifyQRPayload{
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
				Amount:                  &types.Amount{Value: "100000", Currency: "IDR"},
				FeeAmount:               &types.Amount{Value: "1000", Currency: "IDR"},
				PaidAmount:              &types.Amount{Value: "101000", Currency: "IDR"},
				CurrencyCode:            "IDR",
				TransactionStatus:       "SUCCESS",
				TransactionStatusDesc:   "Transaction successful",
				LatestTransactionStatus: "COMPLETED",
				PaymentMethod:           "QR_CODE",
				GatewayReferenceNo:      "GW001",
			},
			isValid: true,
		},
		{
			name: "missing partner service ID",
			payload: types.MPMNotifyQRPayload{
				CustomerNo:          "CUST001",
				PartnerReferenceNo:  "NOTIFY001",
				OriginalReferenceNo: "QR001",
				TransactionStatus:   "SUCCESS",
				Amount:              &types.Amount{Value: "100000", Currency: "IDR"},
			},
			isValid: false,
		},
		{
			name: "missing transaction status",
			payload: types.MPMNotifyQRPayload{
				PartnerServiceId:    "12345",
				CustomerNo:          "CUST001",
				PartnerReferenceNo:  "NOTIFY001",
				OriginalReferenceNo: "QR001",
				Amount:              &types.Amount{Value: "100000", Currency: "IDR"},
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation - check if required fields are present
			isValid := validateNotifyQRPayload(tt.payload)
			if isValid != tt.isValid {
				t.Errorf("validateNotifyQRPayload() = %v, want %v", isValid, tt.isValid)
			}
		})
	}
}

// TestMPMQRContextHandling tests context handling for QR operations
func TestMPMQRContextHandling(t *testing.T) {
	ctx := context.Background()

	// Test context with timeout
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	defer cancel()

	if ctxWithTimeout.Err() != nil {
		t.Errorf("Context should not be cancelled initially")
	}

	// Cancel context
	cancel()

	if ctxWithTimeout.Err() == nil {
		t.Errorf("Context should be cancelled after cancel()")
	}
}

// Helper functions for validation

func validateGenerateQRPayload(payload types.MPMGenerateQRPayload) bool {
	if payload.PartnerServiceId == "" {
		return false
	}
	if payload.CustomerNo == "" {
		return false
	}
	if payload.PartnerReferenceNo == "" {
		return false
	}
	if payload.MerchantId == "" {
		return false
	}
	if payload.Amount == nil {
		return false
	}
	if payload.MerchantName == "" {
		return false
	}
	if payload.QRType == "" {
		return false
	}
	return true
}

func validateNotifyQRPayload(payload types.MPMNotifyQRPayload) bool {
	if payload.PartnerServiceId == "" {
		return false
	}
	if payload.CustomerNo == "" {
		return false
	}
	if payload.PartnerReferenceNo == "" {
		return false
	}
	if payload.OriginalReferenceNo == "" {
		return false
	}
	if payload.TransactionStatus == "" {
		return false
	}
	if payload.Amount == nil {
		return false
	}
	return true
}
