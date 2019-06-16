package scanner

import (
	"github.com/NovakovIK/flex/storage"
	"os"
	"path/filepath"
)

type Scanner struct {
	CreatedFiles chan string
	RemovedFiles chan string
	storage      *storage.Storage
	path         string
}

func NewScanner(stor *storage.Storage, dirPath string) *Scanner {
	return &Scanner{
		CreatedFiles: make(chan string),
		RemovedFiles: make(chan string),
		storage:      stor,
		path:         dirPath,
	}
}

func (s Scanner) Scan() error {
	mediaList, err := s.storage.MediaDAO.FetchAll()
	if err != nil {
		return err
	}

	pathMap := make(map[string]bool)
	for _, media := range mediaList {
		pathMap[media.Path] = true
	}

	if err = s.scanRemoved(pathMap); err != nil {
		return err
	}

	if err = s.scanCreated(pathMap); err != nil {
		return err
	}

	return nil
}

func (s Scanner) scanRemoved(pathMap map[string]bool) error {
	for path := range pathMap {
		stat, err := os.Stat(path)
		if err != nil && os.IsNotExist(err) || err == nil && stat.IsDir() {
			s.RemovedFiles <- path
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (s Scanner) scanCreated(pathMap map[string]bool) error {
	err := filepath.Walk(s.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		_, isStored := pathMap[path]
		if !info.IsDir() && !isStored {
			s.CreatedFiles <- path
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
