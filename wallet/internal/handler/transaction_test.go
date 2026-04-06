package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wallet/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository to mock the TransactionUsecase for testing handlers
type MockRepository struct{ mock.Mock }

func (m *MockRepository) Create(ctx context.Context, t *domain.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockRepository) GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Transaction), args.Error(1)
}

// TestGetTransactions verify if the handler returns transactions correctly
func TestGetTransactions(t *testing.T) {
	mockRepo := &MockRepository{}
	mockRepo.On("GetAllByUserID", mock.Anything, "123").Return([]domain.Transaction{}, nil)
	SetTransactionUsecase(mockRepo)

	r := gin.Default()
	r.GET("/transactions/:id", GetTransactions)

	req, _ := http.NewRequest("GET", "/transactions/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
	assert.Contains(t, w.Body.String(), "123")
	mockRepo.AssertExpectations(t)
}

// TestCreateTransaction_Success verifies if the handler creates a transaction successfully
func TestCreateTransaction_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := &MockRepository{}
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	SetTransactionUsecase(mockRepo)

	r := gin.Default()
	r.POST("/transactions", CreateTransaction)

	body := `{"user_id":"user1","amount":10.5,"type":"CREDIT"}`
	req, _ := http.NewRequest("POST", "/transactions", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	mockRepo.AssertExpectations(t)
}

// TestCreateTransaction_InvalidBody verifies if the handler returns 400 for an invalid body
func TestCreateTransaction_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	SetTransactionUsecase(nil)

	r := gin.Default()
	r.POST("/transactions", CreateTransaction)

	req, _ := http.NewRequest("POST", "/transactions", strings.NewReader(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
