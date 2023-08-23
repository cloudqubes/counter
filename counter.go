package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Global variable for counter
var count int

func countUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count = count + 1
	json.NewEncoder(w).Encode(count)
}

func countDown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count = count - 1
	json.NewEncoder(w).Encode(count)
}

func main() {
	http.HandleFunc("/count-up", countUp)
	http.HandleFunc("/count-down", countDown)

	fmt.Printf("Starting server on port 9000\n")
	count = 0
	fmt.Printf("Counter :%d\n", count)
	err := http.ListenAndServe(":9000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Shutting down server\n")
	} else if err != nil {
		fmt.Printf("Cannot start server :%s\n", err)
		os.Exit(1)
	}
}
