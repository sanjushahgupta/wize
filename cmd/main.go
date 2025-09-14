package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"wize/internal/domain"
	"wize/internal/handler"
)

func main() {
	var err error
	domain.CFG, err = loadConfig()
	if err != nil {
		fmt.Println("failed to load config", err.Error())
		return
	}

	f, err := setupLogger(domain.CFG.LogFile)
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

func setupLogger(fileName string) (*os.File, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}
	fmt.Println("log file:", absPath)

	// Always append; create if missing
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

func loadConfig() (domain.Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return domain.CFG, err
	}

	domain.CFG.LogFile = v.GetString("log_file")
	domain.CFG.WiseAPIKey = v.GetString("wise_api_key")

	return domain.CFG, nil
}
