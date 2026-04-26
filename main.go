package main

import (
	"log"
	"os"

	"ginBackend/config"
	"ginBackend/internal/controller"
	"ginBackend/internal/middleware"
	"ginBackend/internal/repository"
	"ginBackend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// main is the application entry point.
// It runs in order: load env → connect DB → configure routes → start server.
func main() {
	loadEnv()
	db := initDB()
	router := setupRouter(db)
	startServer(router)
}

// reads environment variables from a .env file if one is present.
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables from system")
	}
}

// establishes the database connection and runs auto-migrations
func initDB() *gorm.DB {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

// setupRouter wires all dependencies and registers all HTTP routes.
func setupRouter(db *gorm.DB) *gin.Engine {
	// Build the dependency chain bottom-up.
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// /user/password requires a valid JWT token.
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.PUT("/password", middleware.JWTAuth(), userController.ChangePassword)
	}

	return router
}

// startServer reads the PORT env var (default 8080) and starts listening.
func startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
