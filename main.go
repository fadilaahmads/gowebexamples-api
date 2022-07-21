package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create New Router
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to My Bookstore!")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// Segmenting Request
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You requested book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}
