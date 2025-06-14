package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/errors"
)

// AsymmetricPayload represents the payload for asymmetric signatures
type AsymmetricPayload struct {
	ClientID  string
	Timestamp time.Time
}

// String formats the payload for signing
// Format: clientID + "|" + timestamp (ISO8601)
func (p AsymmetricPayload) String() string {
	return fmt.Sprintf("%s|%s", p.ClientID, p.Timestamp.Format(time.RFC3339))
}

// AsymmetricSigner handles RSA SHA256 signatures
type AsymmetricSigner struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewAsymmetricSigner creates a new asymmetric signer
func NewAsymmetricSigner(privateKeyPEM, publicKeyPEM []byte) (*AsymmetricSigner, error) {
	// Parse private key
	privateKey, err := parsePrivateKey(privateKeyPEM)
	if err != nil {
		return nil, errors.NewValidationError(fmt.Sprintf("failed to parse private key: %v", err))
	}

	// Parse public key
	publicKey, err := parsePublicKey(publicKeyPEM)
	if err != nil {
		return nil, errors.NewValidationError(fmt.Sprintf("failed to parse public key: %v", err))
	}

	return &AsymmetricSigner{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// NewAsymmetricSignerFromPrivateKey creates a signer with only private key (derives public key)
func NewAsymmetricSignerFromPrivateKey(privateKeyPEM []byte) (*AsymmetricSigner, error) {
	privateKey, err := parsePrivateKey(privateKeyPEM)
	if err != nil {
		return nil, errors.NewValidationError(fmt.Sprintf("failed to parse private key: %v", err))
	}

	return &AsymmetricSigner{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// Sign generates an RSA SHA256 signature for the payload
func (s *AsymmetricSigner) Sign(payload AsymmetricPayload) (string, error) {
	if s.privateKey == nil {
		return "", errors.NewConfigurationError("private key not configured")
	}

	// Convert payload to string
	stringToSign := payload.String()

	// Hash the payload
	hasher := sha256.New()
	hasher.Write([]byte(stringToSign))
	hash := hasher.Sum(nil)

	// Sign the hash
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hash)
	if err != nil {
		return "", errors.NewInternalError(err, "failed to generate RSA signature")
	}

	// Encode to base64
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify verifies an RSA SHA256 signature
func (s *AsymmetricSigner) Verify(payload AsymmetricPayload, signatureB64 string) error {
	if s.publicKey == nil {
		return errors.NewConfigurationError("public key not configured")
	}

	// Decode signature from base64
	signature, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return errors.NewValidationError(fmt.Sprintf("invalid base64 signature: %v", err))
	}

	// Convert payload to string
	stringToSign := payload.String()

	// Hash the payload
	hasher := sha256.New()
	hasher.Write([]byte(stringToSign))
	hash := hasher.Sum(nil)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(s.publicKey, crypto.SHA256, hash, signature)
	if err != nil {
		return errors.NewAuthenticationError(fmt.Sprintf("signature verification failed: %v", err))
	}

	return nil
}

// parsePrivateKey parses a PEM-encoded RSA private key
func parsePrivateKey(privateKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	var privateKey *rsa.PrivateKey
	var err error

	switch block.Type {
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA private key")
		}
	default:
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	if err != nil {
		return nil, err
	}

	// Validate key
	if err := privateKey.Validate(); err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	return privateKey, nil
}

// parsePublicKey parses a PEM-encoded RSA public key
func parsePublicKey(publicKeyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	var publicKey *rsa.PublicKey
	var err error

	switch block.Type {
	case "RSA PUBLIC KEY":
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	case "PUBLIC KEY":
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		publicKey, ok = key.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA public key")
		}
	default:
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
