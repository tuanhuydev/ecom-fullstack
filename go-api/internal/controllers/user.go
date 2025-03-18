package controllers

import (
	"fmt"
	"go-api/internal/database"
	"go-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	routes := server.Group("users")
	routes.GET("/", GetAllUsers)
	routes.GET("/:id", GetUserById)
	routes.POST("/", CreateUser)
	routes.PUT("/:id", UpdateUser)
	routes.DELETE("/:id", DeleteUser)
}

func GetAllUsers(ctx *gin.Context) {
	mockData := [5]int{1, 2, 3, 4, 5}
	ctx.JSON(http.StatusOK, gin.H{
		"data": mockData,
	})
}

func GetUserById(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	err := result.Error
	fmt.Println(user)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}

}

func CreateUser(ctx *gin.Context) {
	fmt.Println("Create User")
}

func UpdateUser(ctx *gin.Context) {
	fmt.Println("Update User")
}

func DeleteUser(ctx *gin.Context) {
	fmt.Println("Delete User")
}
