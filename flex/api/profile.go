package api

import (
	"encoding/json"
	"github.com/NovakovIK/flex-server/flex/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"

	"strconv"
)

var profileHardCode = []data.Profile{
	{
		ProfileID: 1,
		Name:      "Alex",
	},
	{
		ProfileID: 2,
		Name:      "Wanya",
	},
}

func ProfileList(w http.ResponseWriter, r *http.Request) {
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
func ProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID, err := strconv.ParseInt(vars["profileID"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	for _, m := range profileHardCode {
		if m.ProfileID == profileID {
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
