package main

import (
	"fmt"
	"net/http"
	"wize/internal/handler"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/transactions", handler.AllTransactions)

	fmt.Println("listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("something went wrong while listening and serving api", err.Error())
	}

	fmt.Println("server has stopped")
}
