package main

import (
	"net/http"
	"web-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create New Router
	r := mux.NewRouter()

	r.HandleFunc("/", routes.Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", routes.GetBooks).Methods("POST")

	r.HandleFunc("/db/", routes.Database).Methods("GET")

	http.ListenAndServe(":80", r)
}
