package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"nayra/internal/repository"
	"nayra/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

// callProcessCmd initialize CallProcess command
func importBpmnCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "import [file]",
		Short: "Import a BPMN file",
		Run:   importBpmn(),
		Args:  cobra.MinimumNArgs(1),
	}
	return command
}

// importBpmn imports a BPMN file into the storage
func importBpmn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		path := args[0]
		xmlFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer xmlFile.Close()
		content, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		definitions, err := repository.ParseBpmn(content)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = storage.InsertDefinitions(definitions)
		if err != nil {
			panic(err)
		}
		id := definitions.UUID.String()
		loaded, err := storage.LoadDefinitions(id)
		if err != nil {
			panic(err)
		}
		fmt.Println(loaded.UUID)
	}
}
