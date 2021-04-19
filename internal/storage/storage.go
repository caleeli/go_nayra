package storage

type StorageService interface {
	InsertRequest(request Request) error
	UpdateRequest(request Request) error
}
