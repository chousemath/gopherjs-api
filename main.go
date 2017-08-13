package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// WeatherReport represents a typical weather report
type WeatherReport struct {
	ID              int32   `json:"id,omitempty"`
	TemperatureLow  float32 `json:"temperatureLow,omitempty"`
	TemperatureHigh float32 `json:"temperatureHigh,omitempty"`
	Precipitation   float32 `json:"precipitation,omitempty"`
	Humidity        float32 `json:"humidity,omitempty"`
	Wind            float32 `json:"wind,omitempty"`
}

var weather1 = WeatherReport{
	ID:              1,
	TemperatureLow:  10.5,
	TemperatureHigh: 42.1,
	Precipitation:   32.1,
	Humidity:        10.4,
	Wind:            2.3,
}

// This func figures out what address to listen on for traffic
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// GetWeatherReport returns the current weather forecast
func GetWeatherReport(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather1)
		return
	}
	json.NewEncoder(w).Encode(WeatherReport{})
}

func main() {
	// The below four lines are used for deployment on Heroku
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	// set up the router
	router := mux.NewRouter()
	// set up routes
	router.HandleFunc("/weatherreport", GetWeatherReport).Methods("GET")
	// The below four lines are used for deployment on Heroku
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
