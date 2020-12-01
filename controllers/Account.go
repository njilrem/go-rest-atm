package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//GetAccounts ... Get all accounts
func GetAccounts(c *gin.Context) {
	var account []models.Account
	err := models.GetAllAccounts(&account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, account)
	}
 }
 //CreateAccount ... Create Account
 func CreateAccount(c *gin.Context) {
	var account models.Account
	c.BindJSON(&account)
	err := models.CreateAccount(&account)
	if err != nil {
	 log.Warn(err.Error())
	 c.AbortWithStatus(http.StatusInternalServerError)
	} else {
	 c.JSON(http.StatusOK, account)
	}
 }
 //GetAccountByID ... Get the account by id
 func GetAccountByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var account models.Account
	err := models.GetAccountByID(&account, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{
		 "name": account.Name,
		 "surname": account.Email,
		 "cardNum": account.Card[0].CardNum,
		 "expDate": account.Card[0].ExpireDate,
		 "balance": account.Card[0].Balance,
	 })
	}
 }
 //UpdateAccount ... Update the account information
 func UpdateAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id") // if0wjfwef.com/accounts-api/account/5453 { "name": "ddd"}
	err := models.GetAccountByID(&account, id)
	if err != nil {
	 c.JSON(http.StatusNotFound, account)
	}
	c.BindJSON(&account)
	err = models.UpdateAccount(&account, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, account)
	}
 }
 //DeleteAccount ... Delete the account
 func DeleteAccount(c *gin.Context) {
	var account models.Account
	id := c.Params.ByName("id")
	err := models.DeleteAccount(&account, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{"result": id + " is deleted"})
	}
 }

