package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"serviceb/internal/dto"
	"serviceb/internal/services"
	"serviceb/tracer"

	"github.com/sirupsen/logrus"
)

// WeatherHandler handles weather-related HTTP requests.
type WeatherHandler struct {
	ZipCodeService services.ZipCodeService
	WeatherService services.WeatherService
}

// ServeHTTP processes the HTTP request and returns the temperature data.
func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Tracer.Start(r.Context(), "WeatherHandler.ServeHTTP")
	defer span.End()

	cep := r.URL.Query().Get("cep")
	location, err := h.ZipCodeService.GetLocationByZipCode(ctx, cep)
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

	tempC, err := h.WeatherService.GetTemperatureByLocation(ctx, location)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tempF, tempK := h.WeatherService.CalculateTemperature(ctx, tempC)
	temp := dto.Temperature{
		City:       location,
		Celsius:    tempC,
		Fahrenheit: tempF,
		Kelvin:     tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temp)
}
