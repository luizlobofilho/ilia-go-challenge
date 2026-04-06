package repository

import (
	"context"
	"usersvc/internal/domain"
	"usersvc/internal/ports"
)

type UserRepository struct{ DB ports.DatabaseConnection }

func NewUserRepository(db ports.DatabaseConnection) *UserRepository { return &UserRepository{DB: db} }

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var u domain.User
	db := r.DB.GormDB().WithContext(ctx)
	if err := db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	db := r.DB.GormDB().WithContext(ctx)
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}
