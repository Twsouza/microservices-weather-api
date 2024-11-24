package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"serviceb/internal/dto"
	"serviceb/internal/services"

	"github.com/sirupsen/logrus"
)

// WeatherHandler handles weather-related HTTP requests.
type WeatherHandler struct {
	ZipCodeService services.ZipCodeService
	WeatherService services.WeatherService
}

// ServeHTTP processes the HTTP request and returns the temperature data.
func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	location, err := h.ZipCodeService.GetLocationByZipCode(cep)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cep":   cep,
			"error": err,
		}).Error("Failed to get location by ZIP code")

		if errors.Is(err, services.InvalidZipCode) {
			http.Error(w, services.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		} else if errors.Is(err, services.ZipCodeNotFound) {
			http.Error(w, services.ZipCodeNotFound.Error(), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	tempC, err := h.WeatherService.GetTemperatureByLocation(location)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tempF, tempK := h.WeatherService.CalculateTemperature(tempC)
	temp := dto.Temperature{
		City:       location,
		Celsius:    tempC,
		Fahrenheit: tempF,
		Kelvin:     tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temp)
}
