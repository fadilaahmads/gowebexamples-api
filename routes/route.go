package routes

import (
	"web-api/controller"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/create/table/user", controller.CreateUserTable)

	r.HandleFunc("/", controller.Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", controller.GetBooks).Methods("GET")
}
