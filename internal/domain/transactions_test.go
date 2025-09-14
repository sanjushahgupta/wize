package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct{}

func (m *mockRepository) SelectAll(ctx context.Context) ([]Transaction, error) {
	return []Transaction{
		{Date: "2025-01-01", Amount: 3200.00, Description: "January Salary", From: 1234, To: 4567, Category: "Income"},
		{Date: "2025-01-02", Amount: 54.37, Description: "Groceries", From: 1234, To: 7890, Category: "Groceries"},
		{Date: "2025-09-02", Amount: 22.02, Description: "Groceries", From: 7890, To: 2222, Category: "Groceries"},
	}, nil
}

func Test_GetAllTransactionsForValidUser(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), 1234)

	expectedTransactions := []Transaction{
		{Date: "2025-01-01", Amount: 3200.00, Description: "January Salary", From: 1234, To: 4567, Category: "Income"},
		{Date: "2025-01-02", Amount: 54.37, Description: "Groceries", From: 1234, To: 7890, Category: "Groceries"},
	}

	assert.Equal(t, expectedTransactions, trs)
	assert.Equal(t, nil, err)
}

func Test_GetAllTransactionsNoUserProvided(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), 0)

	assert.Equal(t, []Transaction{}, trs)
	assert.Equal(t, errors.New("user is not provided"), err)
}

func Test_GetAllTransactionsUserProvidedButNoTransactions(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), 345689)

	assert.Equal(t, []Transaction{}, trs)
	assert.Equal(t, nil, err)
}
