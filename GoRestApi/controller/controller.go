package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paulgureghian/Go_Programs/GoRestApi/db"
	"github.com/paulgureghian/Go_Programs/GoRestApi/models"
	"log"
	"net/http"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var movie models.Movie
	err := decoder.Decode(&movie)

	if err != nil {
		log.Println(err.Error())
	}

	db.MovieList = append(db.MovieList, movie)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	allMovies := db.MovieList
	json.NewEncoder(w).Encode(allMovies)

}

func GetMovieById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := mux.Vars(r)["id"]

	for _, movie := range db.MovieList {

		if id == movie.Id {

			json.NewEncoder(w).Encode(movie)
			return

		}
	}

	message := models.Message{Message: "Movie not found"}
	json.NewEncoder(w).Encode(message)

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := mux.Vars(r)["id"]

	for index, movie := range db.MovieList {

		if id == movie.Id {

			db.MovieList = append(db.MovieList[:index], db.MovieList[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(db.MovieList)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var movie models.Movie

		json.NewDecoder(r.Body).Decode(&movie)

		for index, movieToUpdate := range db.MovieList {

			if movieToUpdate.Id == movie.Id {

				db.MovieList[index] = movie
			}
		}
}
