package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/models"
	"log"
	"net/http"
)

func GetTransactionById(c *gin.Context) {
	id := c.Params.ByName("id")
	var transaction models.Transaction
	err := models.GetTransactionById(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}

func GetTransactionsById(c *gin.Context) {
	id := c.Params.ByName("id")
	var transactions []models.Transaction
	err := models.GetTransactionsByCardId(&transactions, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"transactions": transactions,
			"quantity": len(transactions),
		})
	}
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = models.CreateTransaction(&transaction)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		err = models.ProcessTransaction(transaction)
		if err != nil {
			log.Println("TRANSACTION FAILED")
			err = models.DeleteTransaction(&transaction)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			c.JSON(http.StatusOK, transaction)
		}
	}
}

func Refill(c *gin.Context) {
	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(transaction)
	err = models.CreateTransaction(&transaction)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		err = models.ProcessRefillTransaction(transaction)
		if err != nil {
			log.Println("TRANSACTION FAILED")
			err = models.DeleteTransaction(&transaction)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			c.JSON(http.StatusOK, transaction)
		}
	}

}