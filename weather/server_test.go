package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeatherSummary(t *testing.T) {
	server := &WeatherServer{}

	t.Run("return weahter summary", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/summary?lat=-13.3&lon=125", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("return weahter summary", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/summary", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "{}"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
