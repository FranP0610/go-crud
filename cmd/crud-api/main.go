package main

import (
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.Connection()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":8080", r)

}
