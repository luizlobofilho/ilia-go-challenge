package main

import (
	"log"
	"os"
	"wallet/internal/handler"
	"wallet/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.JWTMiddleware())
	r.GET("/transactions", handler.GetTransactions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("Wallet service running on port %s", port)
	r.Run(":" + port)
}
