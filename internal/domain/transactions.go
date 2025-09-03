package domain

type Transaction struct {
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

func GetAllTransactions() []Transaction {
	return []Transaction{
		{Date: "Jan 1", Amount: "$200", Description: "Rent"},
		{Date: "Jan 2", Amount: "$12", Description: "Coffee"},
		{Date: "Jan 3", Amount: "$30", Description: "Lunch"},
	}
}
