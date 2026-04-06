package usecase

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/ports"
)

type TransactionUsecase struct {
	repo ports.TransactionRepository
}

func NewTransactionUsecase(r ports.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{repo: r}
}

// Create validates the transaction and delegates to the repository.
func (u *TransactionUsecase) Create(ctx context.Context, t *domain.Transaction) error {
	if t.UserID == "" {
		return ErrInvalidTransaction
	}
	return u.repo.Create(ctx, t)
}

// GetAllByUserID retrieves all transactions for a given user by delegating to the repository.
func (u *TransactionUsecase) GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	return u.repo.GetAllByUserID(ctx, userID)
}

// Domain-level errors
var ErrInvalidTransaction = &domainError{"invalid transaction"}

type domainError struct{ msg string }

func (e *domainError) Error() string { return e.msg }
