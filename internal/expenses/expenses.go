package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	loadExpenses()
	editExpense(2, 100, "food", "Aaaaaaaaaaaaaaaaaaa", nil)
}

func newExpense(id int, amount float64, category string, date time.Time, description string) *expense {
	return &expense{
		Id:          id,
		Amount:      amount,
		Category:    category,
		Date:        time.Now(),
		Description: description,
	}
}

var expenses []*expense // Глобальный слайс трат
func AddExpense(amount float64, category string, description string) *expense {
	var newExpenseId int
	if len(expenses) == 0 {
		newExpenseId = 1
	} else {
		newExpenseId = expenses[len(expenses)-1].Id + 1
	}

	exp := newExpense(newExpenseId, amount, category, time.Now(), description)
	expenses = append(expenses, exp) //сохранение в слайс трат
	saveExpenses()
	return exp
}

func saveExpenses() {
	f, err := os.Create("expenses.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(expenses)
	if err != nil {
		panic(err)
	}
}

func loadExpenses() error {
	f, err := os.Open("expenses.json")
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
			saveExpenses()
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
			saveExpenses()
			return exp, nil
		}
	}
	return nil, fmt.Errorf("Трата с id %d не найдена", id)
}
