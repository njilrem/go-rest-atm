package controllers

import (
	// "github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/services"
	"github.com/njilrem/go-rest-atm/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	AuthAccount(ctx *gin.Context)
}

type authController struct {
	jwtService services.JWTService
}

func AuthHandler(jwtService services.JWTService) AuthController {
	return &authController{
		jwtService: jwtService,
	}
}

func (controller *authController) AuthAccount(ctx *gin.Context) {
	var credentials dto.AuthCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"token": "1"})
}