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
	CardNum    string 		  `json:"cardNum"`
	ExpireDate string 		  `json:"expireDate"`
	HolderID   uint   		  `json:"holderID"`
}

// TableName db table name
func (b *Card) TableName() string {
	return "cards"
}