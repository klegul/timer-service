package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestPayload struct {
	Duration    int    `json:"duration"`     // Duration in seconds
	CallbackURL string `json:"callback_url"` // URL to call after the timer expires
}

func HandleTimerRequest(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	go startTimer(payload.Duration, payload.CallbackURL)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Timer started"))
}

func startTimer(duration int, callbackURL string) {
	time.Sleep(time.Duration(duration) * time.Second)

	resp, err := http.Post(callbackURL, "application/json", nil)
	if err != nil {
		fmt.Printf("Error calling callback URL: %v\n", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/start-timer", HandleTimerRequest)

	fmt.Println("Timer service is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
