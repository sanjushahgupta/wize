package domain

type Transaction struct {
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
	Category    string `json:"category"`
}
