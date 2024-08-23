package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Awesome Project d d !"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
}
