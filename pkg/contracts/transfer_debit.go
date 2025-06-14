package contracts

import (
	"context"

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