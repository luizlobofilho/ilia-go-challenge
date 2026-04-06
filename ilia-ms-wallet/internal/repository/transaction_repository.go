package repository

import (
	"context"
	"wallet/internal/domain"
	"wallet/internal/ports"
)

// TransactionRepository implements ports.TransactionRepository using the DB adapter.
type TransactionRepository struct {
	DB ports.DatabaseConnection
}

func NewTransactionRepository(db ports.DatabaseConnection) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

var _ ports.TransactionRepository = (*TransactionRepository)(nil)

func (r *TransactionRepository) GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	var txs []domain.Transaction
	db := r.DB.GormDB().WithContext(ctx)
	if err := db.Where("user_id = ?", userID).Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

func (r *TransactionRepository) Create(ctx context.Context, t *domain.Transaction) error {
	db := r.DB.GormDB().WithContext(ctx)
	// GORM will set ID for new records. If you want ON CONFLICT behavior replace with Upsert logic.
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}
