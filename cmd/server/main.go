package main

import (
	"engbot/internal/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/health", handlers.HealthHandler)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Println("Starting server on " + port)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
