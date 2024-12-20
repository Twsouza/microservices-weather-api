package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"servicea/internal/dto"
	"servicea/tracer"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	weatherApiServiceTimeout = 1 * time.Second
)

var (
	InvalidZipCode  = errors.New("invalid zipcode")
	ZipCodeNotFound = errors.New("can not find zipcode")
	UnexpectedResp  = errors.New("unexpected response")
)

type (
	WeatherApiService interface {
		// GetTemperaturesByZipCode fetches the celcius, fahrenheit and kelvin temperatures for a given ZIP code.
		GetTemperaturesByZipCode(ctx context.Context, zipCode string) (*dto.WeatherResponse, error)
	}
	WeatherAPIService struct {
		apiServiceURL string
		httpClient    *http.Client // Add httpClient field
	}
)

func NewWeatherAPIService(apiServiceURL string) WeatherApiService {
	return &WeatherAPIService{
		apiServiceURL: apiServiceURL,
		httpClient: &http.Client{
			Timeout:   weatherApiServiceTimeout,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
	}
}

func (w *WeatherAPIService) GetTemperaturesByZipCode(ctx context.Context, zipCode string) (*dto.WeatherResponse, error) {
	_, span := tracer.Tracer.Start(ctx, "WeatherAPIService.GetTemperaturesByZipCode")
	defer span.End()

	client := w.httpClient
	endpoint, err := url.Parse(w.apiServiceURL)
	if err != nil {
		return nil, fmt.Errorf("GetTemperaturesByZipCode: error parsing URL: %w", err)
	}
	endpoint.Path = "/weather"
	query := endpoint.Query()
	query.Set("cep", zipCode)
	endpoint.RawQuery = query.Encode()

	logrus.WithFields(logrus.Fields{
		"endpoint": endpoint.String(),
	}).Info("fetching temperature data ", zipCode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetTemperaturesByZipCode error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnprocessableEntity {
			return nil, fmt.Errorf("GetTemperaturesByZipCode: %w", InvalidZipCode)
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("GetTemperaturesByZipCode: %w", ZipCodeNotFound)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("GetTemperaturesByZipCode: error reading response body: %w", err)
		}

		logrus.WithFields(logrus.Fields{
			"statusCode": resp.StatusCode,
			"body":       string(body),
		}).Error("GetTemperaturesByZipCode unexpected response")

		return nil, fmt.Errorf("GetTemperaturesByZipCode: %w", UnexpectedResp)
	}

	res := &dto.WeatherResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("GetTemperaturesByZipCode error decoding response: %w", err)
	}

	return res, nil
}
