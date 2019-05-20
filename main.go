package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/NovakovIK/flex-server/flex/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
		Duration: 20 * 30,
	},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/media", func(w http.ResponseWriter, r *http.Request) {
		var media []data.Media
		media = mediaHardCode
		mediaJSON, err := json.Marshal(media)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		if _, err = w.Write(mediaJSON); err != nil {
			log.Error(err)
		}

	})

	r.HandleFunc("/api/media/{mediaID}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	waitForClosingConnections := make(chan struct{})
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		signal.Notify(signalChan, syscall.SIGTERM)

		<-signalChan

		log.Infoln("Shutting down server...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
		close(waitForClosingConnections)

	}()
	log.Infof("Starting server on http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-waitForClosingConnections
}
