package models

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"github.com/njilrem/go-rest-atm/config"
	"log"
)

func GetTransactionById(transaction *Transaction, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Find(transaction).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetTransactionsByCardId(transactions *[]Transaction, id string) (err error) {
	if err = config.DB.Where("card_id = ?", id).Find(transactions).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTransaction(transaction *Transaction) (err error) {
	var result bool
	result, err = valid.ValidateStruct(transaction)
	if result {
		if err = config.DB.Create(transaction).Error; err != nil {
			log.Println(err)
			return err
		}
		return nil
	} else {
		return errors.New("validation error")
	}
}

func DeleteTransaction(transaction *Transaction) (err error) {
	config.DB.Where("id = ?", transaction.ID).Delete(transaction)
	return nil
}