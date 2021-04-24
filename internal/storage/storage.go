package storage

import (
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

type StorageService interface {
	LoadRequest(requestId uuid.UUID) (Request, error)
	InsertRequest(request Request) error
	UpdateRequest(request Request) error

	LoadDefinitions(id string) (*bpmn.Definitions, error)
	InsertDefinitions(definitions *bpmn.Definitions) error
}
