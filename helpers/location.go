package helpers

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func GetCurrentCity() (string, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var loc Location
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return "", err
	}

	return loc.City, nil
}
