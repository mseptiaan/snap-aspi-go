package contracts

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// VirtualAccountService defines the interface for Virtual Account operations
// Equivalent to PHP VirtualAccount contract
type VirtualAccountService interface {
	// Virtual Account Management
	CreateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error)
	UpdateVA(ctx context.Context, payload *types.CreateVAPayload) (map[string]any, error)
	DeleteVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error)

	// Virtual Account Inquiries
	InquiryVA(ctx context.Context, payload *types.InquiryVAPayload) (map[string]any, error)
	Inquiry(ctx context.Context, payload *types.InquiryPayload) (map[string]any, error)

	// Payment Operations
	Payment(ctx context.Context, payload *types.PaymentPayload) (map[string]any, error)

	// Status and Reporting
	Status(ctx context.Context, payload *types.StatusPayload) (map[string]any, error)
	Report(ctx context.Context, payload *types.ReportPayload) (map[string]any, error)
	UpdateStatus(ctx context.Context, payload *types.StatusPayload) (map[string]any, error)
}
