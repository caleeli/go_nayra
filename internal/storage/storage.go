package storage

import "github.com/google/uuid"

type StorageService interface {
	LoadRequest(requestId uuid.UUID) (Request, error)
	InsertRequest(request Request) error
	UpdateRequest(request Request) error
}
