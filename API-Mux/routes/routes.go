package routes

import (
	"api-mux/resource"

	"github.com/gorilla/mux"
)

func AppRoutes(r *mux.Router) {

	r.HandleFunc("/v1/users", resource.GetUsers).Methods("GET")
	r.HandleFunc("/v1/users/{id}", resource.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", resource.CreateUser).Methods("POST")
	r.HandleFunc("/v1/user/{id}", resource.UpdateUser).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", resource.DeleteUser).Methods("DELETE")

}
