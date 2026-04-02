package repository

import (
	"context"
	"wallet/internal/model"

	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	Conn *pgx.Conn
}

func NewTransactionRepository(conn *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{Conn: conn}
}

func (r *TransactionRepository) GetAll(ctx context.Context) ([]model.Transaction, error) {
	rows, err := r.Conn.Query(ctx, "SELECT id, user_id, amount, type FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var txs []model.Transaction
	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type); err != nil {
			return nil, err
		}
		txs = append(txs, t)
	}
	return txs, nil
}

func (r *TransactionRepository) Create(ctx context.Context, t model.Transaction) error {
	_, err := r.Conn.Exec(ctx, "INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, $3)", t.UserID, t.Amount, t.Type)
	return err
}
