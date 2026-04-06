package usecase

import (
	"context"
	"testing"
	"wallet/internal/domain"
	"wallet/internal/ports"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct{ mock.Mock }

func (m *MockRepo) Create(ctx context.Context, t *domain.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockRepo) GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Transaction), args.Error(1)
}

var _ ports.TransactionRepository = (*MockRepo)(nil)

// TestCreate_Success verifies that Create successfully creates a valid transaction
func TestCreate_Success(t *testing.T) {
	mockRepo := &MockRepo{}
	uc := NewTransactionUsecase(mockRepo)

	tx := &domain.Transaction{UserID: "user1", Amount: 10.0, Type: "CREDIT"}
	mockRepo.On("Create", mock.Anything, tx).Return(nil)

	err := uc.Create(context.Background(), tx)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestCreate_InvalidTransaction verifies that Create returns an error when the transaction is invalid (e.g., missing UserID)
func TestCreate_InvalidTransaction(t *testing.T) {
	mockRepo := &MockRepo{}
	uc := NewTransactionUsecase(mockRepo)

	tx := &domain.Transaction{UserID: "", Amount: 1.0, Type: "DEBIT"}
	err := uc.Create(context.Background(), tx)
	assert.Equal(t, ErrInvalidTransaction, err)
}

// TestGetAllByUserID verifies that GetAllByUserID returns the expected transactions for a given user ID
func TestGetAllByUserID(t *testing.T) {
	mockRepo := &MockRepo{}
	uc := NewTransactionUsecase(mockRepo)

	expected := []domain.Transaction{{ID: 1, UserID: "user1", Amount: 5.5, Type: "CREDIT"}}
	mockRepo.On("GetAllByUserID", mock.Anything, "user1").Return(expected, nil)

	got, err := uc.GetAllByUserID(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	mockRepo.AssertExpectations(t)
}
