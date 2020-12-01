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
	Name      string 	     `json:"name"  valid:"required"`
	Email     string 	     `json:"email" valid:"required,email"`
	Phone     string 	     `json:"phone" valid:"required,numeric"`
	Address   string 	     `json:"address" valid:"required"`
	Card      []Card 		 `gorm:"foreignKey:HolderID"`
}

// TableName db table name 
func (b *Account) TableName() string {
	return "accounts"
}