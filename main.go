package main

import (
	mysql "dewetour/2pkg/mysql"
	database "dewetour/3database"
	routes "dewetour/7routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.Migrate()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Server 5000 running")
	http.ListenAndServe("localhost:5000", r)
}
