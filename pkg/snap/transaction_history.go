package snap

import (
	"context"

	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// TransactionHistoryService defines the interface for Transaction History operations
type TransactionHistoryService interface {
	// Transaction History List
	TransactionHistoryList(ctx context.Context, payload *types.TransactionHistoryListPayload) (*types.TransactionHistoryResponse, error)
	
	// Transaction History Detail
	TransactionHistoryDetail(ctx context.Context, payload *types.TransactionHistoryDetailPayload) (map[string]any, error)
	
	// Bank Statement
	BankStatement(ctx context.Context, payload *types.BankStatementPayload) (*types.BankStatementResponse, error)
}

// transactionHistoryService implements TransactionHistoryService
type transactionHistoryService struct {
	svc *services.TransactionHistoryService
}

// TransactionHistoryList retrieves transaction history list
func (t *transactionHistoryService) TransactionHistoryList(ctx context.Context, payload *types.TransactionHistoryListPayload) (*types.TransactionHistoryResponse, error) {
	return t.svc.TransactionHistoryList(ctx, payload)
}

// TransactionHistoryDetail retrieves transaction history detail
func (t *transactionHistoryService) TransactionHistoryDetail(ctx context.Context, payload *types.TransactionHistoryDetailPayload) (map[string]any, error) {
	return t.svc.TransactionHistoryDetail(ctx, payload)
}

// BankStatement retrieves bank statement
func (t *transactionHistoryService) BankStatement(ctx context.Context, payload *types.BankStatementPayload) (*types.BankStatementResponse, error) {
	return t.svc.BankStatement(ctx, payload)
}