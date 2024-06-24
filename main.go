package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.formHandler)
	http.HandleFunc("/welcome", handlers.welcomeHandler)
	http.ListenAndServe(":8080", nil)
}
