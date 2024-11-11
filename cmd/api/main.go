package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ChristianIsingizwe/GOMART/internal/database"
	"github.com/ChristianIsingizwe/GOMART/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	err = database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database.")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginUser)
	mux.HandleFunc("/register", handlers.RegisterUser)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port =":8080"
	}

	server := &http.Server{
		Addr: port,
		Handler: mux,
	}

	fmt.Println("Starting server on port: " + port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start the server: %s\n", err)
	}
}
