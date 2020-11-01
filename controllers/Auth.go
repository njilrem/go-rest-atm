package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/dto"
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AuthAccount(c *gin.Context) {
	var credentials dto.AuthCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.Account
	if err := models.GetAccountByCredentials(credentials, &account); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if account.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Card[0].Cvv2), []byte(credentials.Cvv2)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Credentials!"})
		return
	} else {
		card := account.Card[0]
		token := services.JWTAuthService().GenerateToken(card.CardNum, card.ExpireDate, card.Cvv2, false)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

}

func AuthAdmin(c *gin.Context) {
	var credentials dto.AuthAdminCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.Account
	if err := models.GetAuthAccount(credentials, &account); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if account.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Card[0].Cvv2), []byte(credentials.Name)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Credentials!"})
		return
	} else {
		card := account.Card[0]
		token := services.JWTAuthService().GenerateToken(card.CardNum, card.ExpireDate, card.Cvv2, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}