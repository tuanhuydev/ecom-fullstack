package services

import (
	"go-api/internal/database"
	"go-api/internal/models"
	"log"
)

func AuthWithEmailPassword(email string, password string) (string, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return "", nil // TODO: Add authentication logic and return appropriate token
}
