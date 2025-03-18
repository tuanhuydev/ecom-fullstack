package main

import (
	"fmt"
	"go-api/internal/controllers"
	"go-api/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func startServer() {
	fmt.Println("Starting application...")
	var server *gin.Engine = gin.Default()
	controllers.RegisterUserRoutes(server)
	controllers.RegisterProductRoutes(server)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	server.Run()
}

func loadEnv() {
	fmt.Println("Loading environment variables...")

	//  Load .env file first
	err := godotenv.Load()
	if err != nil {
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
	database.RunMigrations()
	startServer()
}
