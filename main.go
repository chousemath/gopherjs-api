package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Weather represents a typical weather report
type Weather struct {
	TemperatureLow  float32 `json:"temperatureLow,omitempty"`
	TemperatureHigh float32 `json:"temperatureHigh,omitempty"`
	Precipitation   float32 `json:"precipitation,omitempty"`
	Humidity        float32 `json:"humidity,omitempty"`
	Wind            float32 `json:"wind,omitempty"`
}

var weather1 = Weather{
	TemperatureLow:  10.5,
	TemperatureHigh: 42.1,
	Precipitation:   32.1,
	Humidity:        10.4,
	Wind:            2.3,
}

func determineListenAddress() (string, error) {
	// determine which address to listen for traffic
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// GetWeather returns the current weather forecast
func GetWeather(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(weather1)
}

func main() {
	// attempt to get the address from Heroku
	addr, err := determineListenAddress()
	// if there is an error getting the address, panic
	if err != nil {
		log.Fatal(err)
	}
	// set up the router
	router := mux.NewRouter()
	// set up routes
	router.HandleFunc("/weather", GetWeather).Methods("GET")
	// spin up a server, panic if there is an error starting the server
	log.Printf("Listening on PORT:%s...", addr)
	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
