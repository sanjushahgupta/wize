package repository

import (
	"context"
	"encoding/json"
	"os"
	"wize/internal/domain"
)

type FileStorage struct{}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fs *FileStorage) SelectAll(_ context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	data, err := os.ReadFile("db.json")
	if err != nil {
		return transactions, err
	}

	if err := json.Unmarshal(data, &transactions); err != nil {
		return transactions, err
	}

	return transactions, nil
}
