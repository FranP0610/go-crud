package http

import (
	"encoding/json"
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var album []domain.Album
	database.DB.Find(&album)
	json.NewEncoder(w).Encode(album)
}

func GetOneAlbum(w http.ResponseWriter, r *http.Request) {
	var album domain.Album
	params := mux.Vars(r)
	database.DB.First(&album, params["id"])
	if album.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Album not found"))
		return
	}
	json.NewEncoder(w).Encode(&album)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album domain.Album
	json.NewDecoder(r.Body).Decode(&album)
	createdAlbum := database.DB.Create(&album)
	err := createdAlbum.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&album)

}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var album domain.Album
	database.DB.First(&album, params["id"])
	if album.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Album not found"))
		return // end the request
	}
	database.DB.Unscoped().Delete(&album, params["id"])
	w.WriteHeader(http.StatusOK)
}
