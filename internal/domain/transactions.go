package domain

type TransactionsService struct {
	storage transactionsStorage
}

func NewTransactionsService(
	storage transactionsStorage,
) *TransactionsService {
	return &TransactionsService{
		storage: storage,
	}
}

func (ts *TransactionsService) GetAllTransactions() ([]Transaction, error) {
	transactions, err := ts.storage.SelectAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
