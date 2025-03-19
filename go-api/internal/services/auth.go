package services

import (
	"errors"
	"go-api/internal/database"
	"go-api/internal/dto"
	"go-api/internal/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) AuthWithEmailPassword(email string, password string) (string, error) {
	var user models.User
	var account models.Account
	userErr := database.DB.Where("email = ?", email).First(&user).Error
	if userErr != nil {
		log.Fatal(userErr)
	}
	accountErr := database.DB.Where("user_id = ?", user.ID).First(&account).Error
	if accountErr != nil {
		log.Fatal(accountErr)
	}

	// Compare the password with the hashed password in the database
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if bcryptErr != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (s *AuthService) RegisterUser(body dto.CreateUserDTO) error {
	var user models.User
	var account models.Account

	// Start a transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create the user
	user = models.User{
		ID:    uuid.New().String(),
		Name:  body.Name,
		Email: body.Email,
	}
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create the account
	account = models.Account{
		UserId:   user.ID,
		Password: string(hashedPassword),
	}
	err = tx.Create(&account).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
