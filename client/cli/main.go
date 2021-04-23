package main

import (
	"nayra/client/cli/cmd"
	"nayra/internal/storage"
	"os"
)

func main() {
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
	nayraCli := cmd.Nayra(db)
	nayraCli.Execute()
}
