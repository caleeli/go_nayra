package main

import (
	"log"
	"nayra/internal/nayra"
	"nayra/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type TransitTokenEvent struct {
	RequestId  string `json:"requestId"`
	TokenId    string `json:"tokenId"`
	Transition string `json:"transition"`
}

type Response struct {
	Success   bool     `json:"status"`
	RequestId *string  `json:"requestId"`
	Trace     []string `json:"trace"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event TransitTokenEvent) (Response, error) {
	// start
	db, err := services.DynamoDB()
	if err != nil {
		return Response{
			Success: false,
		}, err
	}
	nayra.SetupStorageService(db)

	// run
	requestId, _ := uuid.Parse(event.RequestId)
	tokenId, _ := uuid.Parse(event.TokenId)
	transition := event.Transition
	request, err := nayra.TransitToken(requestId, tokenId, transition)
	if err != nil {
		log.Fatal(err)
	}
	return Response{
		RequestId: aws.String(request.GetId().String()),
		Trace:     request.Trace(),
	}, err
}
