package http

import (
	"encoding/json"
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	"github.com/gorilla/mux"
	"net/http"
)

func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("get all artists"))
	var artist []domain.Artist
	database.DB.Find(&artist)
	json.NewEncoder(w).Encode(&artist)

}

func GetOneArtistHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Get just one artist"))
	var artist domain.Artist
	params := mux.Vars(r)
	database.DB.First(&artist, params["id"])
	//json.NewEncoder(w).Encode(params)
	if artist.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Artist not found"))
		return // end the request
	}
	database.DB.Model(&artist).Association("Albums").Find(&artist.Albums)
	json.NewEncoder(w).Encode(&artist)
}

func PostArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist domain.Artist
	json.NewDecoder(r.Body).Decode(&artist)
	cratedArtist := database.DB.Create(&artist)
	err := cratedArtist.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&artist)

	//w.Write([]byte("Post artist"))
}

func DeleteArtistHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var artist domain.Artist
	database.DB.First(&artist, params["id"])
	if artist.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Artist not found"))
		return // end the request
	}
	database.DB.Unscoped().Delete(&artist)
	w.WriteHeader(http.StatusNoContent)
}
