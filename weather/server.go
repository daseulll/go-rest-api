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
		latParam := queryParams.Get("lat")
		lonParam := queryParams.Get("lon")

		lat, latErr := strconv.ParseFloat(latParam, 64)
		lon, lonErr := strconv.ParseFloat(lonParam, 64)

		if latErr != nil || lonErr != nil {
			http.Error(w, "lat and lon are required.", http.StatusBadRequest)
			return
		}

		if !isValidCoordinate(lat, -90, 90) || !isValidCoordinate(lon, -180, 180) {
			http.Error(w, "lat must be between -90 and 90, lon must be between -180 and 180.", http.StatusBadRequest)
			return
		}

		summary := s.weatherService.GetWeatherSummary(lat, lon)

		w.Header().Set("content-type", jsonContentType)
		json.NewEncoder(w).Encode(summary)
	}
}

func isValidCoordinate(coord float64, min, max float64) bool {
	return coord >= min && coord <= max
}
