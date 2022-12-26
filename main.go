package main

import (
	mysql "dewetour/2pkg/mysql"
	database "dewetour/3database"
	routes "dewetour/7routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()

	database.Migrate()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Server 5000 running")
	http.ListenAndServe("localhost:5000", r)
}
