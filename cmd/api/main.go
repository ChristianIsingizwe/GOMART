package main

import (
	"log"

	"github.com/ChristianIsingizwe/GOMART/internal/database"
)

func main() {

	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database.")
	}
}
