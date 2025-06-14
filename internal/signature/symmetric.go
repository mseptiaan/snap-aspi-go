package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/errors"
)

// SymmetricPayload represents the payload for symmetric signatures
type SymmetricPayload struct {
	HTTPMethod  string
	EndpointURL string
	AccessToken string
	RequestBody interface{}
	Timestamp   time.Time
}

// String formats the payload for signing
// Format: HTTPMethod + ":" + EndpointUrl + ":" + AccessToken + ":" + Lowercase(HexEncode(SHA-256(minify(RequestBody)))) + ":" + TimeStamp
func (p SymmetricPayload) String() string {
	// Hash the request body
	bodyHash := p.hashRequestBody()

	// Format timestamp
	timestampStr := p.Timestamp.Format(time.RFC3339)

	// Build string to sign
	stringToSign := fmt.Sprintf("%s:%s:%s:%s:%s",
		p.HTTPMethod,
		p.EndpointURL,
		p.AccessToken,
		bodyHash,
		timestampStr,
	)

	return stringToSign
}

// hashRequestBody creates SHA-256 hash of the request body
func (p SymmetricPayload) hashRequestBody() string {
	if p.RequestBody == nil {
		return ""
	}

	// For now, we'll assume the body is already JSON-minified
	// In a production implementation, you might want to add JSON minification
	bodyStr := fmt.Sprintf("%v", p.RequestBody)

	// Create SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(bodyStr))
	hash := hasher.Sum(nil)

	// Convert to lowercase hex
	return strings.ToLower(hex.EncodeToString(hash))
}

// SymmetricSigner handles HMAC SHA512 signatures
type SymmetricSigner struct {
	clientSecret string
}

// NewSymmetricSigner creates a new symmetric signer
func NewSymmetricSigner(clientSecret string) *SymmetricSigner {
	return &SymmetricSigner{
		clientSecret: clientSecret,
	}
}

// Sign generates an HMAC SHA512 signature for the payload
func (s *SymmetricSigner) Sign(payload SymmetricPayload) (string, error) {
	if s.clientSecret == "" {
		return "", errors.NewConfigurationError("client secret not configured")
	}

	// Convert payload to string
	stringToSign := payload.String()

	// Create HMAC SHA512
	h := hmac.New(sha512.New, []byte(s.clientSecret))
	h.Write([]byte(stringToSign))
	signature := h.Sum(nil)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify verifies an HMAC SHA512 signature
func (s *SymmetricSigner) Verify(payload SymmetricPayload, signatureB64 string) error {
	if s.clientSecret == "" {
		return errors.NewConfigurationError("client secret not configured")
	}

	// Generate expected signature
	expectedSignature, err := s.Sign(payload)
	if err != nil {
		return errors.NewInternalError(err, "failed to generate expected signature")
	}

	// Compare signatures using constant-time comparison
	if !hmac.Equal([]byte(expectedSignature), []byte(signatureB64)) {
		return errors.NewAuthenticationError("signature verification failed")
	}

	return nil
}

// SignString signs a raw string (utility method)
func (s *SymmetricSigner) SignString(stringToSign string) (string, error) {
	if s.clientSecret == "" {
		return "", errors.NewConfigurationError("client secret not configured")
	}

	// Create HMAC SHA512
	h := hmac.New(sha512.New, []byte(s.clientSecret))
	h.Write([]byte(stringToSign))
	signature := h.Sum(nil)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyString verifies a signature for a raw string (utility method)
func (s *SymmetricSigner) VerifyString(stringToSign, signatureB64 string) error {
	if s.clientSecret == "" {
		return errors.NewConfigurationError("client secret not configured")
	}

	// Generate expected signature
	expectedSignature, err := s.SignString(stringToSign)
	if err != nil {
		return errors.NewInternalError(err, "failed to generate expected signature")
	}

	// Compare signatures using constant-time comparison
	if !hmac.Equal([]byte(expectedSignature), []byte(signatureB64)) {
		return errors.NewAuthenticationError("signature verification failed")
	}

	return nil
}
