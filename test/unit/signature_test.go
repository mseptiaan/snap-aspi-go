package unit

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/signature"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAsymmetricSigner(t *testing.T) {
	// Generate test RSA keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Create temporary key files
	tempDir := t.TempDir()
	privateKeyPath := tempDir + "/private_key.pem"
	publicKeyPath := tempDir + "/public_key.pem"

	// Write private key
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	require.NoError(t, err)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = os.WriteFile(privateKeyPath, privateKeyPEM, 0600)
	require.NoError(t, err)

	// Write public key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	err = os.WriteFile(publicKeyPath, publicKeyPEM, 0644)
	require.NoError(t, err)

	tests := []struct {
		name    string
		payload signature.AsymmetricPayload
		wantErr bool
	}{
		{
			name: "successful_b2b_signature",
			payload: signature.AsymmetricPayload{
				ClientID:  "test-client-id",
				Timestamp: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "empty_client_id",
			payload: signature.AsymmetricPayload{
				ClientID:  "",
				Timestamp: time.Now(),
			},
			wantErr: false, // This might still work, depends on validation
		},
		{
			name: "future_timestamp",
			payload: signature.AsymmetricPayload{
				ClientID:  "test-client-id",
				Timestamp: time.Now().Add(24 * time.Hour),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create asymmetric signer
			signer, err := signature.NewAsymmetricSigner(privateKeyPEM, publicKeyPEM)
			require.NoError(t, err)

			// Test signing
			signatureValue, err := signer.Sign(tt.payload)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, signatureValue)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, signatureValue)

				// Verify signature format (should be base64 encoded)
				assert.Greater(t, len(signatureValue), 100) // RSA 2048 signature should be ~344 chars base64 encoded

				// Test verification (returns error, not bool)
				err = signer.Verify(tt.payload, signatureValue)
				assert.NoError(t, err)

				// Test with modified payload (should fail verification)
				modifiedPayload := tt.payload
				modifiedPayload.ClientID = "different-client-id"
				err = signer.Verify(modifiedPayload, signatureValue)
				assert.Error(t, err)
			}
		})
	}
}

func TestSymmetricSigner(t *testing.T) {
	tests := []struct {
		name         string
		clientSecret string
		payload      signature.SymmetricPayload
		wantErr      bool
	}{
		{
			name:         "successful_symmetric_signature",
			clientSecret: "test-client-secret-12345",
			payload: signature.SymmetricPayload{
				HTTPMethod:  "POST",
				EndpointURL: "/api/v1.0/transfer-va/create-va",
				AccessToken: "sample-access-token",
				RequestBody: map[string]interface{}{
					"partnerServiceId":   "test-service",
					"customerNo":         "123456789",
					"virtualAccountNo":   "8808001234567890",
					"virtualAccountName": "Test Account",
					"trxId":              "TRX-001",
					"totalAmount": map[string]interface{}{
						"value":    "100000.00",
						"currency": "IDR",
					},
				},
				Timestamp: time.Now(),
			},
			wantErr: false,
		},
		{
			name:         "empty_access_token",
			clientSecret: "test-client-secret-12345",
			payload: signature.SymmetricPayload{
				HTTPMethod:  "POST",
				EndpointURL: "/api/v1.0/transfer-va/create-va",
				AccessToken: "",
				RequestBody: map[string]interface{}{},
				Timestamp:   time.Now(),
			},
			wantErr: false, // Empty access token is actually allowed by implementation
		},
		{
			name:         "empty_client_secret",
			clientSecret: "",
			payload: signature.SymmetricPayload{
				HTTPMethod:  "POST",
				EndpointURL: "/api/v1.0/transfer-va/create-va",
				AccessToken: "sample-access-token",
				RequestBody: map[string]interface{}{},
				Timestamp:   time.Now(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create symmetric signer
			signer := signature.NewSymmetricSigner(tt.clientSecret)

			// Test signing
			signatureValue, err := signer.Sign(tt.payload)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, signatureValue)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, signatureValue)

				// Verify signature format (should be base64 encoded HMAC-SHA512)
				assert.Greater(t, len(signatureValue), 80) // Base64 SHA512 should be ~88 chars

				// Test verification (returns error, not bool)
				err = signer.Verify(tt.payload, signatureValue)
				assert.NoError(t, err)

				// Test with modified payload (should fail verification)
				modifiedPayload := tt.payload
				modifiedPayload.AccessToken = "different-access-token"
				err = signer.Verify(modifiedPayload, signatureValue)
				assert.Error(t, err)
			}
		})
	}
}

func TestSignatureConcurrency(t *testing.T) {
	// Test concurrent signature generation
	signer := signature.NewSymmetricSigner("test-secret")

	payload := signature.SymmetricPayload{
		HTTPMethod:  "POST",
		EndpointURL: "/api/v1.0/test",
		AccessToken: "test-token",
		RequestBody: map[string]interface{}{"test": "data"},
		Timestamp:   time.Now(),
	}

	// Run multiple signatures concurrently
	const numGoroutines = 10
	results := make(chan string, numGoroutines)
	errors := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			sig, err := signer.Sign(payload)
			if err != nil {
				errors <- err
				return
			}
			results <- sig
		}()
	}

	// Collect results
	signatures := make([]string, 0, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		select {
		case sig := <-results:
			signatures = append(signatures, sig)
		case err := <-errors:
			t.Fatalf("Concurrent signature generation failed: %v", err)
		}
	}

	// All signatures should be identical for the same payload
	firstSig := signatures[0]
	for i, sig := range signatures {
		assert.Equal(t, firstSig, sig, "Signature %d should match first signature", i)
	}
}

// Benchmark tests for signature performance
func BenchmarkAsymmetricSigner_Sign(b *testing.B) {
	// Setup
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	tempDir := b.TempDir()
	privateKeyPath := tempDir + "/private_key.pem"
	publicKeyPath := tempDir + "/public_key.pem"

	// Write test keys
	privateKeyBytes, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	os.WriteFile(privateKeyPath, privateKeyPEM, 0600)

	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	os.WriteFile(publicKeyPath, publicKeyPEM, 0644)

	signer, _ := signature.NewAsymmetricSigner(privateKeyPEM, publicKeyPEM)

	payload := signature.AsymmetricPayload{
		ClientID:  "test-client-id",
		Timestamp: time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = signer.Sign(payload)
	}
}

func BenchmarkSymmetricSigner_Sign(b *testing.B) {
	signer := signature.NewSymmetricSigner("test-client-secret-12345")

	payload := signature.SymmetricPayload{
		HTTPMethod:  "POST",
		EndpointURL: "/api/v1.0/transfer-va/create-va",
		AccessToken: "sample-access-token",
		RequestBody: map[string]interface{}{
			"partnerServiceId": "test-service",
			"customerNo":       "123456789",
		},
		Timestamp: time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = signer.Sign(payload)
	}
}
