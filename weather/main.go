package main

import (
	"log"
	"net/http"
)

func main() {
	server := &WeatherServer{}
	log.Fatal(http.ListenAndServe(":8000", server))
}
