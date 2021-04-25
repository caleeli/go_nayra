package services

import (
	"context"
	"fmt"
	"nayra/internal/bpmn"
	"nayra/internal/repository"
	"nayra/internal/storage"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type tDynamoDB struct {
	svc *dynamodb.Client
	ctx *context.Context
}

func DynamoDB() (storage.StorageService, error) {
	service := &tDynamoDB{}
	endpoint := os.Getenv("AWS_ENDPOINT_URL")
	fmt.Println("AWS_ENDPOINT_URL", endpoint)

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if endpoint != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithEndpointResolver(service))
	}
	if err != nil {
		return nil, err
	}
	svc := dynamodb.NewFromConfig(cfg)
	service.svc = svc
	service.ctx = &ctx
	return service, nil
}
func (db *tDynamoDB) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL: os.Getenv("AWS_ENDPOINT_URL"),
	}, nil
}

func (db *tDynamoDB) LoadRequest(requestId uuid.UUID) (request repository.Request, err error) {
	key, err := attributevalue.MarshalMap(struct {
		Id string
	}{
		Id: requestId.String(),
	})
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("requests"),
	}
	output, err := db.svc.GetItem(*db.ctx, input)
	if err != nil {
		return nil, err
	}
	srequest := &storage.SRequest{}
	err = attributevalue.UnmarshalMap(output.Item, srequest)
	if err != nil {
		return nil, err
	}
	request, err = storage.UnmarshalRequest(srequest)
	return request, err
}

func (db *tDynamoDB) InsertRequest(request repository.Request) (err error) {
	req := storage.MarshalRequest(request)
	item, err := attributevalue.MarshalMap(req)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("requests"),
	}

	_, err = db.svc.PutItem(*db.ctx, input)
	return err
}

func (db *tDynamoDB) UpdateRequest(request repository.Request) (err error) {
	req := storage.MarshalRequest(request)
	key, err := attributevalue.MarshalMap(struct {
		Id string
	}{
		Id: request.GetId().String(),
	})
	if err != nil {
		return err
	}
	item, err := attributevalue.MarshalMap(req)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		Key: key,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":instances": item["Instances"],
		},
		UpdateExpression: aws.String("SET Instances = :instances"),
		TableName:        aws.String("requests"),
	}

	_, err = db.svc.UpdateItem(*db.ctx, input)
	return err
}

func (db *tDynamoDB) InsertDefinitions(definitions *bpmn.Definitions) (err error) {
	sdefinitions := storage.MarshalDefinitions(definitions)
	item, err := attributevalue.MarshalMap(sdefinitions)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("processes"),
	}
	_, err = db.svc.PutItem(*db.ctx, input)
	return err
}

func (db *tDynamoDB) LoadDefinitions(id string) (*bpmn.Definitions, error) {
	key, err := attributevalue.MarshalMap(struct {
		Id string
	}{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("processes"),
	}
	output, err := db.svc.GetItem(*db.ctx, input)
	if err != nil {
		return nil, err
	}
	sdefinitions := &storage.SDefinitions{}
	err = attributevalue.UnmarshalMap(output.Item, sdefinitions)
	if err != nil {
		return nil, err
	}
	definitions, err := storage.UnmarshalDefinitions(sdefinitions)
	if err != nil {
		return nil, err
	}
	definitions.UUID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	definitions.ParseTree()
	return definitions, err

}
