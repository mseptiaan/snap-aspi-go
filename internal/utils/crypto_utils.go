package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/mseptiaan/snap-aspi-go/internal/errors"
)

// CryptoUtils provides utility functions for cryptographic operations
type CryptoUtils struct{}

// NewCryptoUtils creates a new CryptoUtils instance
func NewCryptoUtils() *CryptoUtils {
	return &CryptoUtils{}
}

// LoadPrivateKey loads an RSA private key from a PEM file
func (c *CryptoUtils) LoadPrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	if keyPath == "" {
		return nil, errors.NewConfigurationError("private key path not configured")
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to read private key file: %v", err))
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.NewConfigurationError("failed to decode PEM block")
	}

	var privateKey *rsa.PrivateKey
	switch block.Type {
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, errors.NewConfigurationError(fmt.Sprintf("failed to parse private key: %v", err))
		}
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.NewConfigurationError("not an RSA private key")
		}
	default:
		return nil, errors.NewConfigurationError(fmt.Sprintf("unsupported key type: %s", block.Type))
	}

	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to parse private key: %v", err))
	}

	// Validate key
	if err := privateKey.Validate(); err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("invalid private key: %v", err))
	}

	return privateKey, nil
}

// LoadPublicKey loads an RSA public key from a PEM file
func (c *CryptoUtils) LoadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	if keyPath == "" {
		return nil, errors.NewConfigurationError("public key path not configured")
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to read public key file: %v", err))
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.NewConfigurationError("failed to decode PEM block")
	}

	var publicKey *rsa.PublicKey
	switch block.Type {
	case "RSA PUBLIC KEY":
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	case "PUBLIC KEY":
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, errors.NewConfigurationError(fmt.Sprintf("failed to parse public key: %v", err))
		}
		var ok bool
		publicKey, ok = key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.NewConfigurationError("not an RSA public key")
		}
	default:
		return nil, errors.NewConfigurationError(fmt.Sprintf("unsupported key type: %s", block.Type))
	}

	if err != nil {
		return nil, errors.NewConfigurationError(fmt.Sprintf("failed to parse public key: %v", err))
	}

	return publicKey, nil
}

// SignWithPrivateKey signs data with an RSA private key
func (c *CryptoUtils) SignWithPrivateKey(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	if privateKey == nil {
		return "", errors.NewConfigurationError("private key not configured")
	}

	// Hash the data
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)

	// Sign the hash
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	if err != nil {
		return "", errors.NewInternalError(err, "failed to generate RSA signature")
	}

	// Encode to base64
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyWithPublicKey verifies a signature with an RSA public key
func (c *CryptoUtils) VerifyWithPublicKey(publicKey *rsa.PublicKey, data []byte, signatureB64 string) error {
	if publicKey == nil {
		return errors.NewConfigurationError("public key not configured")
	}

	// Decode signature from base64
	signature, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return errors.NewValidationError(fmt.Sprintf("invalid base64 signature: %v", err))
	}

	// Hash the data
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature)
	if err != nil {
		return errors.NewAuthenticationError(fmt.Sprintf("signature verification failed: %v", err))
	}

	return nil
}

// GenerateRandomBytes generates random bytes of the specified length
func (c *CryptoUtils) GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, errors.NewInternalError(err, "failed to generate random bytes")
	}
	return bytes, nil
}

// GenerateRandomString generates a random string of the specified length
func (c *CryptoUtils) GenerateRandomString(length int) (string, error) {
	bytes, err := c.GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}