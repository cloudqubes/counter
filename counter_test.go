package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountUpAndDown(t *testing.T) {
	t.Run("Count Up", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/count-up", nil)
		response := httptest.NewRecorder()

		countUp(response, request)

		got := response.Body.String()
		want := "1\n"

		if got != want {
			t.Errorf("Expecting %q. Got %q", want, got)
		}
	})

	t.Run("Count Up the second time", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/count-up", nil)
		response := httptest.NewRecorder()

		countUp(response, request)

		got := response.Body.String()
		want := "2\n"

		if got != want {
			t.Errorf("Expecting %q. Got %q", want, got)
		}
	})

	t.Run("Count down", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/count-down", nil)
		response := httptest.NewRecorder()

		countDown(response, request)

		got := response.Body.String()
		want := "1\n"

		if got != want {
			t.Errorf("Expecting %q. Got %q", want, got)
		}
	})

	t.Run("Count down again", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/count-down", nil)
		response := httptest.NewRecorder()

		countDown(response, request)

		got := response.Body.String()
		want := "0\n"

		if got != want {
			t.Errorf("Expecting %q. Got %q", want, got)
		}
	})
}
