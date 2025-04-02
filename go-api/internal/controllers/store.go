package controllers

import (
	"go-api/internal/dto"
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	StoreService *services.StoreService
}

func NewStoreController() *StoreController {
	return &StoreController{
		StoreService: services.NewStoreService(),
	}
}

func (c *StoreController) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("stores")
	routes.GET("/", c.GetAllStores)
	routes.GET("/:id", c.GetStoreById)
	routes.POST("/", c.CreateStore)
	routes.PUT("/:id", c.UpdateStore)
	routes.DELETE("/:id", c.DeleteStore)
}

func (c *StoreController) GetAllStores(ctx *gin.Context) {
	query := dto.PaginationQueryDTO{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}

	stores, err := c.StoreService.GetAllStores(query)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve stores"})
		return
	}
	ctx.JSON(200, stores)
}
func (c *StoreController) GetStoreById(ctx *gin.Context) {
	id := ctx.Param("id")
	store, err := c.StoreService.GetStoreByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Store not found"})
		return
	}
	ctx.JSON(200, store)
}

func (c *StoreController) CreateStore(ctx *gin.Context) {
	var body dto.CreateStoreDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	createdStore, err := c.StoreService.CreateStore(body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create store"})
		return
	}
	ctx.JSON(201, createdStore)
}

func (c *StoreController) UpdateStore(ctx *gin.Context) {
	id := ctx.Param("id")
	var body dto.UpdateStoreDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedStore, err := c.StoreService.UpdateStore(id, body)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Store not found"})
		return
	}
	ctx.JSON(200, updatedStore)
}

func (c *StoreController) DeleteStore(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.StoreService.DeleteStore(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Store not found"})
		return
	}
	ctx.JSON(204, nil)
}
