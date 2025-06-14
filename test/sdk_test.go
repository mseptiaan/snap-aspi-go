package test

import (
	"context"
	"testing"

	"github.com/mseptiaan/snap-aspi-go/pkg/snap"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSDKInitialization(t *testing.T) {
	tests := []struct {
		name    string
		config  snap.Config
		wantErr bool
	}{
		{
			name: "valid_config",
			config: snap.Config{
				BaseURL:        "https://sandbox.aspi-indonesia.or.id",
				ClientID:       "test-client-id",
				ClientSecret:   "test-client-secret",
				PrivateKeyPath: "../keys/private_key.pem",
				PublicKeyPath:  "../keys/public_key.pem",
				Environment:    "sandbox",
			},
			wantErr: false,
		},
		{
			name: "missing_base_url",
			config: snap.Config{
				ClientID:       "test-client-id",
				ClientSecret:   "test-client-secret",
				PrivateKeyPath: "../keys/private_key.pem",
			},
			wantErr: true,
		},
		{
			name: "missing_client_id",
			config: snap.Config{
				BaseURL:        "https://sandbox.aspi-indonesia.or.id",
				ClientSecret:   "test-client-secret",
				PrivateKeyPath: "../keys/private_key.pem",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := snap.NewClient(tt.config)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
				
				// Test service access
				assert.NotNil(t, client.VirtualAccount())
				assert.NotNil(t, client.MPM())
				assert.NotNil(t, client.Auth())
			}
		})
	}
}

func TestBankPresets(t *testing.T) {
	bankPresets := &snap.BankPresets{}
	
	supportedBanks := []string{"BCA", "BNI", "BRI", "MANDIRI", "CIMB", "PERMATA"}
	
	for _, bankCode := range supportedBanks {
		t.Run(bankCode, func(t *testing.T) {
			config := bankPresets.GetBankConfig(bankCode)
			assert.NotNil(t, config, "Bank config should exist for %s", bankCode)
			
			// Check that bank-specific endpoints are set
			if config.VirtualAccount != nil {
				assert.NotEmpty(t, config.VirtualAccount.CreateVA, "CreateVA endpoint should be set for %s", bankCode)
			}
			if config.MPM != nil {
				assert.NotEmpty(t, config.MPM.Transfer, "Transfer endpoint should be set for %s", bankCode)
			}
		})
	}
}

func TestVirtualAccountPayloadValidation(t *testing.T) {
	tests := []struct {
		name    string
		payload types.CreateVAPayload
		isValid bool
	}{
		{
			name: "valid_payload",
			payload: types.CreateVAPayload{
				PartnerServiceId:   "12345",
				CustomerNo:         "CUST001",
				VirtualAccountNo:   "8808001234567890",
				VirtualAccountName: "Test Customer",
				TrxId:              "TRX-001",
				TotalAmount:        types.NewAmount(100000, "IDR"),
				ExpiredDate:        "2024-12-31T23:59:59+07:00",
			},
			isValid: true,
		},
		{
			name: "missing_partner_service_id",
			payload: types.CreateVAPayload{
				CustomerNo:         "CUST001",
				VirtualAccountNo:   "8808001234567890",
				VirtualAccountName: "Test Customer",
				TrxId:              "TRX-001",
				TotalAmount:        types.NewAmount(100000, "IDR"),
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation - check required fields
			hasPartnerServiceId := tt.payload.PartnerServiceId != ""
			hasCustomerNo := tt.payload.CustomerNo != ""
			hasVirtualAccountNo := tt.payload.VirtualAccountNo != ""
			hasTotalAmount := tt.payload.TotalAmount != nil

			isValid := hasPartnerServiceId && hasCustomerNo && hasVirtualAccountNo && hasTotalAmount

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestMPMPayloadValidation(t *testing.T) {
	tests := []struct {
		name    string
		payload types.MPMTransferPayload
		isValid bool
	}{
		{
			name: "valid_mpm_payload",
			payload: types.MPMTransferPayload{
				PartnerServiceId:     "12345",
				CustomerNo:           "CUST001",
				PartnerReferenceNo:   "REF001",
				Amount:               types.NewAmount(50000, "IDR"),
				BeneficiaryAccountNo: "1234567890",
				BeneficiaryName:      "Jane Doe",
				Currency:             "IDR",
			},
			isValid: true,
		},
		{
			name: "missing_amount",
			payload: types.MPMTransferPayload{
				PartnerServiceId:     "12345",
				CustomerNo:           "CUST001",
				PartnerReferenceNo:   "REF001",
				BeneficiaryAccountNo: "1234567890",
				BeneficiaryName:      "Jane Doe",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation - check required fields
			hasPartnerServiceId := tt.payload.PartnerServiceId != ""
			hasCustomerNo := tt.payload.CustomerNo != ""
			hasAmount := tt.payload.Amount != nil
			hasBeneficiaryAccountNo := tt.payload.BeneficiaryAccountNo != ""

			isValid := hasPartnerServiceId && hasCustomerNo && hasAmount && hasBeneficiaryAccountNo

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

func TestAmountType(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		currency string
		expected string
	}{
		{
			name:     "idr_amount",
			amount:   100000,
			currency: "IDR",
			expected: "100000.00 IDR",
		},
		{
			name:     "usd_amount",
			amount:   100.50,
			currency: "USD",
			expected: "100.50 USD",
		},
		{
			name:     "default_currency",
			amount:   50000,
			currency: "",
			expected: "50000.00 IDR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount := types.NewAmount(tt.amount, tt.currency)
			
			assert.NotNil(t, amount)
			assert.Equal(t, tt.expected, amount.String())
			
			// Test conversion back to float64
			value, err := amount.Float64()
			assert.NoError(t, err)
			assert.Equal(t, tt.amount, value)
		})
	}
}

func TestContextHandling(t *testing.T) {
	ctx := context.Background()
	assert.NotNil(t, ctx)

	// Test context with timeout
	ctxWithTimeout, cancel := context.WithCancel(ctx)
	defer cancel()

	assert.NotNil(t, ctxWithTimeout)
	assert.NoError(t, ctxWithTimeout.Err())

	// Cancel context
	cancel()
	assert.Error(t, ctxWithTimeout.Err())
}

// Benchmark tests for performance
func BenchmarkSDKInitialization(b *testing.B) {
	config := snap.Config{
		BaseURL:        "https://sandbox.aspi-indonesia.or.id",
		ClientID:       "test-client-id",
		ClientSecret:   "test-client-secret",
		PrivateKeyPath: "../keys/private_key.pem",
		PublicKeyPath:  "../keys/public_key.pem",
		Environment:    "sandbox",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = snap.NewClient(config)
	}
}

func BenchmarkAmountCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = types.NewAmount(100000, "IDR")
	}
}