package services

import (
	"go-api/internal/database"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/pkg"
	"math"
	"time"

	"github.com/google/uuid"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAllProducts(query dto.PaginationQueryDTO) (pkg.PaginatedResponse, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.SortBy == "" {
		query.SortBy = "created_at"
	}
	if query.SortOrder == "" {
		query.SortOrder = "DESC"
	}
	db := database.DB.Model(&models.Product{}).Where("deleted_at IS NULL")

	// Count
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return pkg.PaginatedResponse{}, err
	}

	// Pagination
	offset := (query.Page - 1) * query.PageSize
	totalPages := int(math.Ceil(float64(totalCount) / float64(query.PageSize)))

	var products []models.Product
	if err := database.DB.Order(query.SortBy + " " + query.SortOrder).Limit(query.PageSize).Offset(offset).Find(&products).Error; err != nil {
		return pkg.PaginatedResponse{}, err
	}
	return pkg.PaginatedResponse{
		Data: products,
		Pagination: pkg.PaginationMeta{
			CurrentPage:  query.Page,
			PageSize:     query.PageSize,
			TotalRecords: totalCount,
			TotalPages:   totalPages,
		},
	}, nil
}

func (s *ProductService) GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *ProductService) CreateProduct(body dto.CreateProductDTO) (string, error) {
	productModel := models.Product{
		ID:       uuid.New().String(),
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
	}
	if body.Thumbnail != nil && len(*body.Thumbnail) > 0 {
		productModel.Thumbnail = body.Thumbnail
	}

	if err := database.DB.Create(&productModel).Error; err != nil {
		return "", err
	}

	return productModel.ID, nil
}

func (s *ProductService) UpdateProduct(id string, body dto.UpdateProductDTO) error {
	var existingProduct models.Product
	if err := database.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
		return err
	}

	updatedProduct := map[string]any{}
	if body.Name != nil && len(*body.Name) > 0 {
		updatedProduct["name"] = body.Name
	}

	if body.Price != nil && *body.Price > 0 {
		updatedProduct["price"] = body.Price
	}

	if body.Quantity != nil && *body.Quantity >= 0 {
		updatedProduct["quantity"] = body.Quantity
	}

	if body.Thumbnail != nil && len(*body.Thumbnail) > 0 {
		updatedProduct["thumbnail"] = body.Thumbnail
	}

	// Update the product with the provided fields
	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(id string) (string, error) {
	updateProduct := map[string]any{
		"deleted_at": time.Now(),
	}

	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updateProduct).Error; err != nil {
		return "", err
	}
	return id, nil
}
