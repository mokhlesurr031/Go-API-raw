package views

import (
	"encoding/json"
	"movie_review_apis/mgs"
	"movie_review_apis/models"
	"movie_review_apis/querydir"
	"net/http"
	"strconv"
)

type JsonResponseSingle struct {
	Type    string       `json:"type"`
	Data    models.Movie `json:"data"`
	Message string       `json:"message"`
}

type JsonResponseMultiple struct {
	Type    string         `json:"type"`
	Data    []models.Movie `json:"data"`
	Message string         `json:"message"`
}

// GetMoviesAPI ...
func GetMoviesAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	movies, err := querydir.MovieGetQuery()
	if err == nil {
		var response = JsonResponseMultiple{Type: "success", Data: movies}
		json.NewEncoder(w).Encode(response)
	} else {
		mgs.CheckErr(err)
	}

}

// CreateMovieAPI ...
//func CreateMovieAPI(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("content-type", "application/json")
//	db := conn.GetDB()
//	var mv models.Movie
//	err := json.NewDecoder(r.Body).Decode(&mv)
//	if err != nil {
//		mgs.CheckErr(err)
//	} else {
//		var response JsonResponseSingle
//		if mv.Name == "" || mv.Year == "" {
//			response = JsonResponseSingle{Type: "error", Message: "movieID or movieName missing"}
//		} else {
//			movie := models.Movie{Name: mv.Name, Year: mv.Year}
//			db.Create(&movie)
//			response = JsonResponseSingle{Type: "success", Message: "Movie has been inserted"}
//			json.NewEncoder(w).Encode(response)
//		}
//
//	}
//}

// DetailsMovieAPI ...
func DetailsMovieAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	movie, err := querydir.MovieDetailQuery(idInt)
	if err != nil {
		res := JsonResponseSingle{Type: "error", Data: models.Movie{}}
		json.NewEncoder(w).Encode(res)
	} else {
		var response = JsonResponseSingle{Type: "success", Data: *movie}
		json.NewEncoder(w).Encode(response)
	}

}

// DeleteMovieAPI ...
func DeleteMovieAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var response JsonResponseSingle

	if id == "" {
		response = JsonResponseSingle{Type: "error", Message: "You are missing something"}
	} else {
		idInt, _ := strconv.Atoi(id)
		status, err := querydir.MovieDeleteQuery(idInt)
		if err != nil {
			mgs.CheckErr(err)
		}
		response = JsonResponseSingle{Type: "success", Message: strconv.Itoa(status)}
	}
	json.NewEncoder(w).Encode(response)

}

type UpdateMovieRequest struct {
	FirstName string
	LastName  string
	Email     string
}

// UpdateMovieAPI ...
func UpdateMovieAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	var pm models.Movie
	err := json.NewDecoder(r.Body).Decode(&pm)
	if err != nil {
		mgs.CheckErr(err)
	} else {
		querydir.MovieUpdateQuery(idInt, &pm)

		//response = JsonResponseSingle{Type: "success", Message: res}
	}
	//json.NewEncoder(w).Encode(response)
}

/*

type User struct{
  ID int
  Name string
  Email string
}


{
   "first_name": "Mahin",
   "last_name": "Dev",
   "email": "mahin@dev.com"
}
*/
