package main

import (
	"log"
	"os"

	"usersvc/internal/adapter/db"
	dbconfig "usersvc/internal/adapter/db/config"
	"usersvc/internal/handler"
	"usersvc/internal/middleware"
	"usersvc/internal/repository"
	"usersvc/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := dbconfig.GetDBConfig()
	adapter, err := db.NewDatabaseConnectionAdapter(cfg.DSN())
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	repo := repository.NewUserRepository(adapter)
	uc := usecase.NewUserUsecase(repo)
	handler.SetUserUsecase(uc)

	r := gin.Default()
	r.Use(middleware.JWTMiddleware())
	// expose GET /users/:id
	r.GET("/users/:id", handler.GetUserByID)
	r.POST("/users", handler.CreateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("Users service running on port %s", port)
	r.Run(":" + port)
}
