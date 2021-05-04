package main

import (
	"nayra/internal/nayra"
	"nayra/internal/services"
	"nayra/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type CallProcessEvent struct {
	DefinitionsId string `json:"definitionsId"`
	ProcessId     string `json:"processId"`
}

type Response struct {
	Success bool             `json:"success"`
	Id      *string          `json:"id"`
	Data    storage.SRequest `json:"data"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event CallProcessEvent) (Response, error) {
	// start
	db, err := services.DynamoDB()
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	nayra.SetupStorageService(db)
	sqs, err := services.SQS()
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	services.SetupQueueService(sqs)

	// run
	definitionsId := event.DefinitionsId
	processId := event.ProcessId
	request, err := nayra.CallProcess(definitionsId, processId)
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	return Response{
		Success: true,
		Id:      aws.String(request.GetId().String()),
		Data:    storage.MarshalRequest(request),
	}, nil
}
