package controllers

import (
	"fmt"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{UserService: services.NewUserService()}
}

func (c *UserController) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("users")
	routes.GET("/", c.GetAllUsers)
	routes.GET("/:id", c.GetUserById)
	routes.POST("/", c.CreateUser)
	routes.PUT("/:id", c.UpdateUser)
	routes.DELETE("/:id", c.DeleteUser)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": []models.User{},
	})
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var user models.User
	user, err := c.UserService.GetUserByField("id", id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var body dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.UserService.CreateUser(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	fmt.Println("Update User")
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	fmt.Println("Delete User")
}
