package main

import (
	"nayra/internal/nayra"
	"nayra/internal/repository"
	"nayra/internal/services"
	"nayra/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type ImportBpmnEvent struct {
	Bpmn []byte `json:"bpmn"`
}

type Response struct {
	Success       bool    `json:"success"`
	DefinitionsId *string `json:"definitionsId"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event ImportBpmnEvent) (Response, error) {
	// start
	db, err := services.DynamoDB()
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	nayra.SetupStorageService(db)

	// run
	content := event.Bpmn
	definitions, err := repository.ParseBpmn(content)
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	storage.InsertDefinitions(definitions)
	id := definitions.UUID.String()
	loaded, err := storage.LoadDefinitions(id)
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	return Response{
		Success:       true,
		DefinitionsId: aws.String(loaded.UUID.String()),
	}, nil
}
