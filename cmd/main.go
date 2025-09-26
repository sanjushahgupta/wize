package main

import (
	"log"
	"net/http"
	"wize/internal/domain"
	"wize/internal/handler"

	"github.com/spf13/viper"
)

func main() {
	var err error
	domain.CFG, err = loadConfig()
	if err != nil {
		log.Println("failed to load config", err.Error())
		return
	}

	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/transactions", handler.AllTransactions)

	log.Println("listening on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("something went wrong while listening and serving api", err.Error())
	}

	log.Println("server has stopped")
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
