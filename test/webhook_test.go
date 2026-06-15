package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"gitty/src"

	"github.com/joho/godotenv"
)

func TestWebhook(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("hello world")

	http.HandleFunc("/webhook", src.WebhookHandler)

	go http.ListenAndServe(":"+port, nil)
}
