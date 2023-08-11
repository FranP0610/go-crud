package main

import (
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	"github.com/FranP0610/go-crud/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
}

func main() {
	// Initialize DB connection
	database.Connection()

	// Begin migrations of the 3 databases
	database.DB.AutoMigrate(&domain.Track{})
	database.DB.AutoMigrate(domain.Album{})
	database.DB.AutoMigrate(&domain.Artist{})

	// Defining the router using Gorilla/MUX and Http mod
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":8080", r)

}
