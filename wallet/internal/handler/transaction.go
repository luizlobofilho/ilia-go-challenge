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

// GetTransactions returns all transactions (uses usecase when available)
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

// CreateTransaction create a new transaction (uses usecase when available)
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
