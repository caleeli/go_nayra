package cli

import (
	"log"
	"nayra/internal/nayra"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

// Nayra cli
func Nayra() *cobra.Command {
	rootCmd := &cobra.Command{Use: "nayra"}
	rootCmd.AddCommand(initCallProcessCmd())
	return rootCmd
}

// initCallProcessCmd initialize beers command
func initCallProcessCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "call",
		Short: "Call a process",
		Run:   callProcess(),
	}
	command.Flags().StringP("definitions", "d", "", "id of the definitions")
	command.Flags().StringP("process", "p", "", "id of the process")
	return command
}

func callProcess() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		definitionsId, _ := cmd.Flags().GetString("definitions")
		processId, _ := cmd.Flags().GetString("process")
		request, err := nayra.CallProcess(definitionsId, processId)
		if err != nil {
			log.Fatal(err)
		}
		request.TraceLog()
	}
}
