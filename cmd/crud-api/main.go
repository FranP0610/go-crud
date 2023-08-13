package main

import (
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	myhttp "github.com/FranP0610/go-crud/pkg/http"
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
	// Health check of the service
	r.HandleFunc("/health_check/", myhttp.HealthCheckHandler)
	// Artist routes
	r.HandleFunc("/artist", myhttp.GetArtistsHandler).Methods("GET")
	r.HandleFunc("/artist/{id}", myhttp.GetOneArtistHandler).Methods("GET")
	r.HandleFunc("/artist", myhttp.PostArtistHandler).Methods("POST")
	r.HandleFunc("/artist/{id}", myhttp.DeleteArtistHandler).Methods("DELETE")
	// Albums routes
	r.HandleFunc("/albums", myhttp.GetAlbums).Methods("GET")
	r.HandleFunc("/album/{id}", myhttp.GetOneAlbum).Methods("GET")
	r.HandleFunc("/album", myhttp.CreateAlbum).Methods("POST")
	r.HandleFunc("/album/{id}", myhttp.DeleteAlbum).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
