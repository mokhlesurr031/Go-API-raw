package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"movie_review_apis/views"
	"net/http"
)

func main() {

	//http.HandleFunc("/", views.HomePage)
	//http.HandleFunc("/movies", views.MoviesAPI)
	//http.HandleFunc("/movies/delete", views.DeleteMovieAPI)
	//http.HandleFunc("/movies/update", views.UpdateMovieAPI)
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	log.Fatalln("ListenAndServe: ", err)
	//}

	r := chi.NewRouter()
	r.Get("/", views.HomePage)
	r.Get("/movies", views.MoviesAPI)
	r.Post("/movies", views.MoviesAPI)
	r.Get("/movies/details", views.DetailsMovieAPI)
	r.Delete("/movies/delete", views.DeleteMovieAPI)
	r.Put("/movies/update", views.UpdateMovieAPI)

	http.ListenAndServe(":3000", r)
}
