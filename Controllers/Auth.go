package controllers

import (
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/services"
	"github.com/njilrem/go-rest-atm/dto"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	AuthAccount(ctx *gin.Context) string
}

type authController struct {
	jwtService services.JWTService
}

func AuthHandler(jwtService services.JWTService) AuthController {
	return &authController{
		jwtService: jwtService
	}
}

func AuthAccount(ctx *gin.Context) string {
	var credentials dto.AuthCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}