package controllers

import (
	"go-api/internal/dto"
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		ProductService: services.NewProductService(),
	}
}

func (c ProductController) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("products")
	routes.GET("", c.GetAllProducts)
	routes.GET(":id", c.GetProductById)
	routes.POST("", c.CreateProduct)
	routes.PUT(":id", c.UpdateProduct)
	routes.DELETE(":id", c.DeleteProduct)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	var query dto.PaginationQueryDTO
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
	}

	response, err := c.ProductService.GetAllProducts(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
func (c *ProductController) GetProductById(ctx *gin.Context) {

	product, err := c.ProductService.GetProductById(ctx.Param("id"))
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

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var body dto.CreateProductDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newProductId, error := c.ProductService.CreateProduct(body)

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

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var body dto.UpdateProductDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	error := c.ProductService.UpdateProduct(ctx.Param("id"), body)
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

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, error := c.ProductService.DeleteProduct(ctx.Param("id"))
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    id,
		"message": "Product deleted successfully",
	})
}
