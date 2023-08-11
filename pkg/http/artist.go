package http

import (
	"encoding/json"
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	"net/http"
)

func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all artists"))
}

func GetArtistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get just one artist"))
}

func PostArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist domain.Artist
	json.NewDecoder(r.Body).Decode(&artist)
	cratedArtist := database.DB.Create(&artist)
	err := cratedArtist.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&artist)

	//w.Write([]byte("Post artist"))
}

func DeleteArtistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete artist"))
}
