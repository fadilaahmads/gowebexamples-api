package routes

import (
	"net/http"
	"web-api/controller"
	"web-api/services"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/forms/", controller.Forms)
	r.HandleFunc("/table/create/user", controller.CreateUserTable)
	r.HandleFunc("/user/delete/{id}", controller.DeleteSingleUser)
	r.HandleFunc("/user/add/", controller.InsertNewUser)
	r.HandleFunc("/user/get/all", controller.GetAllUser)
	r.Handle("/static/", http.StripPrefix("/static/", *services.StaticFiles()))
	r.HandleFunc("/todo/", controller.Todo)
	r.HandleFunc("/", controller.Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", controller.GetBooks).Methods("GET")
}
