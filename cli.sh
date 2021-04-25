#!/bin/bash
export `less .env | xargs`
go run client/cli/main.go "$@"
