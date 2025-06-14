package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// BalanceInquiryService defines the interface for Balance Inquiry operations
type BalanceInquiryService interface {
	// Balance Inquiry
	BalanceInquiry(ctx context.Context, payload *types.BalanceInquiryPayload) (*types.BalanceInquiryResponse, error)
}

// balanceInquiryService implements BalanceInquiryService
type balanceInquiryService struct {
	svc *services.BalanceInquiryService
}

// BalanceInquiry performs balance inquiry
func (b *balanceInquiryService) BalanceInquiry(ctx context.Context, payload *types.BalanceInquiryPayload) (*types.BalanceInquiryResponse, error) {
	return b.svc.BalanceInquiry(ctx, payload)
}