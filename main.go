package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	TITLE    string   `json:"title"`
	DIRECTOR DIRECTOR `json:"director"`
}

type DIRECTOR struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	fandf := Movie{
		ID:    "1",
		TITLE: "Fast and furious",
		DIRECTOR: DIRECTOR{
			Firstname: "ishan",
			Lastname:  "jain",
		},
	}
	movies = append(movies, fandf)

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("Starting web server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Returns the route variable for request...//In this case it was {id}
	// Only returns the variable in the route request and not the body of the request
	// For that we have used the method in createmovie
	vari := mux.Vars(r)

	for i, movie := range movies {
		if movie.ID == vari["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	varii := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == varii["id"] {
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))

	movies := append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var mov Movie
	vari := mux.Vars(r)
	for i, mtu := range movies {
		if mtu.ID == vari["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}

	}
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.ID = vari["id"]
	movies = append(movies, mov)
	json.NewEncoder(w).Encode(movies)
}
