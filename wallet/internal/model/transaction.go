package model

// Transaction represents a wallet transaction
type Transaction struct {
	ID     int     `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
