package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// TransferDebitService defines the interface for Transfer Debit operations
type TransferDebitService interface {
	// Direct Debit Payment
	DirectDebitPayment(ctx context.Context, payload *types.DirectDebitPaymentPayload) (map[string]any, error)
	
	// Direct Debit Status
	DirectDebitStatus(ctx context.Context, payload *types.DirectDebitStatusPayload) (map[string]any, error)
	
	// Direct Debit Cancel
	DirectDebitCancel(ctx context.Context, payload *types.DirectDebitCancelPayload) (map[string]any, error)
	
	// Direct Debit Refund
	DirectDebitRefund(ctx context.Context, payload *types.DirectDebitRefundPayload) (map[string]any, error)
	
	// CPM Generate QR
	CPMGenerateQR(ctx context.Context, payload *types.CPMGenerateQRPayload) (map[string]any, error)
	
	// CPM Payment
	CPMPayment(ctx context.Context, payload *types.CPMPaymentPayload) (map[string]any, error)
	
	// Auth Payment
	AuthPayment(ctx context.Context, payload *types.AuthPaymentPayload) (map[string]any, error)
	
	// Auth Capture
	AuthCapture(ctx context.Context, payload *types.AuthCapturePayload) (map[string]any, error)
	
	// Auth Void
	AuthVoid(ctx context.Context, payload *types.AuthVoidPayload) (map[string]any, error)
	
	// Direct Debit BI-FAST
	DirectDebitBIFAST(ctx context.Context, payload *types.DirectDebitBIFASTPayload) (map[string]any, error)
	
	// Direct Debit BI-FAST Payment
	DirectDebitBIFASTPayment(ctx context.Context, payload *types.DirectDebitBIFASTPaymentPayload) (map[string]any, error)
}

// transferDebitService implements TransferDebitService
type transferDebitService struct {
	svc *services.TransferDebitService
}

// DirectDebitPayment performs direct debit payment
func (t *transferDebitService) DirectDebitPayment(ctx context.Context, payload *types.DirectDebitPaymentPayload) (map[string]any, error) {
	return t.svc.DirectDebitPayment(ctx, payload)
}

// DirectDebitStatus checks direct debit status
func (t *transferDebitService) DirectDebitStatus(ctx context.Context, payload *types.DirectDebitStatusPayload) (map[string]any, error) {
	return t.svc.DirectDebitStatus(ctx, payload)
}

// DirectDebitCancel cancels direct debit
func (t *transferDebitService) DirectDebitCancel(ctx context.Context, payload *types.DirectDebitCancelPayload) (map[string]any, error) {
	return t.svc.DirectDebitCancel(ctx, payload)
}

// DirectDebitRefund refunds direct debit
func (t *transferDebitService) DirectDebitRefund(ctx context.Context, payload *types.DirectDebitRefundPayload) (map[string]any, error) {
	return t.svc.DirectDebitRefund(ctx, payload)
}

// CPMGenerateQR generates QR for CPM
func (t *transferDebitService) CPMGenerateQR(ctx context.Context, payload *types.CPMGenerateQRPayload) (map[string]any, error) {
	return t.svc.CPMGenerateQR(ctx, payload)
}

// CPMPayment performs CPM payment
func (t *transferDebitService) CPMPayment(ctx context.Context, payload *types.CPMPaymentPayload) (map[string]any, error) {
	return t.svc.CPMPayment(ctx, payload)
}

// AuthPayment performs auth payment
func (t *transferDebitService) AuthPayment(ctx context.Context, payload *types.AuthPaymentPayload) (map[string]any, error) {
	return t.svc.AuthPayment(ctx, payload)
}

// AuthCapture performs auth capture
func (t *transferDebitService) AuthCapture(ctx context.Context, payload *types.AuthCapturePayload) (map[string]any, error) {
	return t.svc.AuthCapture(ctx, payload)
}

// AuthVoid performs auth void
func (t *transferDebitService) AuthVoid(ctx context.Context, payload *types.AuthVoidPayload) (map[string]any, error) {
	return t.svc.AuthVoid(ctx, payload)
}

// DirectDebitBIFAST performs direct debit BI-FAST
func (t *transferDebitService) DirectDebitBIFAST(ctx context.Context, payload *types.DirectDebitBIFASTPayload) (map[string]any, error) {
	return t.svc.DirectDebitBIFAST(ctx, payload)
}

// DirectDebitBIFASTPayment performs direct debit BI-FAST payment
func (t *transferDebitService) DirectDebitBIFASTPayment(ctx context.Context, payload *types.DirectDebitBIFASTPaymentPayload) (map[string]any, error) {
	return t.svc.DirectDebitBIFASTPayment(ctx, payload)
}