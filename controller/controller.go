package controller

import (
	"fmt"
	"net/http"
	"web-api/repository"

	"github.com/gorilla/mux"
)

func CreateUserTable(w http.ResponseWriter, r *http.Request) {
	repository.CreateUsersTable()
}

func DeleteSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	repository.DeleteSingleUser(id)
}

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
