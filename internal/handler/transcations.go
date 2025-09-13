package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"wize/internal/domain"
	"wize/internal/repository"
)

func AllTransactions(w http.ResponseWriter, r *http.Request) {
	ts := domain.NewTransactionsService(repository.NewFileStorage())

	user := r.URL.Query().Get("user")

	trs, err := ts.GetAllTransactions(r.Context(), user)
	if err != nil {
		if errors.Is(err, domain.NoUserProvidedError) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	allTrns, err := json.Marshal(trs)
	if err != nil {
		fmt.Println("something went wrong while converting all transactions to byte", err)
	}

	_, err = w.Write(allTrns)
	if err != nil {
		fmt.Println("something went wrong while writing response", err)
	}
}
