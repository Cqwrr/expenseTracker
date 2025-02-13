package expenses

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type expense struct {
	Id          int
	Category    string
	Amount      float64
	Description string
	Date        time.Time
}

// Error implements error.
func (e *expense) Error() string {
	panic("unimplemented")
}

func newExpense(id int, category string, amount float64, description string, date time.Time) *expense {
	return &expense{
		Id:          id,
		Category:    category,
		Amount:      amount,
		Description: description,
		Date:        time.Now(),
	}
}

var expenses []*expense

func AddExpense(amount float64, category string, description string) *expense {
	var NewExpenseID int
	if len(expenses) == 0 {
		NewExpenseID = 1
	} else {
		NewExpenseID = expenses[len(expenses)-1].Id + 1
	}
	exp := newExpense(NewExpenseID, category, amount, description, time.Now())
	expenses = append(expenses, exp)
	saveExpense()
	return exp
}

func saveExpense() {
	if err := os.MkdirAll("data", 0755); err != nil {
		panic(err)
	}

	f, err := os.Create("data/expenses.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(expenses); err != nil {
		panic(err)
	}
}

func loadExpenses() error {
	f, err := os.Open("data/expenses.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&expenses)
	if err != nil {
		return err
	}
	return nil
}

func delExpenses(id int) (*expense, error) {
	for i, exp := range expenses {
		if exp.Id == id {
			deletedExp := exp
			expenses = append(expenses[:i], expenses[i+1:]...)
			for j := i; j < len(expenses); j++ {
				expenses[j].Id = j + 1
			}
			saveExpense()
			return deletedExp, nil
		}
	}
	return nil, fmt.Errorf("Трата с id %d не найдена", id)
}

func editExpense(id int, amount float64, category string, description string, date *time.Time) (*expense, error) {
	for _, exp := range expenses {
		if exp.Id == id {
			exp.Amount = amount
			exp.Category = category
			exp.Description = description
			if date != nil {
				exp.Date = *date
			}
			saveExpense()
			return exp, nil
		}
	}
	return nil, fmt.Errorf("Трата с id %d не найдена", id)
}

func ListExpenses() {

}
