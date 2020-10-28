package models

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	cvv2 string
	cardNum string
	expireDate string
}

// TableName db table name
func (b *Card) TableName() string {
	return "cards"
}