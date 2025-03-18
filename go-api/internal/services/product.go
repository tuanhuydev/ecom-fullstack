package services

import (
	"go-api/internal/database"
	"go-api/internal/dto"
	"go-api/internal/models"
	"time"

	"github.com/google/uuid"
)

func GetAllProductsService() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByIdService(id string) (models.Product, error) {
	var product models.Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func CreateProductService(body dto.CreateProductDTO) (string, error) {
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

func UpdateProductService(id string, body dto.UpdateProductDTO) error {
	var existingProduct models.Product
	if err := database.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
		return err
	}

	updatedProduct := map[string]interface{}{}
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

func DeleteProductService(id string) error {
	// Soft delete the product
	updateProduct := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updateProduct).Error; err != nil {
		return err
	}

	return nil
}
