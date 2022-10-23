package main

import (
	_ "github.com/lib/pq"
	"movie_review_apis/cmd"
)

// Main ...
func main() {
	//log.Println("Starting server....")
	//r := chi.NewRouter()
	//conn.Init()
	//
	//r.Get("/", views.HomePage)
	//r.Get("/movies/get", views.GetMoviesAPI)
	//r.Post("/movies/add", views.CreateMovieAPI)
	//r.Get("/movies/details", views.DetailsMovieAPI)
	//r.Delete("/movies/delete", views.DeleteMovieAPI)
	//r.Put("/movies/update", views.UpdateMovieAPI)
	//
	//http.ListenAndServe(":3000", r)
	//conn.CloseDB()
	cmd.Execute()
}
