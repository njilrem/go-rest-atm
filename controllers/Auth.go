package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/dto"
	"github.com/njilrem/go-rest-atm/models"
	"github.com/njilrem/go-rest-atm/services"
	log "github.com/sirupsen/logrus"
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
		log.Info("Account Authorized")
		token := services.JWTAuthService().GenerateToken(card.CardNum, card.ExpireDate, card.Cvv2, false)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

}

func AuthAccountPIN(c *gin.Context) {
	var credentials dto.AuthCredentialsPIN
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account models.Account
	if err := models.GetAccountByCardNum(credentials, &account); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if account.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Card[0].Pin), []byte(credentials.Pin)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Credentials!"})
		return
	} else {
		card := account.Card[0]
		log.Info("Account Authorized With Pin")
		token := services.JWTAuthService().GenerateToken(card.CardNum, card.ExpireDate, card.Cvv2, false)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"id": card.HolderID,
			"card_id": card.ID,
		})
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
		log.Info("Admin Account Authorized")
		token := services.JWTAuthService().GenerateToken(card.CardNum, card.ExpireDate, card.Cvv2, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}