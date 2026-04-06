package usecase

import (
	"context"
	"usersvc/internal/domain"
	"usersvc/internal/ports"
)

type UserUsecase struct{ repo ports.UserRepository }

func NewUserUsecase(r ports.UserRepository) *UserUsecase { return &UserUsecase{repo: r} }

// GetByID retrieves a user by ID, delegating to the repository.
func (u *UserUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

// Create validates the user and delegates persistence to the repository.
func (u *UserUsecase) Create(ctx context.Context, user *domain.User) error {
	if user == nil || user.Email == "" || user.Name == "" {
		return ErrInvalidUser
	}
	return u.repo.Create(ctx, user)
}

var ErrInvalidUser = &userError{"invalid user"}

type userError struct{ msg string }

func (e *userError) Error() string { return e.msg }
