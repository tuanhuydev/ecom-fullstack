package database

import (
	"fmt"
	"go-api/internal/models"
	"log"
)

func RunMigrations() {
	fmt.Println("Running database migrations...")
	err := DB.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migrations completed successfully")
}
