package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"movie_review_apis/conn"
	"movie_review_apis/views"
	"net/http"
)

// Main ...
func main() {
	log.Println("Starting server....")
	r := chi.NewRouter()
	conn.Init()

	r.Get("/", views.HomePage)
	r.Get("/movies/get", views.GetMoviesAPI)
	//r.Post("/movies/add", views.CreateMovieAPI)
	r.Get("/movies/details", views.DetailsMovieAPI)
	r.Delete("/movies/delete", views.DeleteMovieAPI)
	r.Put("/movies/update", views.UpdateMovieAPI)

	http.ListenAndServe(":3000", r)
	conn.CloseDB()
}
