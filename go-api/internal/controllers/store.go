package controllers

import (
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

func (c *StoreController) GetAllStores(ctx *gin.Context) {}
func (c *StoreController) GetStoreById(ctx *gin.Context) {}
func (c *StoreController) CreateStore(ctx *gin.Context)  {}
func (c *StoreController) UpdateStore(ctx *gin.Context)  {}
func (c *StoreController) DeleteStore(ctx *gin.Context)  {}
