package main

import (
	"log"
	"os"

	database "wallet/internal/adapter/db"
	dbconfig "wallet/internal/adapter/db/config"
	"wallet/internal/handler"
	"wallet/internal/middleware"
	"wallet/internal/repository"
	"wallet/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB and create adapter/repository
	cfg := dbconfig.GetDBConfig()
	adapter, err := database.NewDatabaseConnectionAdapter(cfg.DSN())
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	repoAdapter := repository.NewTransactionRepository(adapter)
	uc := usecase.NewTransactionUsecase(repoAdapter)
	handler.SetTransactionUsecase(uc)

	r := gin.Default()
	r.Use(middleware.JWTMiddleware())
	r.GET("/transactions/:id", handler.GetTransactions)
	r.POST("/transactions", handler.CreateTransaction)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("Wallet service running on port %s", port)
	r.Run(":" + port)
}
