package contracts

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// TransferCreditService defines the interface for Transfer Credit operations
type TransferCreditService interface {
	// Account Inquiry
	AccountInquiry(ctx context.Context, payload *types.AccountInquiryPayload) (map[string]any, error)
	
	// Trigger Transfer
	TriggerTransfer(ctx context.Context, payload *types.TriggerTransferPayload) (map[string]any, error)
	
	// Transfer Status Inquiry
	TransferStatusInquiry(ctx context.Context, payload *types.TransferStatusInquiryPayload) (map[string]any, error)
	
	// Customer Top Up
	CustomerTopUp(ctx context.Context, payload *types.CustomerTopUpPayload) (map[string]any, error)
	
	// Bulk Cashin
	BulkCashin(ctx context.Context, payload *types.BulkCashinPayload) (map[string]any, error)
	
	// Transfer To Bank
	TransferToBank(ctx context.Context, payload *types.TransferToBankPayload) (map[string]any, error)
	
	// Transfer To OTC
	TransferToOTC(ctx context.Context, payload *types.TransferToOTCPayload) (map[string]any, error)
}