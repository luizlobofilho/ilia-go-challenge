package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"usersvc/internal/domain"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUsecase implements UserUsecase for tests
type MockUsecase struct{ mock.Mock }

func (m *MockUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestGetUserByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	m := &MockUsecase{}
	m.On("GetByID", mock.Anything, "123").Return(&domain.User{ID: 123, Email: "a@b.com", Name: "Alice"}, nil)

	SetUserUsecase(m)

	r := gin.Default()
	r.GET("/users/:id", GetUserByID)

	req, _ := http.NewRequest("GET", "/users/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
	m.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	m := &MockUsecase{}
	m.On("GetByID", mock.Anything, "404").Return(nil, assert.AnError)

	SetUserUsecase(m)
	r := gin.Default()
	r.GET("/users/:id", GetUserByID)

	req, _ := http.NewRequest("GET", "/users/404", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	m.AssertExpectations(t)
}
