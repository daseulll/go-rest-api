package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

const jsonContentType = "application/json"

type WeatherServer struct {
	weatherService WeatherService
}

func (s *WeatherServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "")

	if path == "/summary" {
		queryParams := r.URL.Query()
		lat_param := queryParams.Get("lat")
		lon_param := queryParams.Get("lon")

		lat, _ := strconv.ParseFloat(lat_param, 64)
		lon, _ := strconv.ParseFloat(lon_param, 64)

		summary := s.weatherService.GetWeatherSummary(lat, lon)

		w.Header().Set("content-type", jsonContentType)
		json.NewEncoder(w).Encode(summary)
	}
}
