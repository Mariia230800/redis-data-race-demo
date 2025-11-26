package models

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
