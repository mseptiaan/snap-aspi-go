package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// VirtualAccountService defines the interface for Virtual Account operations
type VirtualAccountService interface {
	// CreateVA creates a new Virtual Account
	CreateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error)
	
	// UpdateVA updates an existing Virtual Account
	UpdateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error)
	
	// DeleteVA deletes a Virtual Account
	DeleteVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error)
	
	// InquiryVA performs Virtual Account inquiry
	InquiryVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error)
	
	// Inquiry performs general inquiry
	Inquiry(ctx context.Context, payload *types.InquiryPayload) (map[string]any, error)
	
	// Payment processes a Virtual Account payment
	Payment(ctx context.Context, payload *types.PaymentPayload) (map[string]any, error)
	
	// Status checks transaction status
	Status(ctx context.Context, payload *types.StatusPayload) (map[string]any, error)
	
	// Report generates transaction reports
	Report(ctx context.Context, payload *types.ReportPayload) (map[string]any, error)
	
	// UpdateStatus updates transaction status
	UpdateStatus(ctx context.Context, payload *types.StatusPayload) (map[string]any, error)
}

// virtualAccountService implements VirtualAccountService
type virtualAccountService struct {
	svc *services.VirtualAccountService
}

// CreateVA creates a new Virtual Account
func (v *virtualAccountService) CreateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error) {
	return v.svc.CreateVA(ctx, payload)
}

// UpdateVA updates an existing Virtual Account
func (v *virtualAccountService) UpdateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error) {
	return v.svc.UpdateVA(ctx, payload)
}

// DeleteVA deletes a Virtual Account
func (v *virtualAccountService) DeleteVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error) {
	return v.svc.DeleteVA(ctx, payload)
}

// InquiryVA performs Virtual Account inquiry
func (v *virtualAccountService) InquiryVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error) {
	return v.svc.InquiryVA(ctx, payload)
}

// Inquiry performs general inquiry
func (v *virtualAccountService) Inquiry(ctx context.Context, payload *types.InquiryPayload) (map[string]any, error) {
	return v.svc.Inquiry(ctx, payload)
}

// Payment processes a Virtual Account payment
func (v *virtualAccountService) Payment(ctx context.Context, payload *types.PaymentPayload) (map[string]any, error) {
	return v.svc.Payment(ctx, payload)
}

// Status checks transaction status
func (v *virtualAccountService) Status(ctx context.Context, payload *types.StatusPayload) (map[string]any, error) {
	return v.svc.Status(ctx, payload)
}

// Report generates transaction reports
func (v *virtualAccountService) Report(ctx context.Context, payload *types.ReportPayload) (map[string]any, error) {
	return v.svc.Report(ctx, payload)
}

// UpdateStatus updates transaction status
func (v *virtualAccountService) UpdateStatus(ctx context.Context, payload *types.StatusPayload) (map[string]any, error) {
	return v.svc.UpdateStatus(ctx, payload)
}