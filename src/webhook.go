package src

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	w.WriteHeader(http.StatusOK)
}
