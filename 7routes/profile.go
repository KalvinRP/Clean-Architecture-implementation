package routes

import (
	"dewetour/2pkg/mysql"
	repositories "dewetour/4repositories"
	handlers "dewetour/6handlers"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profile", h.FindProfile).Methods("GET")
	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
	r.HandleFunc("/profile", h.MakeProfile).Methods("POST")
	r.HandleFunc("/profile/{id}", h.EditProfile).Methods("PATCH")
	r.HandleFunc("/profile/{id}", h.DeleteProfile).Methods("DELETE")
}
