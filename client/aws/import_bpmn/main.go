package main

import (
	"nayra/internal/nayra"
	"nayra/internal/repository"
	"nayra/internal/services"
	"nayra/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"
)

type ImportBpmnEvent struct {
	Bpmn []byte `json:"bpmn"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event ImportBpmnEvent) (string, error) {
	// start
	db, err := services.LoadStorage()
	if err != nil {
		panic(err)
	}
	nayra.SetupStorageService(db)

	// run
	content := event.Bpmn
	definitions, err := repository.ParseBpmn(content)
	if err != nil {
		return "FAILURE", err
	}
	storage.InsertDefinitions(definitions)
	id := definitions.UUID.String()
	loaded, err := storage.LoadDefinitions(id)
	if err != nil {
		return "FAILURE", err
	}
	return loaded.ID, nil
}
