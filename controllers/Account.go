package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/njilrem/go-rest-atm/models"
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
	 fmt.Println(err.Error())
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
	 c.JSON(http.StatusOK, account)
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

