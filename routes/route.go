package routes

import (
	"web-api/controller"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/table/create/user", controller.CreateUserTable)
	r.HandleFunc("/user/delete/{id}", controller.DeleteSingleUser)
	r.HandleFunc("/user/add/", controller.InsertNewUser)
	r.HandleFunc("/user/get/all", controller.GetAllUser)
	r.HandleFunc("/", controller.Homepage).Methods("GET")

	r.HandleFunc("/books/{title}/page/{page}", controller.GetBooks).Methods("GET")
}
