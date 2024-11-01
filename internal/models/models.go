package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName     string     `gorm:"not null"`
	LastName      string     `gorm:"not null"`
	Email         string     `gorm:"unique;not null"`
	Password      string     `gorm:"not null"`
	Role          string     `gorm:"not null;default: customer"`
	TokenVersion  uint       `gorm:"not null; default: 1"`
	ShoppingCarts []CartItem `gorm:"foreignKey:UserID"`
	Orders        []Order    `gorm:"foreignKey:UserID"`
	Products      []Product  `gorm:"foreignKey:UserID"`
}

type CartItem struct {
	gorm.Model

	UserID   User `gorm:"not null;index"`
	ProductID Product`gorm:"not null;index"`
	Quantity uint `gorm:"not null;check:quantity > 0"`
}

type OrderItem struct {
	gorm.Model

	OrderID   uint    `gorm:"not null;index"`
	Price     float64 `gorm:"not null;check:quantity >= 0"`
	ProductID uint    `gorm:"not null;index"`
	Quantity  uint    `gorm:"not null;check:quantity > 0"`

	Product Product `gorm:"foreignKey:ProductID;constraints:OnDelete:CASCADE"`
	Order   Order   `gorm:"foreignKey:OrderID;constraints:OnDelete:CASCADE"`
}

type Order struct {
	gorm.Model

	Total      uint        `gorm:"not null;default: 0"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	User       User        `gorm:"foreignKey:UserID;constraints:OnDelete:CASCADE"`
}

type Product struct {
	gorm.Model

	Name          string  `gorm:"not null;"`
	Description   string  `gorm:"not null"`
	Price         float64 `gorm:"not null;default: 0"`
	StockQuantity uint    `gorm:"not null;check:stock_quantity >= 0"`

	CartItems []CartItem `gorm:"foreignKey:ProductID;constraints:OnDelete:CASCADE"`
	OrderItems []OrderItem `gorm:"foreignkey:ProductID;constraints:OnDelete:CASCADE"`
}
