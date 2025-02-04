package main

import (
	"fmt"
	"time"

	"github.com/cqwrr/expenseTracker/internal/expenses"
)

func main() {

	if err := expenses.LoadExpenses(); err != nil {
		fmt.Println("Ошибка загрузки расходов:", err)
		return
	}

	exp := expenses.AddExpense(50.0, "Transport", "Taxi to work", time.Now())
	fmt.Println("Добавлен расход:", exp)

	fmt.Println("Список расходов:")
	for _, e := range expenses.GetExpenses() {
		fmt.Printf("ID: %d, Amount: %.2f, Category: %s, Date: %s, Description: %s\n",
			e.Id, e.Amount, e.Category, e.Date.Format("2006-01-02"), e.Description)
	}
}
