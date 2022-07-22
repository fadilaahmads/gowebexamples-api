package routes

import (
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
