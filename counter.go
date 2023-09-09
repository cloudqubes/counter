package main

import (
	"encoding/json"
	"net/http"
)

// Global variable for counter
var count int

func countUp(w http.ResponseWriter, r *http.Request) {
	count = count + 1
	// fmt.Printf("Counter %q", count)
	json.NewEncoder(w).Encode(count)
}

func countDown(w http.ResponseWriter, r *http.Request) {
	count = count - 1
	// fmt.Printf("Counter %q", count)
	json.NewEncoder(w).Encode(count)
}
