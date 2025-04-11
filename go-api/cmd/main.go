package main

import (
	"fmt"
	"go-api/internal/controllers"
	"go-api/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func startServer() {
	fmt.Println("Starting application...")
	var server *gin.Engine = gin.Default()

	// CORS middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Controller declaration
	userController := controllers.NewUserController()
	userController.RegisterRoutes(server)

	productController := controllers.NewProductController()
	productController.RegisterRoutes(server)

	authController := controllers.NewAuthController()
	authController.RegisterAuthRoutes(server)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() {
	fmt.Println("Loading environment variables...")

	//  Load .env file first
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Validate required environment variables
	requiredEnvs := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("Required environment variable %s is not set", env)
		}
	}
}

func main() {
	loadEnv()
	database.ConnectDB()
	startServer()
}
