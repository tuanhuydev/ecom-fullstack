package dto

type CreateStoreDTO struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateStoreDTO struct {
	Name *string `json:"name" validate:"omitempty,min=3,max=100"`
}
