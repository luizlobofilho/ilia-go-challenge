package handler

import (
	"context"
	"net/http"
	"wallet/internal/domain"

	"github.com/gin-gonic/gin"
)

type TransactionUsecase interface {
	Create(ctx context.Context, t *domain.Transaction) error
	GetAllByUserID(ctx context.Context, userID string) ([]domain.Transaction, error)
}

var transactionUsecase TransactionUsecase

// SetTransactionUsecase allows setting the usecase implementation (e.g., for testing)
func SetTransactionUsecase(u TransactionUsecase) {
	transactionUsecase = u
}

// @Summary Get transactions for a user
// @Description Get all transactions for a user by their ID
// @Tags transactions
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /transactions/{id} [get]
func GetTransactions(c *gin.Context) {
	userID := c.Param("id")
	if transactionUsecase == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase not initialized"})
		return
	}

	txs, err := transactionUsecase.GetAllByUserID(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      userID,
		"transactions": txs,
	})
}

// @Summary Create a new transaction
// @Description Create a new transaction with the provided information
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body domain.Transaction true "Transaction to create"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	var t domain.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if transactionUsecase == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase not initialized"})
		return
	}

	if err := transactionUsecase.Create(context.Background(), &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created", "transaction": t})
}
