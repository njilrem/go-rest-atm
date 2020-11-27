package models

import (
	"errors"
	"fmt"
	"github.com/njilrem/go-rest-atm/config"
	"github.com/njilrem/go-rest-atm/dto"

	//only need init() function
	_ "github.com/go-sql-driver/mysql"
)

//GetAllAccounts Fetch all accounts data
func GetAllAccounts(account *[]Account) (err error) {
	if err = config.DB.Joins("JOIN cards on cards.holder_id=accounts.id").
		Preload("Card").
		Find(account).Error; err != nil {
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
	if err = config.DB.Where("accounts.id = ?", id).
		Joins("JOIN cards on cards.holder_id=accounts.id").
		Preload("Card").
		First(account).Error; err != nil {
			return err
	}
	return nil
}

func GetAccountByCredentials(credentials dto.AuthCredentials, account *Account) (err error) {
	if err = config.DB.Where("cards.card_num = ? AND cards.expire_date = ?", credentials.CardNum, credentials.ExpireDate).
		Joins("JOIN cards on cards.holder_id=accounts.id").
		Preload("Card").
		Find(account).
		Error; err != nil {
			return err
	}
	return nil
}

func GetAccountByCardNum(credentials dto.AuthCredentialsPIN, account *Account) (err error) {
	if err = config.DB.Where("cards.card_num = ?", credentials.CardNum).
		Joins("JOIN cards on cards.holder_id=accounts.id").
		Preload("Card").
		Find(account).
		Error; err != nil {
			return err
	}
	return nil
}

func GetAuthAccount(credentials dto.AuthAdminCredentials, account *Account) (err error) {
	if len(credentials.Phone) != 15 && len(credentials.Name) != 10{
		return errors.New("bad credentials")
	}
	if err = config.DB.Where("name = ? AND phone = ?", credentials.Name, credentials.Phone).
		Find(account).Error; err != nil {
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