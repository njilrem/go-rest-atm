package models

import (
	"fmt"
	"github.com/njilrem/go-rest-atm/config"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetCardById(card *Card, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Find(card).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetCardsByHolderId(card *[]Card, holderID string) (err error) {
	if err = config.DB.Where("holder_id = ?", holderID).Find(card).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateCard(card *Card) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(card.Cvv2), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	card.Cvv2 = string(hash)
	hash, err = bcrypt.GenerateFromPassword([]byte(card.Pin), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	card.Pin = string(hash)
	if err = config.DB.Create(card).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateCard(card *Card) (err error) {
	fmt.Println(card)
	if len(card.Cvv2) == 3 {
		hash, err := bcrypt.GenerateFromPassword([]byte(card.Cvv2), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return err
		}
		card.Cvv2 = string(hash)
	}
	if len(card.Pin) == 4 {
		hash, err := bcrypt.GenerateFromPassword([]byte(card.Pin), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return err
		}
		card.Pin = string(hash)
	}
	config.DB.Save(card)
	return nil
}

func ProcessTransaction(transaction Transaction) (err error) {
	var card Card
	if err = config.DB.Where("id = ?", transaction.CardID).Find(&card).Error; err != nil {
		log.Println(err)
		return err
	}
	card.Balance -= transaction.Amount
	/* IMAGINE PROCESSING A BANK OPERATION OF TRANSFERRING MONEY OVER HERE TSSSSS */
	config.DB.Save(card)
	return nil
}