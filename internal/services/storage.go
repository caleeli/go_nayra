package services

import (
	"nayra/internal/errors"
	"nayra/internal/storage"
	"os"
)

func LoadStorage() (db storage.StorageService, err error) {
	switch os.Getenv("DB_CONNECTION") {
	case "MONGODB":
		db, err = storage.NewMongoDB(
			os.Getenv("DB_USER"),
			os.Getenv("DB__PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	case "DYNAMODB":
		db, err = DynamoDB()
	default:
		err = errors.WrapStorageError(nil, "Undifined environment variable DB_CONNECTION")
	}
	return
}
