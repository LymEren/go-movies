package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mux-master"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = (strconv.Itoa(rand.Intn(100000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// set json content type
// params
// loop over the movies, range
// delete the movie with the id that you have sent
// add a new movie (the movie that we send in the body of postman)

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// We have to define 5 different route and 5 different function for these
func main() {
	r := mux.NewRouter()

	// Added some movie for the test

	movies = append(movies, Movie{ID: "1", Isbn: "111222333", Title: "LOTR", Director: &Director{Firstname: "Peter", Lastname: "Jackson"}})
	movies = append(movies, Movie{ID: "2", Isbn: "111222334", Title: "Gladiator", Director: &Director{Firstname: "Unknown", Lastname: "Soldier"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies{id}", updateMovie).Methods("UPDATE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(" :8000", r))

}
