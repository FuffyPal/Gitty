package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WebhookPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("webhook +")

	var payload WebhookPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "not read data", http.StatusBadRequest)
		return
	}

	fmt.Println("**" + payload.Title + "**:")
	fmt.Println("*" + payload.Description + "*")

	sendDiscord("**" + payload.Title + "**\n" + "*" + payload.Description + "*")

}

func sendDiscord(message string) error {
	webhookURL := os.Getenv("URL")

	payload := map[string]string{
		"content": message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("hello world")

	http.HandleFunc("/webhook", WebhookHandler)
	http.ListenAndServe(":"+port, nil)
}
