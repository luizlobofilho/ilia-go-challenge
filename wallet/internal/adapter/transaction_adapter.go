package adapter

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/ports"
)

// TransactionRepoAdapter wraps an existing TransactionRepository and exposes the ports.TransactionRepository
// interface. Repository now works with domain.Transaction so the adapter is a simple passthrough.
type TransactionRepoAdapter struct {
	inner ports.TransactionRepository
}

func NewTransactionRepoAdapter(inner ports.TransactionRepository) *TransactionRepoAdapter {
	return &TransactionRepoAdapter{inner: inner}
}

var _ ports.TransactionRepository = (*TransactionRepoAdapter)(nil)

func (a *TransactionRepoAdapter) Create(ctx context.Context, t *domain.Transaction) error {
	return a.inner.Create(ctx, t)
}

func (a *TransactionRepoAdapter) GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	return a.inner.GetAllByUserID(ctx, userID)
}
