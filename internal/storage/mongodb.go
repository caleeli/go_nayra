package storage

import (
	"context"
	"fmt"
	"nayra/internal/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageService interface {
	InsertRequest(request Request) error
}

const (
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_HOST     = "localhost"
	DB_PORT     = 27017
	DB_NAME     = "go_nayra"
)

type mongoDB struct {
	client   *mongo.Client
	database *mongo.Database
	requests *mongo.Collection
	ctx      context.Context
}

func NewMongoDB() (StorageService, error) {
	//uri := fmt.Sprintf(
	//	"mongodb://%s:%s@%s:%d",
	//	DB_USER,
	//	DB_PASSWORD,
	//	DB_HOST,
	//	DB_PORT,
	//)
	uri := fmt.Sprintf(
		"mongodb://%s:%d",
		DB_HOST,
		DB_PORT,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.WrapServiceError(err, "Unable to start MongoDB Service")
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel() // releases resources if slowOperation completes before timeout elapses
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(err, "Unable to start MongoDB Service")
	}
	service := &mongoDB{
		client:   client,
		database: client.Database(DB_NAME),
		ctx:      ctx,
	}
	return service, nil
}

func (db *mongoDB) InsertRequest(request Request) (err error) {
	if db.requests == nil {
		db.requests = db.database.Collection("test")
	}
	req := request.(*tRequest)
	_, err = db.requests.InsertOne(db.ctx, marshalRequest(req))
	return
}
