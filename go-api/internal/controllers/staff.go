package controllers

import (
	"go-api/internal/services"

	"github.com/gin-gonic/gin"
)

type StaffController struct {
	StaffService *services.StaffService
}

func NewStaffController() *StaffController {
	return &StaffController{
		StaffService: services.NewStaffService(),
	}
}

func (c *StaffController) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("staff")
	routes.GET("/", c.GetAllStaff)
	routes.GET("/:id", c.GetStaffById)
	routes.POST("/", c.CreateStaff)
	routes.PUT("/:id", c.UpdateStaff)
	routes.DELETE("/:id", c.DeleteStaff)
}

func (c *StaffController) GetAllStaff(ctx *gin.Context)  {}
func (c *StaffController) GetStaffById(ctx *gin.Context) {}
func (c *StaffController) CreateStaff(ctx *gin.Context)  {}
func (c *StaffController) UpdateStaff(ctx *gin.Context)  {}
func (c *StaffController) DeleteStaff(ctx *gin.Context)  {}
