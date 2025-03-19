package pkg

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Loading environment variables...")

	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	// Validate required environment variables
	requiredEnvs := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("Required environment variable %s is not set", env)
		}
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func CalculatePagination(page, pageSize int, totalCount int64) PaginationMeta {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return PaginationMeta{
		CurrentPage:  page,
		PageSize:     pageSize,
		TotalRecords: totalCount,
		TotalPages:   int(math.Ceil(float64(totalCount) / float64(pageSize))),
	}
}
