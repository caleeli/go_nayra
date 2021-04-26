package cmd

import (
	"fmt"
	"log"
	"nayra/internal/nayra"

	"github.com/spf13/cobra"
)

// callProcessCmd initialize CallProcess command
func callProcessCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "call [DEFINITIONS_ID]",
		Short: "Call a process",
		Run:   callProcess(),
		Args:  cobra.MinimumNArgs(1),
	}
	command.Flags().StringP("process", "p", "", "id of the process")
	return command
}

// callProcess start a process by definitions and processId
func callProcess() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		definitionsId := args[0]
		processId, _ := cmd.Flags().GetString("process")
		request, err := nayra.CallProcess(definitionsId, processId)
		if err != nil {
			log.Fatal(err)
		}
		request.TraceLog()
		fmt.Println(request.GetId().String())
	}
}
