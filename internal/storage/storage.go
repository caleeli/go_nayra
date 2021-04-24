package storage

import (
	"nayra/internal/bpmn"
	"nayra/internal/repository"

	"github.com/google/uuid"
)

type StorageService interface {
	LoadRequest(requestId uuid.UUID) (repository.Request, error)
	InsertRequest(request repository.Request) error
	UpdateRequest(request repository.Request) error

	LoadDefinitions(id string) (*bpmn.Definitions, error)
	InsertDefinitions(definitions *bpmn.Definitions) error
}

var config struct {
	storage StorageService
}

func SetupStorageService(storage StorageService) {
	config.storage = storage
}
