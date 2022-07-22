package main

import (
	"net/http"
	"web-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create New Router
	r := mux.NewRouter()

	routes.Routes(r)

	http.ListenAndServe(":4444", r)
}
