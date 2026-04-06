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

// GetUserByID godoc
// @Summary Get user by id
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/{id} [get]
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User to create"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users [post]
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
