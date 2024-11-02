package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ChristianIsingizwe/GOMART/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectToDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbName=%s ",
		os.Getenv("DB_HOST"),os.Getenv("PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.CartItem{}, &models.OrderItem{}); err != nil{
		log.Fatalf("Auto migration failed:  %v", err)
	}

	DB= database
	return nil
}