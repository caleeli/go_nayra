package cmd

import (
	"log"
	"nayra/internal/nayra"
	"nayra/internal/storage"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

// Nayra cmd
func Nayra(store storage.StorageService) *cobra.Command {
	nayra.SetupStorageService(store)
	rootCmd := &cobra.Command{Use: "nayra"}
	rootCmd.AddCommand(callProcessCmd())
	rootCmd.AddCommand(transitTokenCmd())
	rootCmd.AddCommand(importBpmnCmd())
	return rootCmd
}

// callProcessCmd initialize CallProcess command
func callProcessCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "call",
		Short: "Call a process",
		Run:   callProcess(),
	}
	command.Flags().StringP("definitions", "d", "", "id of the definitions")
	command.Flags().StringP("process", "p", "", "id of the process")
	return command
}

// callProcess start a process by definitions and processId
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

// transitTokenCmd initialize TransitToken command
func transitTokenCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "transit [TRANSITION]",
		Short: "Transit a token",
		Run:   transitToken(),
		Args:  cobra.MinimumNArgs(1),
	}
	command.Flags().StringP("request", "r", "", "id of the request")
	command.Flags().StringP("token", "t", "", "id of the token")
	cobra.CheckErr(command.MarkFlagRequired("request"))
	cobra.CheckErr(command.MarkFlagRequired("token"))
	return command
}

// transitToken start a process by definitions and processId
func transitToken() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		requestFlag, _ := cmd.Flags().GetString("request")
		tokenFlag, _ := cmd.Flags().GetString("token")
		requestId, _ := uuid.Parse(requestFlag)
		tokenId, _ := uuid.Parse(tokenFlag)
		request, err := nayra.TransitToken(requestId, tokenId, args[0])
		if err != nil {
			log.Fatal(err)
		}
		request.TraceLog()
	}
}
