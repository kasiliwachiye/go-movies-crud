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

// Define movie struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Define director struct
type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Define movie slice
var movies []Movie

// Handler functions
// get all movies
func getMovies (w http.ResponseWriter, r *http.Request) {
  // Setting the content type as json
  w.Header().Set("Content-Type", "application/json")
  // Encode movies and send them
  json.NewEncoder(w).Encode(movies)
}

// get movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type as json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
  movieID := params["id"]
	for _, movie := range movies {
		if movie.ID == movieID {
			// Encode and send the found movie
			json.NewEncoder(w).Encode(movie)
			return // Exit the function after sending the response
		}
	}
	// If no movie is found, return a not found response
	http.NotFound(w, r)
}

// create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
  // Setting the content type as json
	w.Header().Set("Content-Type", "application/json")
  var movie Movie
  _ = json.NewDecoder(r.Body).Decode(&movie)
  movie.ID = strconv.Itoa((rand.Intn(100000000)))
  movies = append(movies, movie)
  json.NewEncoder(w).Encode(movie)
}

// update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type as json
	w.Header().Set("Content-Type", "application/json")

	// Get movie ID from URL parameters
	params := mux.Vars(r)
	movieID := params["id"]

	// Decode the request body to extract the updated movie information
	var updatedMovie Movie
	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Loop through the movies slice to find the movie with the specified ID
	for index, movie := range movies {
		if movie.ID == movieID {
			// Update the movie fields with the new values
			movies[index].Isbn = updatedMovie.Isbn
			movies[index].Title = updatedMovie.Title
			movies[index].Director = updatedMovie.Director

			// Encode and send the updated movie as the response
			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}

	// If no movie is found with the specified ID, return a not found response
	http.NotFound(w, r)
}

// delete movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type as json
	w.Header().Set("Content-Type", "application/json")

	// Get the ID of the movie to delete from the request parameters
	params := mux.Vars(r)
	movieID := params["id"]

	// Declare a variable to store the deleted movie
	var deletedMovie Movie

	// Loop through the movies slice to find the movie with the specified ID
	for index, movie := range movies {
		if movie.ID == movieID {
			// Store the deleted movie
			deletedMovie = movies[index]

			// Delete the movie from the slice by appending slices before and after the item
			movies = append(movies[:index], movies[index+1:]...)

			// Return the deleted movie in the response
			json.NewEncoder(w).Encode(deletedMovie)
			return
		}
	}

	// If no movie is found with the specified ID, return a not found response
	http.NotFound(w, r)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: strconv.Itoa((rand.Intn(100000000))), Isbn: "438227", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}},
	)
  movies = append(movies,
		Movie{ID: strconv.Itoa((rand.Intn(100000000))), Isbn: "45455", Title: "Avatar", Director: &Director{FirstName: "James", LastName: "Cameron"}},)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", r); err != nil {
		// Log any errors
		log.Fatal(err)
	}
}