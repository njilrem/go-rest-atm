package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/models"
	"net/http"
	"time"
)

func InitData(c *gin.Context){
	var account models.Account
	account.Name = "Dereka"
	account.Email = "derekaBood@mail.com"
	account.Phone = "+673829295333"
	account.Address = "Drezden city, Center"

	err := models.CreateAccount(&account)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var card models.Card
	card.CardNum = "5527532727647849"
	card.Balance = 500
	card.ExpireDate = "8/24/2028"
	card.HolderID = account.ID
	card.Pin = "2222"
	card.Cvv2 = "222"

	err = models.CreateCard(&card)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var transaction models.Transaction
	transaction.CardNum = "2221000230411810"
	transaction.CardID = card.ID
	transaction.TransactionDate = time.Now()
	transaction.Amount = 29.9

}
