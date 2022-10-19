package views

import (
	"encoding/json"
	"fmt"
	"movie_review_apis/conn"
	"movie_review_apis/mgs"
	"movie_review_apis/models"
	"net/http"
)

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    []models.Movie `json:"data"`
	Message string         `json:"message"`
}

type PutMovie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}

// GetMoviesAPI ...
func GetMoviesAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var movies []models.Movie
	db := conn.GetDB()
	db.Find(&movies)
	fmt.Println(movies)
	var response = JsonResponse{Type: "success", Data: movies}
	json.NewEncoder(w).Encode(response)
}

// CreateMovieAPI ...
func CreateMovieAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	db := conn.GetDB()
	var mv models.Movie
	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		mgs.CheckErr(err)
	} else {
		var response JsonResponse
		if mv.Name == "" || mv.Year == "" {
			response = JsonResponse{Type: "error", Message: "movieID or movieName missing"}
		} else {
			movie := models.Movie{Name: mv.Name, Year: mv.Year}
			db.Create(&movie)
			response = JsonResponse{Type: "success", Message: "Movie has been inserted"}
			json.NewEncoder(w).Encode(response)
		}

	}
}

// DetailsMovieAPI ...
func DetailsMovieAPI(w http.ResponseWriter, r *http.Request) {
	db := conn.GetDB()
	id := r.URL.Query().Get("id")
	var movie []models.Movie
	db.Find(&movie, "id=?", id)
	var response = JsonResponse{Type: "success", Data: movie}
	json.NewEncoder(w).Encode(response)
}

// DeleteMovieAPI ...
func DeleteMovieAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var response JsonResponse
	if id == "" {
		response = JsonResponse{Type: "error", Message: "You are missing something"}
	} else {
		db := conn.GetDB()
		var movie models.Movie
		db.Where("id=?", id).Delete(&movie)
		response = JsonResponse{Type: "success", Message: "Movie Deleted"}
	}
	json.NewEncoder(w).Encode(response)

}

// UpdateMovieAPI ...
func UpdateMovieAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	db := conn.GetDB()
	var movie models.Movie
	var pm PutMovie
	var response JsonResponse
	err := json.NewDecoder(r.Body).Decode(&pm)
	if err != nil {
		mgs.CheckErr(err)
	} else {
		id := r.URL.Query().Get("id")
		db.Find(&movie, "id=?", id)
		movie.Name = pm.Name
		movie.Year = pm.Year
		db.Where("id=?", id).Save(&movie)
		response = JsonResponse{Type: "success", Message: "Movie Updated"}
	}
	json.NewEncoder(w).Encode(response)
}
