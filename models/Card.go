package models

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"github.com/njilrem/go-rest-atm/config"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetCardById(card *Card, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Find(card).Error; err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func GetCardsByHolderId(card *[]Card, holderID string) (err error) {
	if err = config.DB.Where("holder_id = ?", holderID).Find(card).Error; err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func CreateCard(card *Card) (err error) {
	var result bool
	result, err = valid.ValidateStruct(card)
	if result {
		hash, err := bcrypt.GenerateFromPassword([]byte(card.Cvv2), bcrypt.DefaultCost)
		if err != nil {
			log.Warn(err)
			return err
		}
		card.Cvv2 = string(hash)
		hash, err = bcrypt.GenerateFromPassword([]byte(card.Pin), bcrypt.DefaultCost)
		if err != nil {
			log.Warn(err)
			return err
		}
		card.Pin = string(hash)
		if err = config.DB.Create(card).Error; err != nil {
			log.Warn(err)
			return err
		}
		return nil
	} else {
		return errors.New("validation error")
	}
}

func UpdateCard(card *Card) (err error) {
	if len(card.Cvv2) == 3 {
		hash, err := bcrypt.GenerateFromPassword([]byte(card.Cvv2), bcrypt.DefaultCost)
		if err != nil {
			log.Warn(err)
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
		log.Warn(err)
		return err
	}
	if transaction.Amount > card.Balance {
		log.Info("Insufficient funds!")
		return errors.New("")
	}
	card.Balance -= transaction.Amount

	var receiverCard Card
	if err = config.DB.Where("card_num = ?", transaction.CardNum).Find(&receiverCard).Error; err != nil {
		/* IMAGINE PROCESSING A BANK OPERATION OF TRANSFERRING MONEY TO the other bank OVER HERE*/
		log.Info("Transaction is incomplete, receiver's card belongs to other bank")
	}
	if len(receiverCard.Cvv2) > 3 {
		receiverCard.Balance += transaction.Amount
		var transactionReceiver Transaction
		transactionReceiver.Amount = transaction.Amount
		transactionReceiver.Comment = "Receive Transfer"
		transactionReceiver.TransactionDate = time.Now()
		transactionReceiver.CardID = receiverCard.ID
		transactionReceiver.CardNum = receiverCard.CardNum
		err = CreateTransaction(&transactionReceiver); if err != nil {
			log.Warn("Transaction creation failed")
		}
	}


	config.DB.Save(card)
	return nil
}

func ProcessRefillTransaction(transaction Transaction) (err error) {
	var refillingCard Card
	if err = config.DB.Where("id = ?", transaction.CardID).Find(&refillingCard).Error; err != nil {
		log.Warn(err)
		return err
	}
	refillingCard.Balance += transaction.Amount
	config.DB.Save(refillingCard)
	return nil
}