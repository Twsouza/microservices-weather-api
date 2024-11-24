package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetLocationByZipCode tests the GetLocationByZipCode method.
func TestGetLocationByZipCode(t *testing.T) {
	// Mock the HTTP response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"localidade": "Test City"}`))
	}))
	defer ts.Close()

	service := NewViaCEPService(ts.URL)

	location, err := service.GetLocationByZipCode(context.Background(), "12345678")
	if err != nil {
		t.Fatalf("GetLocationByZipCode returned error: %v", err)
	}

	if location != "Test City" {
		t.Errorf("GetLocationByZipCode = %s; want %s", location, "Test City")
	}
}
