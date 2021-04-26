package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"nayra/internal/nayra"
	"nayra/internal/storage"

	"github.com/spf13/cobra"
)

// showRequestCmd initialize showRequest command
func showRequestCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "request [ID]",
		Short: "Get request info",
		Run:   showRequest(),
		Args:  cobra.MinimumNArgs(1),
	}
	return command
}

// showRequest start a process by definitions and processId
func showRequest() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		requestId := args[0]
		request, err := nayra.GetRequest(requestId)
		if err != nil {
			log.Fatal(err)
		}
		res, _ := json.Marshal(storage.MarshalRequest(request))
		fmt.Println(string(res))
	}
}
