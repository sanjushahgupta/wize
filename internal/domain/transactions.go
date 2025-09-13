package domain

import (
	"context"
	"errors"
	"strings"
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

func (ts *TransactionsService) GetAllTransactions(ctx context.Context, user string) ([]Transaction, error) {
	if user == "" {
		return make([]Transaction, 0), NoUserProvidedError
	}

	transactions, err := ts.storage.SelectAll(ctx)
	if err != nil {
		return transactions, err
	}

	normalizedUser := strings.ToLower(user)
	myTransactions := make([]Transaction, 0)
	for _, transaction := range transactions {
		normalizedFrom := strings.ToLower(transaction.From)
		normalizedTo := strings.ToLower(transaction.To)
		if normalizedFrom == normalizedUser || normalizedTo == normalizedUser {
			myTransactions = append(myTransactions, transaction)
		}
	}

	return myTransactions, nil
}
