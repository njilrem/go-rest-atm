package tests

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/njilrem/go-rest-atm/models"
	"testing"
)

func TestValidCreateAccount(t *testing.T) {
	var account models.Account
	account.Name = "Test"
	account.Address = "Address"
	account.Phone = "380508565432"
	account.Email = "jinJon@gmail.com"
	result, err := valid.ValidateStruct(account)
	if err != nil {
		t.Error(err)
	}
	if !result {
		t.Error(err)
	}
}

func TestInvalidEmailCreateAccount(t *testing.T) {
	var account models.Account
	account.Name = "Jing"
	account.Email = "aaaaaaa"
	account.Phone = "380508565432"
	account.Address = "STREET"
	_, err := valid.ValidateStruct(account)
	if err != nil {
		t.Logf("%v", err)
	}
}

func TestMissingRequiredFieldCreateAccount(t *testing.T) {
	var account models.Account
	account.Name = "Jing"
	account.Email = "aaaaaaa"
	account.Phone = "380508565432"
	_, err := valid.ValidateStruct(account)
	if err != nil {
		t.Logf("%v", err)
	}
}
