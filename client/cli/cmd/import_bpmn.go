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
		Use:   "import",
		Short: "Import a BPMN file",
		Run:   importBpmn(),
		Args:  cobra.MinimumNArgs(1),
	}
	return command
}

// callProcess start a process by definitions and processId
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
		//$bpmn = repository.ParseBpmn($content)
		//$record = storage.MarshalProcess($bpmn)
		//storage.insert($record)
		definitions, err := repository.ParseBpmn(content)
		if err != nil {
			log.Fatal(err)
			return
		}
		storage.InsertDefinitions(definitions)
		id := definitions.UUID.String()
		loaded, err := storage.LoadDefinitions(id)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(loaded.UUID)
	}
}
