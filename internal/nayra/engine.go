package nayra

import (
	"nayra/internal/errors"
	"nayra/internal/storage"
)

func CallProcess(definitionsId string, processId string) (request storage.Request, err error) {
	definitions, err := storage.LoadBpmn(definitionsId)
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
	uuids := request.GetInstanceUuids()
	instance := request.GetInstance(uuids[0])
	process.Execute.CreateToken(instance)
	if err != nil {
		return nil, err
	}
	storage.SaveRequest(request)
	// @todo queue
	instance.NextTick(definitions)
	storage.SaveRequest(request)
	return
}
