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
	"syscall"
)

func main() {
	mediaDir := flag.String("media-dir", "~/Videos", "path to direcotry with videos")
	flag.Parse()

	s := storage.NewStorage()

	scan := scanner.NewScanner(s, *mediaDir)
	sync := scanner.NewSyncUtil(s, scan)
	go scan.Scan()
	go sync.Start()

	router := mux.NewRouter()
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(flex.NewExecutableSchema(flex.Config{Resolvers: resolvers.NewResolver(s)})))

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
