package routes

import (
	"dewetour/2pkg/mysql"
	repositories "dewetour/4repositories"
	handlers "dewetour/6handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", h.FindAcc).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetAcc).Methods("GET")
	r.HandleFunc("/user", h.MakeAcc).Methods("POST")
	r.HandleFunc("/user/{id}", h.EditAcc).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteAcc).Methods("DELETE")
}
