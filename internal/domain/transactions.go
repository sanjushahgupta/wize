package domain

import (
	"context"
	"errors"
	"log"
)

type TransactionsService struct {
	storage transactionsStorage
}

var NoUserProvidedError = errors.New("user is not provided")

func NewTransactionsService(
	storage transactionsStorage,
) *TransactionsService {
	return &TransactionsService{
		storage: storage,
	}
}

func (ts *TransactionsService) GetAllTransactions(ctx context.Context, userID int64) ([]Transaction, error) {
	if userID == 0 {
		return make([]Transaction, 0), NoUserProvidedError
	}

	transactions, err := ts.storage.SelectAll(ctx)
	log.Println("all transactions", transactions)
	if err != nil {
		return transactions, err
	}

	myTransactions := make([]Transaction, 0)

	for _, transaction := range transactions {
		log.Println("userID", transaction.From)
		if transaction.From == userID || transaction.To == userID {
			myTransactions = append(myTransactions, transaction)
		}
	}

	return myTransactions, nil
}
