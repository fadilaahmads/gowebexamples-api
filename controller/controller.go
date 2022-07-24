package controller

import (
	"fmt"
	"net/http"
	"text/template"
	"web-api/repository"
	"web-api/services"

	"github.com/gorilla/mux"
)

func CreateUserTable(w http.ResponseWriter, r *http.Request) {
	repository.CreateUsersTable()
}

func DeleteSingleUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	repository.DeleteSingleUser()
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	repository.GetAllUser()
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	repository.InsertNewUser()
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

func Todo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/layout.html"))
	tmpl.Execute(w, services.TodoTemplate())
}
