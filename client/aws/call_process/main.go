package main

import (
	"log"
	"nayra/internal/nayra"
	"nayra/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
)

type CallProcessEvent struct {
	DefinitionsId string `json:"definitionsId"`
	ProcessId     string `json:"processId"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event CallProcessEvent) (string, error) {
	// start
	db, err := services.LoadStorage()
	if err != nil {
		panic(err)
	}
	nayra.SetupStorageService(db)

	// run
	definitionsId := event.DefinitionsId
	processId := event.ProcessId
	request, err := nayra.CallProcess(definitionsId, processId)
	if err != nil {
		log.Fatal(err)
	}
	request.TraceLog()
	return request.GetId().String(), nil
}
