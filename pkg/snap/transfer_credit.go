package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
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

// transferCreditService implements TransferCreditService
type transferCreditService struct {
	svc *services.TransferCreditService
}

// AccountInquiry performs account inquiry
func (t *transferCreditService) AccountInquiry(ctx context.Context, payload *types.AccountInquiryPayload) (map[string]any, error) {
	return t.svc.AccountInquiry(ctx, payload)
}

// TriggerTransfer performs transfer
func (t *transferCreditService) TriggerTransfer(ctx context.Context, payload *types.TriggerTransferPayload) (map[string]any, error) {
	return t.svc.TriggerTransfer(ctx, payload)
}

// TransferStatusInquiry checks transfer status
func (t *transferCreditService) TransferStatusInquiry(ctx context.Context, payload *types.TransferStatusInquiryPayload) (map[string]any, error) {
	return t.svc.TransferStatusInquiry(ctx, payload)
}

// CustomerTopUp performs customer top up
func (t *transferCreditService) CustomerTopUp(ctx context.Context, payload *types.CustomerTopUpPayload) (map[string]any, error) {
	return t.svc.CustomerTopUp(ctx, payload)
}

// BulkCashin performs bulk cashin
func (t *transferCreditService) BulkCashin(ctx context.Context, payload *types.BulkCashinPayload) (map[string]any, error) {
	return t.svc.BulkCashin(ctx, payload)
}

// TransferToBank performs transfer to bank
func (t *transferCreditService) TransferToBank(ctx context.Context, payload *types.TransferToBankPayload) (map[string]any, error) {
	return t.svc.TransferToBank(ctx, payload)
}

// TransferToOTC performs transfer to OTC
func (t *transferCreditService) TransferToOTC(ctx context.Context, payload *types.TransferToOTCPayload) (map[string]any, error) {
	return t.svc.TransferToOTC(ctx, payload)
}