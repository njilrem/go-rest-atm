package tests

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/njilrem/go-rest-atm/models"
	"testing"
	"time"
)

func TestValidCreateTransaction(t *testing.T) {
	var transaction models.Transaction
	transaction.CardNum = "5395171188879325"
	transaction.Amount = 1000
	transaction.CardID = 3
	transaction.TransactionDate = time.Now()
	transaction.Comment = "Transfer"
	result, err := valid.ValidateStruct(transaction)
	if err != nil {
		t.Error(err)
	}
	if !result {
		t.Error(err)
	}
}

func TestInvalidCardNumTransaction(t *testing.T) {
	var transaction models.Transaction
	transaction.CardNum = "5395171188879325444"
	transaction.Amount = 1000
	transaction.CardID = 3
	transaction.TransactionDate = time.Now()
	transaction.Comment = "Transfer"
	_, err := valid.ValidateStruct(transaction)
	if err != nil {
		t.Log(err)
	}
}

func TestInvalidAmountTransaction(t *testing.T) {
	var transaction models.Transaction
	transaction.CardNum = "5395171188879325"
	transaction.Amount = -100
	transaction.CardID = 3
	transaction.TransactionDate = time.Now()
	transaction.Comment = "Transfer"
	_, err := valid.ValidateStruct(transaction)
	if err != nil {
		t.Log(err)
	}
}

func TestInvalidMissingDateTransaction(t *testing.T) {
	var transaction models.Transaction
	transaction.CardNum = "5395171188879325"
	transaction.Amount = 1000
	transaction.CardID = 3
	transaction.Comment = "Transfer"
	_, err := valid.ValidateStruct(transaction)
	if err != nil {
		t.Log(err)
	}
}

func TestInvalidCardIDTransaction(t *testing.T) {
	var transaction models.Transaction
	transaction.CardNum = "5395171188879325"
	transaction.Amount = 1000
	transaction.TransactionDate = time.Now()
	transaction.Comment = "Transfer"
	_, err := valid.ValidateStruct(transaction)
	if err != nil {
		t.Log(err)
	}
}
