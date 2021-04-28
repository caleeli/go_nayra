package main

import (
	"nayra/client/cli/cmd"
	"nayra/internal/services"
)

func main() {
	db, err := services.LoadStorage()
	if err != nil {
		panic(err)
	}
	queue, err := services.LogQueue()
	if err != nil {
		panic(err)
	}
	nayraCli := cmd.Nayra(db, queue)
	nayraCli.Execute()
}
