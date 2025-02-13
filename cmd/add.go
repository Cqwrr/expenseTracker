package main

import (
	"fmt"
	"github.com/cqwrr/expenseTracker/internal/expenses"
	"github.com/spf13/cobra"
)

var Description string
var Amount float64
var Category string

func newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds new Expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddExpenceCmd(args)
		},
	}

	addCmd.Flags().StringVarP(&Description, "description", "d", "", "Description of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "Amount of the expense")
	addCmd.MarkFlagFilename("amount")
	addCmd.Flags().StringVarP(&Category, "category", "c", "general", "Category of the expence")

	return addCmd
}

func RunAddExpenceCmd(args []string) error {
	if Amount < 0 {
		return fmt.Errorf("ammount cannot be negative")
	}
	return expenses.AddExpense(Amount, Category, Description)
}
