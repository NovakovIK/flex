package api

import (
	"github.com/NovakovIK/flex-server/flex/storage"
)

type Server struct {
	storage *storage.Storage
}

func NewServer(s *storage.Storage) *Server {
	return &Server{s}
}
