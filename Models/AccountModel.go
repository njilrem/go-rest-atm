package models

// Account Model
type Account struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

// TableName db table name 
func (b *Account) TableName() string {
	return "account"
}