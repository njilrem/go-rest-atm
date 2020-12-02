package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/config"
	"github.com/njilrem/go-rest-atm/models"
	"net/http"
	"time"
)

func InitData(c *gin.Context){
	var account models.Account
	account.Name = "Dereka"
	account.Email = "derekaBood@mail.com"
	account.Phone = "673829295333"
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
	transaction.Comment = "First Transaction"

	err = models.CreateTransaction(&transaction)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var accountL models.Account
	accountL.Name = "Dereka Brother"
	accountL.Email = "derekaBoodAA@mail.com"
	accountL.Phone = "673829295333"
	accountL.Address = "Drezden city, Center"

	err = models.CreateAccount(&accountL)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var cardL models.Card
	cardL.CardNum = "5507770244348111"
	cardL.Balance = 500
	cardL.ExpireDate = "8/24/2028"
	cardL.HolderID = accountL.ID
	cardL.Pin = "2222"
	cardL.Cvv2 = "222"

	err = models.CreateCard(&cardL)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var transactionL models.Transaction
	transactionL.CardNum = "2221000230411810"
	transactionL.CardID = cardL.ID
	transactionL.TransactionDate = time.Now()
	transactionL.Amount = 29.9
	transactionL.Comment = "First Transaction"

	err = models.CreateTransaction(&transactionL)
	if err != nil {
		fmt.Printf(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

}

func DropData(c *gin.Context) {
	config.DB.Exec("DROP TABLE transactions")
	config.DB.Exec("DROP TABLE cards")
	config.DB.Exec("DROP TABLE accounts")
}