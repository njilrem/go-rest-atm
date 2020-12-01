package tests

import (
	"github.com/njilrem/go-rest-atm/models"
	"testing"
)

func TestCreate(t *testing.T) {

	var account models.Account
	account.Name = "Test"
	account.Address = "Address"
	account.Phone = "380508565432"
	account.Email = "jin shon"
	err := models.CreateAccount(&account)
	if err != nil {
		t.Error(err)
	}
}