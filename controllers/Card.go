package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/models"
	"net/http"
)

func GetCardsByHolderID(c *gin.Context) {
	id := c.Params.ByName("id")
	var card []models.Card
	err := models.GetCardsByHolderId(&card, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, card)
	}
}

func CreateCard(c *gin.Context) {
	var card models.Card
	_ = c.BindJSON(&card)
	err := models.CreateCard(&card)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, card)
	}
}

func UpdateCard(c *gin.Context) {
	var card models.Card
	id := c.Params.ByName("id")
	err := models.GetCardById(&card, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	_ = c.BindJSON(&card)
	err = models.UpdateCard(&card)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, card)
	}
}

func GetBalance(c *gin.Context) {
	id := c.Params.ByName("id")
	var card models.Card
	err := models.GetCardById(&card, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"balance": card.Balance,
			"id": card.ID,
			"card_num": card.CardNum,
			"holder_id": card.HolderID,
		})
	}
}