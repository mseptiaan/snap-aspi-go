package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
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

// registrationService implements RegistrationService
type registrationService struct {
	svc *services.RegistrationService
}

// Register performs user registration
func (r *registrationService) Register(ctx context.Context, payload *types.RegistrationPayload) (map[string]any, error) {
	return r.svc.Register(ctx, payload)
}

// RegisterCard performs card registration
func (r *registrationService) RegisterCard(ctx context.Context, payload *types.CardRegistrationPayload) (map[string]any, error) {
	return r.svc.RegisterCard(ctx, payload)
}

// BindAccount performs account binding
func (r *registrationService) BindAccount(ctx context.Context, payload *types.AccountBindingPayload) (map[string]any, error) {
	return r.svc.BindAccount(ctx, payload)
}

// VerifyOTP performs OTP verification
func (r *registrationService) VerifyOTP(ctx context.Context, payload *types.OTPVerificationPayload) (map[string]any, error) {
	return r.svc.VerifyOTP(ctx, payload)
}