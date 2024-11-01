package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	TokenVersion uint   `gorm:"not null; default: 1"`
}

type CartItem struct {
	gorm.Model

	Quantity uint `gorm:"not null; default: 0"`
}

type OrderItem struct {
	gorm.Model

	Price uint `gorm:"not null;default: 0"`
	Quantity uint `gorm:"not null;default:0"`
}

type Order struct {
	gorm.Model

	
}
