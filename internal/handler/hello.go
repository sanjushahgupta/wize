package handler

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		log.Println("something went wrong when trying to write in api response", err.Error())
	}
}
