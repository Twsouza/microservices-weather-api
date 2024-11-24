package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"serviceb/tracer"

	"github.com/sirupsen/logrus"
)

var (
	InvalidZipCode  = errors.New("invalid zipcode")
	ZipCodeNotFound = errors.New("can not find zipcode")
)

// ZipCodeService defines the interface for fetching location by ZIP code.
type ZipCodeService interface {
	GetLocationByZipCode(ctx context.Context, zipCode string) (string, error)
}

// ViaCEPService implements ZipCodeService using the ViaCEP API.
type ViaCEPService struct {
	BaseURL string
}

func NewViaCEPService(baseURL string) *ViaCEPService {
	return &ViaCEPService{
		BaseURL: baseURL,
	}
}

// viaCEPResponse represents the response from the ViaCEP API.
type viaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

// GetLocationByZipCode fetches the city name for a given ZIP code.
func (v *ViaCEPService) GetLocationByZipCode(ctx context.Context, zipCode string) (string, error) {
	ctx, span := tracer.Tracer.Start(ctx, "ViaCEPService.GetLocationByZipCode")
	defer span.End()

	apiURL, err := url.Parse(v.BaseURL)
	if err != nil {
		return "", fmt.Errorf("GetLocationByZipCode: error parsing URL: %w", err)
	}
	apiURL.Path = fmt.Sprintf("/ws/%s/json", zipCode)

	logrus.WithFields(logrus.Fields{
		"apiURL": apiURL.String(),
	}).Info("Fetching location by ZIP code", zipCode)

	resp, err := http.Get(apiURL.String())
	if err != nil {
		return "", fmt.Errorf("GetLocationByZipCode: error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return "", fmt.Errorf("GetLocationByZipCode: %w", InvalidZipCode)
	}

	var data viaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("GetLocationByZipCode: error decoding response: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"data": data,
	}).Info("City found for ZIP code ", zipCode)

	if data.Erro != "" {
		return "", fmt.Errorf("GetLocationByZipCode: %w", ZipCodeNotFound)
	}

	return data.Localidade, nil
}
