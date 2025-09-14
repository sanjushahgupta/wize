package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"wize/internal/domain"
	"wize/internal/repository"
)

func AllTransactions(w http.ResponseWriter, r *http.Request) {
	ts := domain.NewTransactionsService(repository.NewWiseClient(domain.CFG.WiseAPIKey))

	userIDString := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))

		slog.Error(
			"failed to convert userID from string to int",
			slog.String("error_message", err.Error()),
			slog.String("user_id", userIDString),
		)
		return
	}

	trs, err := ts.GetAllTransactions(r.Context(), int64(userID))
	if err != nil {
		if errors.Is(err, domain.NoUserProvidedError) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			slog.Warn(
				"bad request",
				slog.Int("user", userID),
				slog.String("error", err.Error()),
			)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))

		slog.Error(
			"internal server error",
			slog.String("error_message", err.Error()),
			slog.Int("user", userID),
		)

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
