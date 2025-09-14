package domain

type Transaction struct {
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	From        int64   `json:"from"`
	To          int64   `json:"to"`
	Category    string  `json:"category"`
}
