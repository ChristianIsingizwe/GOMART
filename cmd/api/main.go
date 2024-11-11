package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ChristianIsingizwe/GOMART/internal/database"
	"github.com/ChristianIsingizwe/GOMART/internal/handlers"
)

func main() {

	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database.")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginUser)
	mux.HandleFunc("/register", handlers.RegisterUser)


	server := &http.Server{
		Addr: os.Getenv("APP_PORT"),
		Handler: mux,
	}

	fmt.Println("Starting server on port: " + os.Getenv("APP_PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start the server: %s\n", err)
	}
}
