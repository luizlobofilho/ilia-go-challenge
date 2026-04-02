package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransactions returns a list of transactions for the authenticated user
func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transactions": []interface{}{},
	})
}
