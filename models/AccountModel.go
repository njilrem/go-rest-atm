package models

import (
	"gorm.io/gorm"
	"time"
)

// Account Model
type Account struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Card []Card `gorm:"foreignKey:HolderID"`
}

// TableName db table name 
func (b *Account) TableName() string {
	return "accounts"
}