package main

import (
	"nayra/client/cli"
)

func main() {
	nayraCli := cli.Nayra()
	nayraCli.Execute()
}
