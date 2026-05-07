package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {}

func getMovie(w http.ResponseWriter, r *http.Request) {}

func createMovie(w http.ResponseWriter, r *http.Request) {}

func updateMove(w http.ResponseWriter, r *http.Request) {}

func deleteMovie(w http.ResponseWriter, r *http.Request) {}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Micheal", Director: &Director{FirstName: "Micheal", LastName: "Jackson"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23456", Title: "Swiftie", Director: &Director{FirstName: "Taylor", LastName: "Swift"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMove).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("8000", r))
}
