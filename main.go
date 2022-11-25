package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `jason:"isbn"`
	Title    string    `jason:"title"`
	Director *Director `jason:"director"`
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
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "Jhone", Lastname: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "sebin", Lastname: "joe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("String server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
