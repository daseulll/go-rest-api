package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetWeatherSummary(t *testing.T) {
	server := &WeatherServer{}

	t.Run("return weahter summary", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/summary?lat=-13.3&lon=125", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got Summary
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into Summary struct, '%v'", response.Body, err)
		}

		want := Summary{
			SummaryStatement{
				"폭우가 내리고 있어요.",
				"어제보다 4도 가량 춥습니다. 최고기온은 10도 최저기온은 3도 입니다.",
				"내일 폭설이 내릴 수도 있으니 외출 시 주의하세요.",
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		if response.Code != http.StatusOK {
			t.Errorf("got %q, want %q", response.Code, http.StatusOK)
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
