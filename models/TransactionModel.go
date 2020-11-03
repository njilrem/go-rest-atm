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
	CardID    uint `json:"card_id"`
	CardNum   string `json:"card_num"`
	Amount    float64 `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}

func (b* Transaction) TableName() string {
	return "transactions"
}