package main

import (
	"nayra/internal/nayra"
	"nayra/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type CallProcessEvent struct {
	DefinitionsId string `json:"definitionsId"`
	ProcessId     string `json:"processId"`
}

type Response struct {
	Success   bool     `json:"success"`
	RequestId *string  `json:"requestId"`
	Trace     []string `json:"trace"`
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

	// run
	definitionsId := event.DefinitionsId
	processId := event.ProcessId
	request, err := nayra.CallProcess(definitionsId, processId)
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	request.TraceLog()
	return Response{
		Success:   true,
		RequestId: aws.String(request.GetId().String()),
		Trace:     request.Trace(),
	}, nil
}
