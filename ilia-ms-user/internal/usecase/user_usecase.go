package usecase

import (
	"context"
	"usersvc/internal/domain"
	"usersvc/internal/ports"
)

type UserUsecase struct{ repo ports.UserRepository }

func NewUserUsecase(r ports.UserRepository) *UserUsecase { return &UserUsecase{repo: r} }

func (u *UserUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}
