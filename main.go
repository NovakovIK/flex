package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NovakovIK/flex-server/flex/api"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/media", api.MediaList).Methods("GET")
	r.HandleFunc("/api/media/{mediaID}", api.MediaByID).Methods("GET")
	r.HandleFunc("/api/profile", api.ProfileList).Methods("GET")
	r.HandleFunc("/api/profile/{profileID}", api.ProfileByID).Methods("GET")

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
