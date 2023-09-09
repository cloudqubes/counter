package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

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
