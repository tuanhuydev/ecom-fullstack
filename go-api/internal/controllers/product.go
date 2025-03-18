package controllers

import (
	"go-api/internal/dto"
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(server *gin.Engine) {
	routes := server.Group("products")
	routes.GET("/", GetAllProducts)
	routes.GET("/:id", GetProductById)
	routes.POST("/", CreateProduct)
	routes.PUT("/:id", UpdateProduct)
	routes.DELETE("/:id", DeleteProduct)
}

func GetAllProducts(ctx *gin.Context) {
	products, err := services.GetAllProductsService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
func GetProductById(ctx *gin.Context) {

	product, err := services.GetProductByIdService(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func CreateProduct(ctx *gin.Context) {
	var body dto.CreateProductDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newProductId, error := services.CreateProductService(body)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    newProductId,
	})
}

func UpdateProduct(ctx *gin.Context) {
	var body dto.UpdateProductDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	error := services.UpdateProductService(ctx.Param("id"), body)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    ctx.Param("id"),
		"message": "Product updated successfully",
	})
}

func DeleteProduct(ctx *gin.Context) {
	error := services.DeleteProductService(ctx.Param("id"))
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    ctx.Param("id"),
		"message": "Product deleted successfully",
	})
}
