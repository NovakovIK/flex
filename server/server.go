package main

import (
	"context"
	"flag"
	"github.com/99designs/gqlgen/handler"
	"github.com/NovakovIK/flex"
	"github.com/NovakovIK/flex/resolvers"
	"github.com/NovakovIK/flex/scanner"
	"github.com/NovakovIK/flex/storage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	mediaDir := flag.String("media-dir", "~/Videos", "path to directory with videos")
	flag.Parse()

	s := storage.NewStorage()

	scan := scanner.NewScanner(s, *mediaDir)
	sync := scanner.NewSyncUtil(s, scan)

	go func() {
		if err := scan.Scan(); err != nil {
			log.Error(err)
		}
	}()

	go func() {
		if err := sync.Start(); err != nil {
			log.Error(err)
		}
	}()

	router := mux.NewRouter()
	router.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(flex.NewExecutableSchema(flex.Config{Resolvers: resolvers.NewResolver(s)})))
	router.HandleFunc("/videos/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		media, err := s.MediaDAO.FetchByID(id)
		if err != nil || len(media) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err)
			return
		}

		http.ServeFile(w, r, media[0].Path)
	})
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./server/static/")))
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
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
