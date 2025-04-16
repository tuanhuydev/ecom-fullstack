package controllers

import (
	"go-api/internal/database"
	"go-api/internal/models"
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CollectionController struct {
	ProductService *services.ProductService
}

func NewCollectionController(productService *services.ProductService) *CollectionController {
	return &CollectionController{
		ProductService: productService,
	}
}

func (c CollectionController) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("collections")
	routes.GET("/", c.GetAllCollections)
	routes.GET("/:id", c.GetCollectionById)
	routes.POST("/", c.CreateCollection)
	routes.PUT("/:id", c.UpdateCollection)
	routes.DELETE("/:id", c.DeleteCollection)
	routes.GET("/:id/products", c.GetProductsByCollection)

}

func (c CollectionController) GetAllCollections(ctx *gin.Context) {
	var collections []models.Collection
	if err := database.DB.Find(&collections).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch collections"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": collections})
}

func (c CollectionController) GetCollectionById(ctx *gin.Context) {
	id := ctx.Param("id")
	var collection models.Collection
	if err := database.DB.First(&collection, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": collection})
}

func (c CollectionController) CreateCollection(ctx *gin.Context) {
	var collection models.Collection
	if err := ctx.ShouldBindJSON(&collection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&collection).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create collection"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": collection})
}

func (c CollectionController) UpdateCollection(ctx *gin.Context) {
	id := ctx.Param("id")
	var collection models.Collection
	if err := database.DB.First(&collection, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}

	var updatedData models.Collection
	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&collection).Updates(updatedData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update collection"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": collection})
}

func (c CollectionController) DeleteCollection(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := database.DB.Delete(&models.Collection{}, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete collection"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Collection deleted successfully"})
}

func (c CollectionController) GetProductsByCollection(ctx *gin.Context) {
	collectionID := ctx.Param("id")

	products, err := c.ProductService.GetProductsByCollection(collectionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products for the collection"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}
