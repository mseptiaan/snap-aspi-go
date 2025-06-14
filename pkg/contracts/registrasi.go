package contracts

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// RegistrationService defines the interface for Registration operations
type RegistrationService interface {
	// User Registration
	Register(ctx context.Context, payload *types.RegistrationPayload) (map[string]any, error)
	
	// Card Registration
	RegisterCard(ctx context.Context, payload *types.CardRegistrationPayload) (map[string]any, error)
	
	// Account Binding
	BindAccount(ctx context.Context, payload *types.AccountBindingPayload) (map[string]any, error)
	
	// OTP Verification
	VerifyOTP(ctx context.Context, payload *types.OTPVerificationPayload) (map[string]any, error)
}