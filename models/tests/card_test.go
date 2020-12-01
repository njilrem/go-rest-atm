package tests

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/njilrem/go-rest-atm/models"
	"testing"
)

func TestValidCreateCard(t *testing.T) {
	var card models.Card
	card.CardNum = "5347816310496245"
	card.ExpireDate = "8/24/2029"
	card.Pin = "4444"
	card.Cvv2 = "222"
	card.Balance = 0
	card.HolderID = 1
	_, err := valid.ValidateStruct(card)
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidCardNum(t *testing.T) {
	var card models.Card
	card.CardNum = "5347816310496245222"
	card.ExpireDate = "8/24/2029"
	card.Pin = "4444"
	card.Cvv2 = "222"
	card.Balance = 0
	card.HolderID = 1
	_, err := valid.ValidateStruct(card)
	if err != nil {
		t.Logf("%v", err)
	}
}

func TestMissingFieldHolderIDCard(t *testing.T) {
	var card models.Card
	card.CardNum = "5347816310496245"
	card.Pin = "4444"
	card.Cvv2 = "222"
	card.Balance = 0
	_, err := valid.ValidateStruct(card)
	if err != nil {
		t.Logf("%v", err)
	}
}

func TestMissingFieldsCard(t *testing.T) {
	var card models.Card
	card.CardNum = "5347816310496245"
	card.ExpireDate = "8/24/2029"
	_, err := valid.ValidateStruct(card)
	if err != nil {
		t.Logf("%v", err)
	}
}

