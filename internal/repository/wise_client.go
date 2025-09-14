package repository

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"wize/internal/domain"
)

type WiseClient struct {
	apiKey string
}

type wiseTransfer struct {
	User           int     `json:"user"`
	TargetAccount  int64   `json:"targetAccount"`
	SourceAccount  int64   `json:"sourceAccount"`
	Reference      string  `json:"reference"`
	Rate           float64 `json:"rate"`
	Created        string  `json:"created"`
	SourceCurrency string  `json:"sourceCurrency"`
	SourceValue    float64 `json:"sourceValue"`
	TargetCurrency string  `json:"targetCurrency"`
	TargetValue    float64 `json:"targetValue"`
}

const wiseTransfersEndPoint = "https://api.sandbox.transferwise.tech/v1/transfers"

func NewWiseClient(apiKey string) *WiseClient {
	return &WiseClient{
		apiKey: apiKey,
	}
}

func (c *WiseClient) SelectAll(ctx context.Context) ([]domain.Transaction, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", wiseTransfersEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	httpClient := http.Client{Timeout: time.Second * 10}
	httpResponse, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return []domain.Transaction{}, errors.New("unexpected status code while fetching transaction from wise")
	}

	var listOfWiseTransfers []wiseTransfer
	err = json.NewDecoder(httpResponse.Body).Decode(&listOfWiseTransfers)
	if err != nil {
		return []domain.Transaction{}, err
	}

	transactions := make([]domain.Transaction, len(listOfWiseTransfers))
	for i, t := range listOfWiseTransfers {
		transactions[i] = domain.Transaction{
			Date:        t.Created,
			Amount:      t.SourceValue,
			Description: t.Reference,
			From:        t.SourceAccount,
			To:          t.TargetAccount,
			Category:    "",
		}
	}

	return transactions, nil
}
