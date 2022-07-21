package main

import (
	"database/sql"
	"fmt"
	"net/http"

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

	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

	// Initialize the first connection to the database, to see if everything works correctly.
	err := db.Ping()

	r.HandleFunc("/", Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", GetBooks).Methods("POST")

	http.ListenAndServe(":80", r)
}
