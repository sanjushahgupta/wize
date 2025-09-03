package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wize/internal/domain"
)

func AllTransactions(w http.ResponseWriter, r *http.Request) {
	trs := domain.GetAllTransactions()

	allTrns, err := json.Marshal(trs)
	if err != nil {
		fmt.Println("something went wrong while converting all transactions to byte", err)
	}

	_, err = w.Write(allTrns)
	if err != nil {
		fmt.Println("something went wrong while writing response", err)
	}
}
