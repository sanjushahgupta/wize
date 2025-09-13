package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"wize/internal/handler"
)

func main() {
	f, err := setupLogger()
	if err != nil {
		fmt.Println("failed to set up logger:", err)
		os.Exit(1)
	}
	defer f.Close()

	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/transactions", handler.AllTransactions)

	fmt.Println("listening on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("something went wrong while listening and serving api", err.Error())
	}

	fmt.Println("server has stopped")
}

func setupLogger() (*os.File, error) {
	const filename = "myserver.log"
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println("log file:", absPath)

	// Always append; create if missing
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	sysLogHandler := slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(sysLogHandler))

	log.SetOutput(slog.NewLogLogger(sysLogHandler, slog.LevelInfo).Writer())

	return f, nil
}
