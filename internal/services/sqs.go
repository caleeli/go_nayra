package services

import (
	"context"
	"encoding/json"
	"nayra/internal/errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type tSQSService struct {
	svc *sqs.Client
	ctx *context.Context
	url *string
}

func SQS() (QueueService, error) {
	service := &tSQSService{}
	endpoint := os.Getenv("AWS_ENDPOINT_URL")
	service.url = aws.String(os.Getenv("SQS_URL"))

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if endpoint != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithEndpointResolver(service))
	}
	if err != nil {
		return nil, errors.WrapServiceError(err, "Unable to init SQS config")
	}
	svc := sqs.NewFromConfig(cfg)
	service.svc = svc
	service.ctx = &ctx
	return service, nil
}

func (queue *tSQSService) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL: os.Getenv("AWS_ENDPOINT_URL"),
	}, nil
}

func (queue *tSQSService) NotifyEvent(event string, body interface{}) (err error) {
	msgBody, err := json.Marshal(body)
	_, err = queue.svc.SendMessage(*queue.ctx, &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Event": {
				DataType:    aws.String("String"),
				StringValue: aws.String(event),
			},
		},
		MessageBody: aws.String(string(msgBody)),
		QueueUrl:    queue.url,
	})
	return
}
