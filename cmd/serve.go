package cmd

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"movie_review_apis/conn"
	"movie_review_apis/views"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func StartServer() {
	port := viper.GetInt("server.port")

	log.Println("Starting server....")
	r := chi.NewRouter()
	conn.Init()

	r.Get("/", views.HomePage)
	r.Get("/movies/get", views.GetMoviesAPI)
	r.Post("/movies/add", views.CreateMovieAPI)
	r.Get("/movies/details", views.DetailsMovieAPI)
	r.Delete("/movies/delete", views.DeleteMovieAPI)
	r.Put("/movies/update", views.UpdateMovieAPI)

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	conn.CloseDB()
}
