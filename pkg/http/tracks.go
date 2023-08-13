package http

import (
	"encoding/json"
	"github.com/FranP0610/go-crud/pkg/database"
	"github.com/FranP0610/go-crud/pkg/domain"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTracksHandler(w http.ResponseWriter, r *http.Request) {
	var tracks []domain.Track
	database.DB.Find(&tracks)
	json.NewEncoder(w).Encode(&tracks)
}

func GetOneTrackHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var track domain.Track
	database.DB.First(&track, params["id"])
	if track.ID == 0 {
		w.Write([]byte("Track Not Found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&track)
}

func CreateTrackHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var track domain.Track
	json.NewDecoder(r.Body).Decode(&track)
	createdTrack := database.DB.Create(&track)
	err := createdTrack.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&track)
}

func DeleteTrackHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var track domain.Track
	database.DB.First(&track, params["id"])
	if track.ID == 0 {
		w.Write([]byte("Track Not Found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	database.DB.Unscoped().Delete(&track)
	w.WriteHeader(http.StatusNoContent)

}
