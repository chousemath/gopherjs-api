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
	Day             string  `json:"day,omitempty"`
	TemperatureLow  float32 `json:"temperatureLow,omitempty"`
	TemperatureHigh float32 `json:"temperatureHigh,omitempty"`
	Precipitation   float32 `json:"precipitation,omitempty"`
	Humidity        float32 `json:"humidity,omitempty"`
	Wind            float32 `json:"wind,omitempty"`
}

// WeatherReports represents a typical weather report
type WeatherReports struct {
	Reports []WeatherReport `json:"reports,omitempty"`
}

var weatherReports = WeatherReports{
	Reports: []WeatherReport{
		WeatherReport{
			ID:              1,
			Day:             "Monday",
			TemperatureLow:  10.5,
			TemperatureHigh: 52.1,
			Precipitation:   32.1,
			Humidity:        11.4,
			Wind:            23.3,
		},
		WeatherReport{
			ID:              2,
			Day:             "Tuesday",
			TemperatureLow:  40.5,
			TemperatureHigh: 28.1,
			Precipitation:   92.1,
			Humidity:        12.4,
			Wind:            52.3,
		},
		WeatherReport{
			ID:              3,
			Day:             "Wednesday",
			TemperatureLow:  37.5,
			TemperatureHigh: 420.1,
			Precipitation:   19.1,
			Humidity:        77.4,
			Wind:            55.3,
		},
		WeatherReport{
			ID:              4,
			Day:             "Thursday",
			TemperatureLow:  15.5,
			TemperatureHigh: 76.1,
			Precipitation:   43.1,
			Humidity:        91.4,
			Wind:            69.3,
		},
		WeatherReport{
			ID:              5,
			Day:             "Friday",
			TemperatureLow:  1.5,
			TemperatureHigh: 5.1,
			Precipitation:   2.1,
			Humidity:        1.4,
			Wind:            3.3,
		},
		WeatherReport{
			ID:              6,
			Day:             "Saturday",
			TemperatureLow:  102.5,
			TemperatureHigh: 352.1,
			Precipitation:   432.1,
			Humidity:        711.4,
			Wind:            823.3,
		},
		WeatherReport{
			ID:              7,
			Day:             "Sunday",
			TemperatureLow:  90.5,
			TemperatureHigh: 82.1,
			Precipitation:   72.1,
			Humidity:        61.4,
			Wind:            53.3,
		},
	},
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
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(WeatherReport{
		ID:              1,
		Day:             "Monday",
		TemperatureLow:  10.5,
		TemperatureHigh: 52.1,
		Precipitation:   32.1,
		Humidity:        11.4,
		Wind:            23.3,
	})
}

// GetWeatherReports returns a bunch of fake weather reports
func GetWeatherReports(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherReports)
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
	router.HandleFunc("/weatherreports", GetWeatherReports).Methods("GET")
	// The below four lines are used for deployment on Heroku
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
