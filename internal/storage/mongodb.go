package storage

import (
	"context"
	"fmt"
	"nayra/internal/bpmn"
	"nayra/internal/errors"
	"nayra/internal/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client    *mongo.Client
	database  *mongo.Database
	requests  *mongo.Collection
	processes *mongo.Collection
	ctx       context.Context
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

func (db *mongoDB) getRequests() *mongo.Collection {
	if db.requests == nil {
		db.requests = db.database.Collection("requests")
	}
	return db.requests
}

func (db *mongoDB) getProcessses() *mongo.Collection {
	if db.processes == nil {
		db.processes = db.database.Collection("processes")
	}
	return db.processes
}

func (db *mongoDB) LoadRequest(requestId uuid.UUID) (request repository.Request, err error) {
	requests := db.getRequests()
	filter := bson.D{{Key: "id", Value: requestId.String()}}
	srequest := &SRequest{}
	err = requests.FindOne(db.ctx, filter).Decode(srequest)
	if err != nil {
		return nil, err
	}
	request, err = UnmarshalRequest(srequest)
	return request, err
}

func (db *mongoDB) InsertRequest(request repository.Request) (err error) {
	requests := db.getRequests()
	_, err = requests.InsertOne(db.ctx, MarshalRequest(request))
	return
}

func (db *mongoDB) UpdateRequest(request repository.Request) (err error) {
	if db.requests == nil {
		db.requests = db.database.Collection("requests")
	}
	filter := bson.D{{Key: "id", Value: request.GetId().String()}}
	bm := MarshalRequest(request)
	update := bson.D{{Key: "$set", Value: bm}}

	_, err = db.requests.UpdateOne(
		db.ctx,
		filter,
		update,
	)
	return
}

func (db *mongoDB) InsertDefinitions(definitions *bpmn.Definitions) (err error) {
	processes := db.getProcessses()
	record := MarshalDefinitions(definitions)
	_, err = processes.InsertOne(db.ctx, record)
	return
}

func (db *mongoDB) LoadDefinitions(id string) (*bpmn.Definitions, error) {
	requests := db.getProcessses()
	filter := bson.D{{Key: "id", Value: id}}
	sdefinitions := &SDefinitions{}
	err := requests.FindOne(db.ctx, filter).Decode(sdefinitions)
	if err != nil {
		return nil, err
	}
	definitions, err := UnmarshalDefinitions(sdefinitions)
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
