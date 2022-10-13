package main

import (
	_ "github.com/lib/pq"
	"log"
	"movie_review_apis/views"
	"net/http"
)

func main() {
	http.HandleFunc("/", views.HomePage)
	http.HandleFunc("/movies", views.MoviesAPI)
	http.HandleFunc("/movies/delete", views.DeleteMovieAPI)
	http.HandleFunc("/movies/update", views.UpdateMovieAPI)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
