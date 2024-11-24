package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperaturesByZipCode(t *testing.T) {
	testCases := []struct {
		name           string
		zipCode        string
		responseStatus int
		responseBody   string
		expectedError  error
	}{
		{
			name:           "Successful response",
			zipCode:        "12345",
			responseStatus: http.StatusOK,
			responseBody:   `{"city":"City","temp_C":22.3,"temp_F":72.14,"temp_K":295.45}`,
			expectedError:  nil,
		},
		{
			name:           "Invalid zip code",
			zipCode:        "invalid",
			responseStatus: http.StatusUnprocessableEntity,
			responseBody:   `Invalid zip code`,
			expectedError:  InvalidZipCode,
		},
		{
			name:           "Zip code not found",
			zipCode:        "99999",
			responseStatus: http.StatusNotFound,
			responseBody:   `can not find zipcode`,
			expectedError:  ZipCodeNotFound,
		},
		{
			name:           "Unexpected status code",
			zipCode:        "12345",
			responseStatus: http.StatusInternalServerError,
			responseBody:   `Internal server error`,
			expectedError:  UnexpectedResp,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup a mock HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.responseStatus)
				w.Write([]byte(tc.responseBody))
			}))
			defer server.Close()

			// Create service instance with mock server URL and client
			service := &WeatherAPIService{
				apiServiceURL: server.URL,
				httpClient:    server.Client(),
			}

			// Call the method
			_, err := service.GetTemperaturesByZipCode(context.Background(), tc.zipCode)

			fmt.Printf("Unwrapped error: %v\n", errors.Unwrap(err))
			fmt.Printf("Expected error: %v\n", tc.expectedError)

			// Check for expected error
			if err != nil && tc.expectedError == nil {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectedError != nil {
				t.Errorf("expected error: %v, got nil", tc.expectedError)
			} else if err != nil && tc.expectedError != nil && !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
