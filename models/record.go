package models

type Record struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
