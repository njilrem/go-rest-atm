package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CardID    uint `json:"card_id" valid:"required"`
	CardNum   string `json:"card_num" valid:"required, creditcard"`
	Amount    float64 `json:"amount" valid:"required"`
	TransactionDate time.Time `json:"transaction_date" valid:"required"`
	Comment   string `json:"comment" valid:"required"`
}

func (b* Transaction) TableName() string {
	return "transactions"
}