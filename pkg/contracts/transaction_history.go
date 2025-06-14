package contracts

import (
	"context"

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