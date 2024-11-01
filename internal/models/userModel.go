package models

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	TokenVersion uint `gorm:"not null; default: 1"`
	CreatedAt time.Time `json:"autoCreateTime"`
	UpdatedAt time.Time`json:"autoUpdateTime"`
}