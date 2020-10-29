package models

import (
	"fmt"
	"github.com/njilrem/go-rest-atm/config"
)

func GetCardById(card *Card, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Find(card).Error; err != nil {
		return err
	}
	return nil
}

func GetCardsByHolderId(card *[]Card, holderID string) (err error) {
	if err = config.DB.Where("holderID = ?", holderID).Find(card).Error; err != nil {
		return err
	}
	return nil
}

func CreateCard(card *Card) (err error) {
	if err = config.DB.Create(card).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCard(card *Card) (err error) {
	fmt.Println(card)
	config.DB.Save(card)
	return nil
}