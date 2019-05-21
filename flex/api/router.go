package api

import (
	"github.com/NovakovIK/flex-server/flex/storage"
	"github.com/gorilla/mux"
)

func NewRouter(s *storage.Storage) *mux.Router {
	server := NewServer(s)

	r := mux.NewRouter()
	r.HandleFunc("/api/media", server.MediaList).Methods("GET")
	r.HandleFunc("/api/media/{mediaID}", server.MediaByID).Methods("GET")
	r.HandleFunc("/api/profile", server.ProfileList).Methods("GET")
	r.HandleFunc("/api/profile/{profileID}", server.ProfileByID).Methods("GET")

	return r
}
