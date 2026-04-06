package repository

import (
	"context"
	"testing"
	"wallet/internal/domain"
	"wallet/internal/ports"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// testDBAdapter is a simple implementation of ports.DatabaseConnection for testing purposes
type testDBAdapter struct{ DB *gorm.DB }

func (a *testDBAdapter) GormDB() *gorm.DB { return a.DB }

// newInMemoryDB creates an in-memory *gorm.DB and applies necessary migrations
func newInMemoryDB(t *testing.T) ports.DatabaseConnection {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm sqlite: %v", err)
	}
	if err := db.AutoMigrate(&domain.Transaction{}); err != nil {
		t.Fatalf("failed migrate: %v", err)
	}
	return &testDBAdapter{DB: db}
}

// TestTransactionRepository_GetAllByUserID verifies that GetAllByUserID returns correct transactions for a given user ID
func TestTransactionRepository_GetAllByUserID(t *testing.T) {
	adapter := newInMemoryDB(t)
	db := adapter.(*testDBAdapter).GormDB()

	// fixture
	if err := db.Create(&domain.Transaction{UserID: "user1", Amount: 10.5, Type: "CREDIT"}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}
	if err := db.Create(&domain.Transaction{UserID: "user2", Amount: 3.0, Type: "DEBIT"}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}

	repo := NewTransactionRepository(adapter)
	txs, err := repo.GetAllByUserID(context.Background(), "user1")

	assert.NoError(t, err)
	assert.Len(t, txs, 1)
	assert.Equal(t, "user1", txs[0].UserID)
	assert.Equal(t, 10.5, txs[0].Amount)
	assert.Equal(t, "CREDIT", txs[0].Type)
}

// TestTransactionRepository_Create verifies that Create persists a new transaction and sets its ID
func TestTransactionRepository_Create(t *testing.T) {
	adapter := newInMemoryDB(t)
	repo := NewTransactionRepository(adapter)

	tx := &domain.Transaction{UserID: "user3", Amount: 5.0, Type: "DEBIT"}
	err := repo.Create(context.Background(), tx)
	assert.NoError(t, err)
	assert.NotZero(t, tx.ID)

	// verify persisted
	db := adapter.(*testDBAdapter).GormDB()
	var got domain.Transaction
	if err := db.First(&got, "user_id = ?", "user3").Error; err != nil {
		t.Fatalf("query persisted: %v", err)
	}
	assert.Equal(t, tx.UserID, got.UserID)
	assert.Equal(t, tx.Amount, got.Amount)
	assert.Equal(t, tx.Type, got.Type)
}
