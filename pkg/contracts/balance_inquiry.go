package contracts

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// BalanceInquiryService defines the interface for Balance Inquiry operations
type BalanceInquiryService interface {
	// Balance Inquiry
	BalanceInquiry(ctx context.Context, payload *types.BalanceInquiryPayload) (*types.BalanceInquiryResponse, error)
}