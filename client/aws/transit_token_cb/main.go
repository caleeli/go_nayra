package main

import (
	"log"
	"nayra/internal/nayra"
	"nayra/internal/services"
	"nayra/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type TransitTokenEvent struct {
	RequestId  string `json:"requestId"`
	TokenId    string `json:"tokenId"`
	Transition string `json:"transition"`
}

type TransitTokenEventCB struct {
	RequestId string
	TokenId   string
	Script    string
}

type Response struct {
	Success bool             `json:"status"`
	Id      *string          `json:"id"`
	Data    storage.SRequest `json:"data"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(eventCB TransitTokenEventCB) (Response, error) {
	event := TransitTokenEvent{
		eventCB.RequestId,
		eventCB.TokenId,
		"COMPLETE",
	}
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
	log.Println(event.RequestId, event.TokenId, event.Transition)
	requestId, _ := uuid.Parse(event.RequestId)
	tokenId, _ := uuid.Parse(event.TokenId)
	transition := event.Transition
	request, err := nayra.TransitToken(requestId, tokenId, transition)
	if err != nil {
		log.Fatal(err)
	}
	return Response{
		Success: true,
		Id:      aws.String(request.GetId().String()),
		Data:    storage.MarshalRequest(request),
	}, err
}
