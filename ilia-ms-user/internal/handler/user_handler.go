package handler

import (
	"context"
	"net/http"
	"usersvc/internal/domain"
	"usersvc/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
}

var userUsecase UserUsecase

func SetUserUsecase(u UserUsecase) { userUsecase = u }

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if userUsecase == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase not initialized"})
		return
	}
	u, err := userUsecase.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func CreateUser(c *gin.Context) {
	var u domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if userUsecase == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase not initialized"})
		return
	}
	if err := userUsecase.Create(context.Background(), &u); err != nil {
		if err == usecase.ErrInvalidUser {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created", "user": u})
}
