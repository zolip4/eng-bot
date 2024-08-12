package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// API route
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")

	// Serve static files from the React app
	fs := http.FileServer(http.Dir("./resources/build"))
	router.PathPrefix("/").Handler(fs)

	// Start the server on port 8080
	http.ListenAndServe(":80", router)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go gfgfg"))
}
