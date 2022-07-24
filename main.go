package main

import (
	"log"
	"net/http"
	"web-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create New Router
	r := mux.NewRouter()

	routes.Routes(r)

	err := http.ListenAndServe(":4444", r)
	if err != nil {
		log.Fatal(err)
	}
}
