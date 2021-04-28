package nayra

import (
	"nayra/internal/errors"
	"nayra/internal/repository"
	"nayra/internal/storage"

	"github.com/google/uuid"
)

func GetRequest(requestId string) (request repository.Request, err error) {
	id, err := uuid.Parse(requestId)
	if err != nil {
		return nil, errors.WrapEngineError(err, "Invalid request id")
	}
	request, err = storage.LoadRequest(id)
	if err != nil {
		return nil, errors.WrapEngineError(err, "Unable to load request %s", requestId)
	}
	return
}
