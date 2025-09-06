package domain

type transactionsStorage interface {
	SelectAll() ([]Transaction, error)
}
