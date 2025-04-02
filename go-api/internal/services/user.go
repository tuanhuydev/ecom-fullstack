package services

import (
	"fmt"
	"go-api/internal/database"
	"go-api/internal/dto"
	"go-api/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users %w", err)
	}
	return users, nil
}

func (s *UserService) GetUserByField(field string, value any) (models.User, error) {
	var user models.User
	result := database.DB.Where(field+" = ?", value).First(&user)
	if result.Error != nil {
		return models.User{}, fmt.Errorf("failed to get user by %s: %w", field, result.Error)
	}
	return user, nil
}

func (s *UserService) saveUser(body dto.CreateUserDTO) (models.User, error) {
	// Create the user
	if _, err := s.GetUserByField("email", body.Email); err == nil {
		return models.User{}, fmt.Errorf("user with email %s already exists", body.Email)

	}
	user := models.User{
		Name:  body.Name,
		Email: body.Email,
	}
	if createUserErr := database.DB.Create(&user).Error; createUserErr != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", createUserErr)
	}
	return user, nil

}

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, hashPasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if hashPasswordError != nil {
		return "", fmt.Errorf("failed to hash password: %w", hashPasswordError)
	}
	return string(hashedPassword), nil
}

func (s *UserService) CreateUser(body dto.CreateUserDTO) (string, error) {
	user, err := s.saveUser(body)
	if err != nil {
		return "", err
	}

	// Hash the password
	hashedPassword, hashPasswordErr := s.hashPassword(body.Password)

	if hashPasswordErr != nil {
		return "", hashPasswordErr
	}

	account := models.Account{
		Password: string(hashedPassword),
		UserId:   user.ID,
	}
	if createAccountErr := database.DB.Create(&account).Error; createAccountErr != nil {
		return "", createAccountErr
	}

	return user.ID, nil
}

func (s *UserService) UpdateUser(userId string, body dto.UpdateUserDTO) error {
	var existingUser models.User
	if err := database.DB.Where("id = ?", userId).First(&existingUser).Error; err != nil {
		return err
	}

	updatedUser := map[string]any{}
	if len(body.Name) > 0 {
		updatedUser["name"] = body.Name
	}

	// Update the product with the provided fields
	if err := database.DB.Model(&models.Product{}).Where("id = ?", userId).Updates(existingUser).Error; err != nil {
		return err
	}

	return nil

}

func (s *UserService) DeleteUser(userId string) (string, error) {
	updateUser := map[string]any{
		"deleted_at": time.Now(),
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", userId).Updates(updateUser).Error; err != nil {
		return "", err
	}
	return userId, nil
}
