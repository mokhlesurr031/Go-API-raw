package models

type Movie struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Year string `json:"year"`
}

type PutMovie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}
