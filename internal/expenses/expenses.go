package expenses

import (
	"encoding/json"
	"os"
	"time"
)

type expense struct {
	Id          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

var expenses []expense
var filePath = "expenses.json"

func saveExpenses() error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}
func loadExpenses() error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			expenses = []expense{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &expenses)
}

func addExpense(Amount float64, Category string, Description string, Date time.Time) expense {
	id := len(expenses) + 1
	exp := expense{Id: id, Amount: Amount, Category: Category, Date: Date, Description: Description}
	expenses = append(expenses, exp)
	saveExpenses()
	return exp
}

func getExpenses() []expense {
	return expenses
}
