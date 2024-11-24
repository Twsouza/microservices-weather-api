package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"servicea/internal/dto"
	"servicea/internal/services"
	"servicea/tracer"

	"github.com/sirupsen/logrus"
)

var (
	InvalidZipCode = errors.New("invalid zipcode")
)

// WeatherAPIHandler handles weather-related HTTP requests.
type WeatherAPIHandler struct {
	ZipCodeService services.ZipCodeService
	WeatherService services.WeatherApiService
}

// ServeHTTP processes the HTTP request and returns the temperature data.
func (h *WeatherAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Tracer.Start(r.Context(), "WeatherAPIHandler.ServeHTTP")
	defer span.End()

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	req := dto.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !h.ZipCodeService.IsValidCEP(ctx, req.Cep) {
		http.Error(w, InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		return
	}

	wr, err := h.WeatherService.GetTemperaturesByZipCode(ctx, req.Cep)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request": req,
			"error":   err,
		}).Error("failed to get temperature data")

		if errors.Is(err, services.InvalidZipCode) {
			http.Error(w, services.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
			return
		}

		if errors.Is(err, services.ZipCodeNotFound) {
			http.Error(w, services.ZipCodeNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wr)
}
