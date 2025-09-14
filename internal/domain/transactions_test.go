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
		{Date: "2025-01-01", Amount: "3200.00", Description: "January Salary", From: "Delivery Hero", To: "Sanju", Category: "Income"},
		{Date: "2025-01-02", Amount: "54.37", Description: "Groceries", From: "Sanju", To: "Rewe", Category: "Groceries"},
		{Date: "2025-09-02", Amount: "22.02", Description: "Groceries", From: "Sushil", To: "Lidl", Category: "Groceries"},
	}, nil
}

func Test_GetAllTransactionsForValidUser(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), "Sanju")

	expectedTransactions := []Transaction{
		{Date: "2025-01-01", Amount: "3200.00", Description: "January Salary", From: "Delivery Hero", To: "Sanju", Category: "Income"},
		{Date: "2025-01-02", Amount: "54.37", Description: "Groceries", From: "Sanju", To: "Rewe", Category: "Groceries"},
	}

	assert.Equal(t, expectedTransactions, trs)
	assert.Equal(t, nil, err)
}

func Test_GetAllTransactionsNoUserProvided(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), "")

	assert.Equal(t, []Transaction{}, trs)
	assert.Equal(t, errors.New("user is not provided"), err)
}

func Test_GetAllTransactionsUserProvidedButNoTransactions(t *testing.T) {
	ts := NewTransactionsService(&mockRepository{})
	trs, err := ts.GetAllTransactions(context.Background(), "Shyam")

	assert.Equal(t, []Transaction{}, trs)
	assert.Equal(t, nil, err)
}
