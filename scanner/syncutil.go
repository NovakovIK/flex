package scanner

import (
	"github.com/NovakovIK/flex/storage"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"path"
)

type SyncUtil struct {
	storage *storage.Storage
	scanner *Scanner
}

func NewSyncUtil(storage *storage.Storage, scanner *Scanner) *SyncUtil {
	return &SyncUtil{
		storage: storage,
		scanner: scanner,
	}
}

func (s SyncUtil) Start() error {
	for {
		select {
		case filePath := <-s.scanner.CreatedFiles:
			log.Info("Created file: " + filePath)
			media := storage.Media{
				Name:     path.Base(filePath),
				Path:     filePath,
				Status:   storage.Available,
				Created:  int(rand.Int63()),
				Duration: int(rand.Int63()),
			}

			if err := s.storage.MediaDAO.Insert(media); err != nil {
				return err
			}

		case filePath := <-s.scanner.RemovedFiles:
			log.Info("Removed file: " + filePath)
			if err := s.storage.MediaDAO.DeleteByPath(filePath); err != nil {
				return err
			}

		}
	}
}
