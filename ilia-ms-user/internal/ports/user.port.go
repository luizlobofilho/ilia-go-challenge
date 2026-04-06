package ports

import (
	"context"
	"usersvc/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
}
