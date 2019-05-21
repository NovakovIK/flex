package api

import (
	"encoding/json"
	"github.com/NovakovIK/flex-server/flex/storage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"

	"strconv"
)

var profileHardCode = []storage.Profile{}

func (s *Server) ProfileList(w http.ResponseWriter, r *http.Request) {
	profile := profileHardCode
	profileJSON, err := json.Marshal(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	if _, err = w.Write(profileJSON); err != nil {
		log.Error(err)
	}
}

func (s *Server) ProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID, err := strconv.ParseInt(vars["profileID"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	for _, m := range profileHardCode {
		if m.ID == profileID {
			mediaJSON, err := json.Marshal(m)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Error(err)
				return
			}

			if _, err = w.Write(mediaJSON); err != nil {
				log.Error(err)
			}
			return
		}
	}
	http.NotFound(w, r)
}
