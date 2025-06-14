package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// MPMService defines the interface for MPM (Merchant Payment Management) operations
type MPMService interface {
	// Transfer executes an MPM transfer credit transaction
	Transfer(ctx context.Context, payload *types.MPMTransferPayload) (map[string]any, error)
	
	// Inquiry performs MPM transfer inquiry
	Inquiry(ctx context.Context, payload *types.MPMInquiryPayload) (map[string]any, error)
	
	// Status checks MPM transfer status
	Status(ctx context.Context, payload *types.MPMStatusPayload) (map[string]any, error)
	
	// Refund processes MPM transfer refund
	Refund(ctx context.Context, payload *types.MPMRefundPayload) (map[string]any, error)
	
	// BalanceInquiry performs MPM balance inquiry
	BalanceInquiry(ctx context.Context, payload *types.MPMBalanceInquiryPayload) (map[string]any, error)
	
	// AccountInquiry performs MPM account inquiry
	AccountInquiry(ctx context.Context, payload *types.MPMAccountInquiryPayload) (map[string]any, error)
	
	// History retrieves MPM transaction history
	History(ctx context.Context, payload *types.MPMHistoryPayload) (map[string]any, error)
	
	// GenerateQR generates QR code for MPM payment
	GenerateQR(ctx context.Context, payload *types.MPMGenerateQRPayload) (*types.MPMQRResponse, error)
	
	// DecodeQR decodes a QR code for MPM payment
	DecodeQR(ctx context.Context, payload *types.MPMDecodeQRPayload) (map[string]any, error)
	
	// ApplyOTT applies a one-time token for payment redirect
	ApplyOTT(ctx context.Context, payload *types.MPMApplyOTTPayload) (map[string]any, error)
	
	// Payment processes a payment for MPM
	Payment(ctx context.Context, payload *types.MPMPaymentPayload) (map[string]any, error)
	
	// QueryPayment queries payment status for MPM
	QueryPayment(ctx context.Context, payload *types.MPMQueryPaymentPayload) (map[string]any, error)
	
	// CancelPayment cancels a payment for MPM
	CancelPayment(ctx context.Context, payload *types.MPMCancelPaymentPayload) (map[string]any, error)
	
	// NotifyQR handles QR MPM payment notification
	NotifyQR(ctx context.Context, payload *types.MPMNotifyQRPayload) (*types.MPMNotifyResponse, error)
}

// mpmService implements MPMService
type mpmService struct {
	svc *services.MPMService
}

// Transfer executes an MPM transfer credit transaction
func (m *mpmService) Transfer(ctx context.Context, payload *types.MPMTransferPayload) (map[string]any, error) {
	return m.svc.Transfer(ctx, payload)
}

// Inquiry performs MPM transfer inquiry
func (m *mpmService) Inquiry(ctx context.Context, payload *types.MPMInquiryPayload) (map[string]any, error) {
	return m.svc.Inquiry(ctx, payload)
}

// Status checks MPM transfer status
func (m *mpmService) Status(ctx context.Context, payload *types.MPMStatusPayload) (map[string]any, error) {
	return m.svc.Status(ctx, payload)
}

// Refund processes MPM transfer refund
func (m *mpmService) Refund(ctx context.Context, payload *types.MPMRefundPayload) (map[string]any, error) {
	return m.svc.Refund(ctx, payload)
}

// BalanceInquiry performs MPM balance inquiry
func (m *mpmService) BalanceInquiry(ctx context.Context, payload *types.MPMBalanceInquiryPayload) (map[string]any, error) {
	return m.svc.BalanceInquiry(ctx, payload)
}

// AccountInquiry performs MPM account inquiry
func (m *mpmService) AccountInquiry(ctx context.Context, payload *types.MPMAccountInquiryPayload) (map[string]any, error) {
	return m.svc.AccountInquiry(ctx, payload)
}

// History retrieves MPM transaction history
func (m *mpmService) History(ctx context.Context, payload *types.MPMHistoryPayload) (map[string]any, error) {
	return m.svc.History(ctx, payload)
}

// GenerateQR generates QR code for MPM payment
func (m *mpmService) GenerateQR(ctx context.Context, payload *types.MPMGenerateQRPayload) (*types.MPMQRResponse, error) {
	return m.svc.GenerateQR(ctx, payload)
}

// DecodeQR decodes a QR code for MPM payment
func (m *mpmService) DecodeQR(ctx context.Context, payload *types.MPMDecodeQRPayload) (map[string]any, error) {
	return m.svc.DecodeQR(ctx, payload)
}

// ApplyOTT applies a one-time token for payment redirect
func (m *mpmService) ApplyOTT(ctx context.Context, payload *types.MPMApplyOTTPayload) (map[string]any, error) {
	return m.svc.ApplyOTT(ctx, payload)
}

// Payment processes a payment for MPM
func (m *mpmService) Payment(ctx context.Context, payload *types.MPMPaymentPayload) (map[string]any, error) {
	return m.svc.Payment(ctx, payload)
}

// QueryPayment queries payment status for MPM
func (m *mpmService) QueryPayment(ctx context.Context, payload *types.MPMQueryPaymentPayload) (map[string]any, error) {
	return m.svc.QueryPayment(ctx, payload)
}

// CancelPayment cancels a payment for MPM
func (m *mpmService) CancelPayment(ctx context.Context, payload *types.MPMCancelPaymentPayload) (map[string]any, error) {
	return m.svc.CancelPayment(ctx, payload)
}

// NotifyQR handles QR MPM payment notification
func (m *mpmService) NotifyQR(ctx context.Context, payload *types.MPMNotifyQRPayload) (*types.MPMNotifyResponse, error) {
	return m.svc.NotifyQR(ctx, payload)
}