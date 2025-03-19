package dto

type CreateProductDTO struct {
	Name      string  `json:"name" validate:"required,min=3,max=50"`
	Price     float64 `json:"price" validate:"required,gt=0"`
	Quantity  int     `json:"quantity" validate:"required,gte=0"`
	Thumbnail *string `json:"thumbnail" validate:"omitempty,url"`
}
type UpdateProductDTO struct {
	Name      *string  `json:"name" validate:"omitempty,min=3,max=50"`
	Price     *float64 `json:"price" validate:"omitempty,gt=0"`
	Quantity  *int     `json:"quantity" validate:"omitempty,gte=0"`
	Thumbnail *string  `json:"thumbnail" validate:"omitempty,url"`
}

type ProductQueryDTO struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
	SortBy    string `form:"sortBy"`
	SortOrder string `form:"sortOrder"`
}
