package controllers

import (
	"go-api/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(server *gin.Engine) {
	routes := server.Group("auth")
	routes.POST("/login", Login)
}

func Login(ctx *gin.Context) {
	var credentials dto.AuthEmailPasswordDTO
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// token, err := services.LoginService(credentials)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid credentials",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
}
