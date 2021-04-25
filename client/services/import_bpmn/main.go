package main

import (
	"nayra/internal/nayra"
	"nayra/internal/repository"
	"nayra/internal/storage"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Bpmn []byte `json:"bpmn"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event InputEvent) (string, error) {
	db, err := storage.NewMongoDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB__PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		panic(err)
	}
	nayra.SetupStorageService(db)
	content := event.Bpmn
	definitions, err := repository.ParseBpmn(content)
	if err != nil {
		return "FAILURE", err
	}
	storage.InsertDefinitions(definitions)
	id := definitions.UUID.String()
	loaded, err := storage.LoadDefinitions(id)
	if err != nil {
		return "FAILURE", err
	}
	return loaded.ID, nil
}
