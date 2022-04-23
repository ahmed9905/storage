package storage

import (
	"fmt"
	"github.com/ahmed9905/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)

	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(ID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[ID]

	if !ok {
		return nil, fmt.Errorf("file %v not found", ID)
	}

	return foundFile, nil
}
