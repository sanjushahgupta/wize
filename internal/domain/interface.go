package domain

import "context"

type transactionsStorage interface {
	SelectAll(ctx context.Context) ([]Transaction, error)
}
