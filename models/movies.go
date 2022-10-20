package models

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
}

type PutMovie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}
