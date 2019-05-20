package api

import (
	"encoding/json"
	"github.com/NovakovIK/flex-server/flex/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"

	"strconv"
)

var mediaHardCode = []data.Media{
	{
		MediaID:  1,
		Name:     "Big Buck Bunny",
		Hash:     nil,
		Duration: 20 * 60,
	},
	{
		MediaID:  2,
		Name:     "Jojo Bizarre Adventure",
		Hash:     nil,
		Duration: 10 * 60,
	},
}

func MediaList(w http.ResponseWriter, r *http.Request) {
	media := mediaHardCode
	mediaJSON, err := json.Marshal(media)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	if _, err = w.Write(mediaJSON); err != nil {
		log.Error(err)
	}

}

func MediaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID, err := strconv.ParseInt(vars["mediaID"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	for _, m := range mediaHardCode {
		if m.MediaID == mediaID {
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
