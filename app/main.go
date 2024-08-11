package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Back")
	})

	fmt.Println("Starting server on :80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
