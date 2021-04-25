package cmd

import (
	"nayra/internal/nayra"
	"nayra/internal/storage"

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
