package models

import (
	"github.com/njilrem/go-rest-atm/config"
	"fmt"
  //only need init() function
	_ "github.com/go-sql-driver/mysql"
)

//GetAllAccounts Fetch all accounts data
func GetAllAccounts(account *[]Account) (err error) {
	if err = config.DB.Find(account).Error; err != nil {
		return err
	}
	return nil
}

//CreateAccount ... Insert New data
func CreateAccount(account *Account) (err error) {
	if err = config.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}

//GetAccountByID ... Fetch only one account by Id
func GetAccountByID(account *Account, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(account).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAccount ... Update account
func UpdateAccount(account *Account, id string) (err error) {
	fmt.Println(account)
	config.DB.Save(account)
	return nil
}

//DeleteAccount ... Delete account
func DeleteAccount(account *Account, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(account)
	return nil
}