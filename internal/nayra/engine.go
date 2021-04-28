package nayra

import (
	"nayra/internal/bpmn"
	"nayra/internal/errors"
	"nayra/internal/repository"
	"nayra/internal/storage"

	"github.com/google/uuid"
)

func SetupStorageService(store storage.StorageService) {
	storage.SetupStorageService(store)
}

func CallProcess(definitionsId string, processId string) (request repository.Request, err error) {
	definitions, err := storage.LoadDefinitions(definitionsId)
	if err != nil {
		return nil, err
	}
	process := definitions.GetProcess(processId)
	if process == nil {
		return nil, errors.WrapEngineError(nil, "Process not found (ID=%s)", processId)
	}
	request, err = storage.NewRequest(definitions, 1)
	if err != nil {
		return nil, err
	}
	instance := request.GetInstance(0)
	bpmn.CreateToken(instance, &process.Execute)
	if err != nil {
		return nil, err
	}
	err = storage.InsertRequest(request)
	if err != nil {
		return nil, err
	}
	// @todo queue
	instance.NextTick(definitions)
	err = storage.UpdateRequest(request)
	if err != nil {
		return nil, err
	}
	return
}

func TransitToken(requestId, tokenId uuid.UUID, transition string) (request repository.Request, err error) {
	request, err = storage.LoadRequest(requestId)
	if err != nil {
		return nil, err
	}
	token, err := request.GetToken(tokenId)
	if err != nil {
		return nil, err
	}
	token.Transition = transition
	definitions := request.GetDefinitions()
	instance := token.Instance
	// @todo queue
	instance.NextTick(definitions)
	err = storage.UpdateRequest(request)
	if err != nil {
		return nil, err
	}
	return
}
