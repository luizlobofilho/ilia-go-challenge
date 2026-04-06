package domain

// Transaction represents domain model for wallet transaction
type Transaction struct {
	ID     int64   `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
