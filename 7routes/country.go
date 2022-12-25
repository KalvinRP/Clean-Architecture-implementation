package routes

import (
	"dewetour/2pkg/mysql"
	repositories "dewetour/4repositories"
	handlers "dewetour/6handlers"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	r.HandleFunc("/country", h.FindCountry).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", h.MakeCountry).Methods("POST")
	r.HandleFunc("/country/{id}", h.EditCountry).Methods("PATCH")
	r.HandleFunc("/country/{id}", h.DeleteCountry).Methods("DELETE")
}
