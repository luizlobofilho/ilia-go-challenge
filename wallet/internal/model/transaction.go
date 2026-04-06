package model

import "wallet/internal/domain"

// Transaction represents a wallet transaction
type Transaction struct {
	ID     int     `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}

func (m Transaction) ToDomain() domain.Transaction {
	return domain.Transaction{
		ID:     int64(m.ID),
		UserID: m.UserID,
		Amount: m.Amount,
		Type:   m.Type,
	}
}
