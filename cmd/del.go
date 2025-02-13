package cmd

import (
	"fmt"
	"github.com/cqwrr/expenseTracker/internal/expenses"
	"github.com/spf13/cobra"
)

var DeleteExpenseId int

func NewDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteExpenseCmd(args)
		},
	}

	deleteCmd.Flags().IntVarP(&DeleteExpenseId, "id", "i", 0, "ID of the expense to delete")
	return deleteCmd
}

func RunDeleteExpenseCmd(args []string) error {
	deleted, err := expenses.DelExpenses(DeleteExpenseId)
	if err != nil {
		return err
	}
	fmt.Printf("Удалена трата: ID %d | Категория: %s | Сумма: %.2f | Описание: %s\n",
		deleted.Id, deleted.Category, deleted.Amount, deleted.Description)
	return nil
}
