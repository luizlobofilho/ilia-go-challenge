package usecase

import (
	"context"
	"testing"
	"usersvc/internal/domain"
	"usersvc/internal/ports"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct{ mock.Mock }
func (m *MockRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*domain.User), args.Error(1)
}

var _ ports.UserRepository = (*MockRepo)(nil)

func TestGetByID(t *testing.T) {
	m := &MockRepo{}
	uc := NewUserUsecase(m)
	m.On("GetByID", mock.Anything, "123").Return(&domain.User{ID: 123, Email: "e@e.com", Name: "E"}, nil)

	u, err := uc.GetByID(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, "E", u.Name)
	m.AssertExpectations(t)
}
