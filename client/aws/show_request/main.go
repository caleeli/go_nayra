package main

import (
	"log"
	"nayra/internal/nayra"
	"nayra/internal/services"
	"nayra/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"
)

type ShowRequestEvent struct {
	RequestId string `json:"requestId"`
}

type Response struct {
	Success bool             `json:"status"`
	Data    storage.SRequest `json:"data"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event ShowRequestEvent) (Response, error) {
	// start
	db, err := services.DynamoDB()
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	nayra.SetupStorageService(db)

	// run
	request, err := nayra.GetRequest(event.RequestId)
	if err != nil {
		log.Fatal(err)
	}
	return Response{
		Success: true,
		Data:    storage.MarshalRequest(request),
	}, nil
}
