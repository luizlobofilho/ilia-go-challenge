package handler

import (
	"context"
	"net/http"
	"usersvc/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
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
