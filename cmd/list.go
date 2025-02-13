package cmd

import (
	"github.com/cqwrr/expenseTracker/internal/expenses"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListExpenseCmd(args)
		},
	}

	listCmd.Flags().StringVarP(&Category, "category", "c", "all", "Filter expenses by category")
	return listCmd
}

func RunListExpenseCmd(args []string) error {
	expenses.ListExpenses()
	return nil
}
