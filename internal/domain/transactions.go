package domain

import "wize/internal/repository"

type Transaction struct {
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
	Category    string `json:"category"`
}

func GetAllTransactions() []Transaction {
	transactions := repository.SelectAll()

	return transactions
}
