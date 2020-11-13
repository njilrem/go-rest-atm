package models

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID         uint           `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Balance    float64		  `json:"balance"`
	Cvv2       string 		  `json:"cvv2"`
	Pin        string		  `json:"pin"`
	CardNum    string 		  `json:"cardNum"`
	ExpireDate string 		  `json:"expireDate"`
	HolderID   uint   		  `json:"holderID"`
	Transaction []Transaction `gorm:"foreignKey:card_id"`
}

// TableName db table name
func (b *Card) TableName() string {
	return "cards"
}