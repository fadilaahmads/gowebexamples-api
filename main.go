package main

import (
	"fmt"
	"net/http"
	"web-api/database"

	"github.com/gorilla/mux"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My Bookstore!")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Segmenting Request
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You requested book: %s on page %s\n", title, page)
}

func main() {
	// Create New Router
	r := mux.NewRouter()

	database.DB()

	r.HandleFunc("/", Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", GetBooks).Methods("POST")

	http.ListenAndServe(":80", r)
}
