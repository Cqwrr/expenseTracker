package expenses

import (
	"encoding/json"
	"fmt"
	"io"
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
	ReadExpensesFile()
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
		// Если ошибка равна EOF, то файл пустой, инициализируем пустой срез расходов.
		if err == io.EOF {
			expenses = []*expense{}
			return nil
		}
		return err
	}
	return nil
}

func DelExpenses(id int) (*expense, error) {
	ReadExpensesFile()

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

func EditExpense(id int, amount float64, category string, description string, date *time.Time) (*expense, error) {
	ReadExpensesFile()

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

func ReadExpensesFile() {
	if err := loadExpenses(); err != nil {
		// Если ошибка не связана с отсутствием файла, вызываем панику
		if !os.IsNotExist(err) {
			panic(fmt.Errorf("ошибка загрузки трат: %v", err))
		}
		// Если файл не найден, инициализируем пустой срез
		expenses = []*expense{}
	}

}

func ListExpenses() {
	if err := loadExpenses(); err != nil {
		fmt.Printf("Ошибка загрузки расходов: %v\n", err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("Нет записей о тратах")
		return
	}

	for _, exp := range expenses {
		fmt.Printf("ID: %d | Категория: %s | Сумма: %.2f | Описание: %s | Дата: %s\n",
			exp.Id, exp.Category, exp.Amount, exp.Description, exp.Date.Format("2006-01-02 15:04:05"))
	}
}
