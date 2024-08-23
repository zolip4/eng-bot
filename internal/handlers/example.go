package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Awesome!"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "OK"}
	json.NewEncoder(w).Encode(response)
}
