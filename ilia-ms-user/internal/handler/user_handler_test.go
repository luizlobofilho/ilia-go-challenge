package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"usersvc/internal/domain"
	"usersvc/internal/usecase"

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

func (m *MockUsecase) Create(ctx context.Context, u *domain.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
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

func TestCreateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	m := &MockUsecase{}
	m.On("Create", mock.Anything, mock.Anything).Return(nil)

	SetUserUsecase(m)
	r := gin.Default()
	r.POST("/users", CreateUser)

	body := `{"email":"a@b.com","name":"Alice"}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	m.AssertExpectations(t)
}

func TestCreateUser_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	SetUserUsecase(nil)
	r := gin.Default()
	r.POST("/users", CreateUser)

	req, _ := http.NewRequest("POST", "/users", strings.NewReader(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestCreateUser_InvalidUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	m := &MockUsecase{}
	m.On("Create", mock.Anything, mock.Anything).Return(usecase.ErrInvalidUser)

	SetUserUsecase(m)
	r := gin.Default()
	r.POST("/users", CreateUser)

	body := `{"email":"","name":""}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	m.AssertExpectations(t)
}
