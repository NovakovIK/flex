package scanner

import (
	"github.com/NovakovIK/flex/storage"
	"github.com/bakape/thumbnailer"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
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
			file, err := os.Open(filePath)
			if err != nil {
				log.Error(err)
				continue
			}

			src, thumb, err := thumbnailer.Process(file, thumbnailer.Options{
				JPEGQuality: 75,
				PNGQuality: struct {
					Min, Max uint
				}{
					Min: 50,
					Max: 100,
				},
				MaxSourceDims: thumbnailer.Dims{
					Width:  3840,
					Height: 2160,
				},
				ThumbDims: thumbnailer.Dims{
					Width:  1920,
					Height: 1080,
				},
			})
			if err != nil {
				log.Error(err)
				continue
			}
			err = file.Close()
			if err != nil {
				log.Fatal(err)
			}

			media := storage.Media{
				Name:      path.Base(filePath),
				Path:      filePath,
				Status:    storage.Available,
				Created:   int(time.Now().UnixNano()),
				Duration:  int(src.Length.Nanoseconds()),
				Thumbnail: thumb.Data,
			}
			if err := s.storage.MediaDAO.Insert(media); err != nil {
				return err
			}
			src.Data = nil
			thumb.Data = nil
			log.Infof("%#v\n", src)
			log.Infof("%#v\n", thumb)
			log.Infof("%#v\n", media)

		case filePath := <-s.scanner.RemovedFiles:
			log.Info("Removed file: " + filePath)
			if err := s.storage.MediaDAO.DeleteByPath(filePath); err != nil {
				return err
			}
		}
	}
}
