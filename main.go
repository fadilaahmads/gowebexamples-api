package main

import (
	"net/http"
	"web-api/database"
	"web-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create New Router
	r := mux.NewRouter()

	database.DB()

	r.HandleFunc("/", routes.Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", routes.GetBooks).Methods("POST")

	http.ListenAndServe(":80", r)
}
