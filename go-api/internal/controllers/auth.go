package controllers

import (
	"go-api/internal/dto"
	"go-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: services.NewAuthService(),
	}
}

func (c *AuthController) RegisterAuthRoutes(server *gin.Engine) {
	routes := server.Group("auth")
	routes.POST("/sign-in", c.SignInWithEmailPassword)
	routes.POST("/sign-up", c.SignUp)
}

func (c *AuthController) SignInWithEmailPassword(ctx *gin.Context) {
	var credentials dto.AuthEmailPasswordDTO
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := c.AuthService.AuthWithEmailPassword(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var body dto.RegisterUserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.AuthService.SignUp(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}
