package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"

	"strconv"
)

func (s *Server) MediaList(w http.ResponseWriter, r *http.Request) {
	media, err := s.storage.MediaDAO.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	mediaJSON, err := json.Marshal(media)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	if _, err = w.Write(mediaJSON); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}
}

func (s *Server) MediaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID, err := strconv.ParseInt(vars["mediaID"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error(err)
		return
	}

	media, err := s.storage.MediaDAO.FetchByID(mediaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	mediaJSON, err := json.Marshal(media)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	if _, err = w.Write(mediaJSON); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}
}
