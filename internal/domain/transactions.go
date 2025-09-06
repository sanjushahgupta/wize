package domain

type Transaction struct {
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
	Category    string `json:"category"`
}

type transactionsStorage interface {
	SelectAll() ([]Transaction, error)
}

type TransactionsService struct {
	Storage transactionsStorage
}

func (ts *TransactionsService) GetAllTransactions() ([]Transaction, error) {
	transactions, err := ts.Storage.SelectAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
