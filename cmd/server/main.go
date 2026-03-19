package main

import (
	"log"
	"net/http"

	"github.com/xcentralnn/Alertless/internal/handler"
)

func main() {
	http.HandleFunc("/webhook", handler.HandleWebhook)

	log.Println("Alertless running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}