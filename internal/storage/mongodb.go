package storage

import (
	"context"
	"fmt"
	"nayra/internal/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client   *mongo.Client
	database *mongo.Database
	requests *mongo.Collection
	ctx      context.Context
}

func NewMongoDB(
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	DB_NAME string,
) (StorageService, error) {
	var uri string
	if DB_HOST == "" {
		DB_HOST = "127.0.0.1"
	}
	if DB_PORT == "" {
		DB_PORT = "27017"
	}
	if DB_NAME == "" {
		DB_NAME = "go_nayra"
	}
	if DB_USER == "" {
		uri = fmt.Sprintf(
			"mongodb://%s:%s",
			DB_HOST,
			DB_PORT,
		)
	} else if DB_PASSWORD == "" {
		uri = fmt.Sprintf(
			"mongodb://%s@%s:%s",
			DB_USER,
			DB_HOST,
			DB_PORT,
		)
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s",
			DB_USER,
			DB_PASSWORD,
			DB_HOST,
			DB_PORT,
		)
	}
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
		db.requests = db.database.Collection("requests")
	}
	_, err = db.requests.InsertOne(db.ctx, marshalRequest(request))
	return
}

func (db *mongoDB) UpdateRequest(request Request) (err error) {
	if db.requests == nil {
		db.requests = db.database.Collection("requests")
	}
	filter := bson.D{{Key: "id", Value: request.GetId().String()}}
	bm := marshalRequest(request)
	update := bson.D{{Key: "$set", Value: bm}}

	_, err = db.requests.UpdateOne(
		db.ctx,
		filter,
		update,
	)
	return
}
