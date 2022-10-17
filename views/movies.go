package views

import (
	"encoding/json"
	"fmt"
	"movie_review_apis/conn"
	"movie_review_apis/mgs"
	"net/http"
)

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

type PutMovie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}

func MoviesAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heyyy Mahin")
	w.Header().Add("content-type", "application/json")
	db := conn.SetupDB()
	fmt.Println("Method: ", r.Method)

	if r.Method == "GET" {
		mgs.PrintMessage("Getting Movie List... ")
		var movies []Movie
		db.Find(&movies)
		var response = JsonResponse{Type: "success", Data: movies}
		json.NewEncoder(w).Encode(response)
	}

	//else if r.Method == "POST" {
	//	var mv Movie
	//
	//	err := json.NewDecoder(r.Body).Decode(&mv)
	//	if err != nil {
	//		mgs.CheckErr(err)
	//	} else {
	//		var response JsonResponse
	//
	//		if mv.Name == "" || mv.Year == "" {
	//			response = JsonResponse{Type: "error", Message: "movieID or movieName missing"}
	//		} else {
	//			mgs.PrintMessage("Inserting Movie into DB")
	//
	//			var lastInsertID int
	//
	//			err := db.QueryRow("INSERT INTO movies(Name, Year) VALUES ($1, $2) returning id;", mv.Name, mv.Year).Scan(&lastInsertID)
	//
	//			mgs.CheckErr(err)
	//			response = JsonResponse{Type: "success", Message: "Movie has been inserted"}
	//
	//			json.NewEncoder(w).Encode(response)
	//		}
	//
	//	}
	//}
}

func DetailsMovieAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Movie Details Data")
	if r.Method == "GET" {
		db := conn.SetupDB()
		id := r.URL.Query().Get("id")
		var movie []Movie
		db.Find(&movie, "id=?", id)
		var response = JsonResponse{Type: "success", Data: movie}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteMovieAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")

		var response JsonResponse
		if id == "" {
			response = JsonResponse{Type: "error", Message: "You are missing something"}
		} else {
			db := conn.SetupDB()
			var movie Movie
			db.Where("id=?", id).Delete(&movie)
			response = JsonResponse{Type: "success", Message: "Movie Deleted"}
		}
		json.NewEncoder(w).Encode(response)
	}

}

func UpdateMovieAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update triggered")
	w.Header().Add("content-type", "application/json")
	if r.Method == "PUT" {
		db := conn.SetupDB()
		var movie Movie
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
}
