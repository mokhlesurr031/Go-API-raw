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
	w.Header().Add("content-type", "application/json")
	db := conn.SetupDB()
	fmt.Println("Method: ", r.Method)

	if r.Method == "GET" {
		mgs.PrintMessage("Getting Movie List... ")
		rows, err := db.Query("SELECT * FROM movies")
		mgs.CheckErr(err)
		fmt.Println("Moviesssss", rows)

		var movies []Movie

		for rows.Next() {
			var id int
			var MovieName string
			var Year string
			err = rows.Scan(&id, &MovieName, &Year)
			mgs.CheckErr(err)
			movies = append(movies, Movie{MovieName, Year})
		}
		var response = JsonResponse{Type: "success", Data: movies}
		json.NewEncoder(w).Encode(response)
	} else if r.Method == "POST" {
		var mv Movie

		err := json.NewDecoder(r.Body).Decode(&mv)
		if err != nil {
			mgs.CheckErr(err)
		} else {
			var response JsonResponse

			if mv.Name == "" || mv.Year == "" {
				response = JsonResponse{Type: "error", Message: "movieID or movieName missing"}
			} else {
				mgs.PrintMessage("Inserting Movie into DB")

				var lastInsertID int

				err := db.QueryRow("INSERT INTO movies(Name, Year) VALUES ($1, $2) returning id;", mv.Name, mv.Year).Scan(&lastInsertID)

				mgs.CheckErr(err)
				response = JsonResponse{Type: "success", Message: "Movie has been inserted"}

				json.NewEncoder(w).Encode(response)
			}

		}
	}
}

func DetailsMovieAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Movie Details Data")
	if r.Method == "GET" {
		db := conn.SetupDB()
		id := r.URL.Query().Get("id")

		row, err := db.Query("SELECT * FROM movies WHERE id=$1", id)
		mgs.CheckErr(err)
		fmt.Println("Moviesssss", row)

		var movies []Movie

		for row.Next() {
			var id int
			var MovieName string
			var Year string
			err = row.Scan(&id, &MovieName, &Year)
			mgs.CheckErr(err)
			movies = append(movies, Movie{MovieName, Year})
		}
		var response = JsonResponse{Type: "success", Data: movies}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteMovieAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete triggered")
	if r.Method == "DELETE" {
		//fmt.Fprintf(w, "Delete Triggered")
		id := r.URL.Query().Get("id")

		var response JsonResponse

		if id == "" {
			response = JsonResponse{Type: "error", Message: "You are missing something"}

		} else {
			db := conn.SetupDB()
			mgs.PrintMessage("Deleting movie from DB")

			_, err := db.Exec("DELETE FROM movies WHERE id=$1", id)
			mgs.CheckErr(err)

			response = JsonResponse{Type: "success", Message: "Movie Deleted"}

		}

		json.NewEncoder(w).Encode(response)
	}

}

func UpdateMovieAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update triggered")
	w.Header().Add("content-type", "application/json")
	db := conn.SetupDB()

	if r.Method == "PUT" {
		id := r.URL.Query().Get("id")
		var response JsonResponse
		fmt.Println(id)
		var pm PutMovie

		err := json.NewDecoder(r.Body).Decode(&pm)
		if err != nil {
			mgs.CheckErr(err)
		} else {
			fmt.Println("PUT DATA", pm.ID, pm.Year, pm.Name)
			mgs.PrintMessage("Deleting movie from DB")
			_, err := db.Exec("UPDATE movies set Name=$1, Year=$2  WHERE id=$3", pm.Name, pm.Year, id)
			mgs.CheckErr(err)

			response = JsonResponse{Type: "success", Message: "Movie Deleted"}

		}
		json.NewEncoder(w).Encode(response)

	}
}
