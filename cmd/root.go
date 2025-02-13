package main

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Expence-tracker",
		Short: "Expence tracker for managing expences",
		Long:  `manage your expences`,
	}

	cmd.AddCommand(newAddCmd())
	return cmd
}
