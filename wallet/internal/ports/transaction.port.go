package ports

import (
	"context"
	"wallet/internal/domain"
)

// Transaction defines the methods used by usecases
type TransactionRepository interface {
	Create(ctx context.Context, t *domain.Transaction) error
	GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error)
}
