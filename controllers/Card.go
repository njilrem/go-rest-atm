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
	c.BindJSON(&card)
	err := models.CreateCard(&card)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
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
	c.BindJSON(&card)
	err = models.UpdateCard(&card)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, card)
	}
}
