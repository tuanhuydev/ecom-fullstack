package dto

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=50"`
}

type UpdateUserDTO struct {
	Name string `json:"name" validate:"omitempty,min=3,max=50"`
}
