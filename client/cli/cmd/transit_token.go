package cmd

import (
	"log"
	"nayra/internal/nayra"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

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
