package dto

type AuthEmailPasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserDTO struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
