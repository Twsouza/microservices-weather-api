package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"servicea/internal/dto"
	"servicea/internal/services"
	mockServices "servicea/mocks/services"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServeHTTP(t *testing.T) {
	t.Run("Valid CEP", func(t *testing.T) {
		mockAPI := mockServices.NewMockWeatherApiService(t)
		expectedRes := &dto.WeatherResponse{
			City:  "City",
			TempC: 25,
			TempF: 77,
			TempK: 298.15,
		}
		mockAPI.EXPECT().GetTemperaturesByZipCode(mock.Anything, "12345678").Return(expectedRes, nil)
		mockZip := mockServices.NewMockZipCodeService(t)
		mockZip.EXPECT().IsValidCEP(mock.Anything, "12345678").Return(true)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"cep": "12345678"}`))
		require.NoError(t, err)

		h := &WeatherAPIHandler{
			ZipCodeService: mockZip,
			WeatherService: mockAPI,
		}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, `{"city":"City","temp_C":25,"temp_F":77,"temp_K":298.15}`, w.Body.String())
	})

	t.Run("Invalid CEP", func(t *testing.T) {
		mockAPI := mockServices.NewMockWeatherApiService(t)
		mockZip := mockServices.NewMockZipCodeService(t)
		mockZip.EXPECT().IsValidCEP(mock.Anything, "invalid").Return(false)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"cep": "invalid"}`))
		require.NoError(t, err)

		h := &WeatherAPIHandler{
			ZipCodeService: mockZip,
			WeatherService: mockAPI,
		}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.Contains(t, w.Body.String(), "invalid zipcode")
	})

	t.Run("Method Not Allowed", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		h := &WeatherAPIHandler{}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusMethodNotAllowed, w.Code)
		require.Contains(t, w.Body.String(), "invalid method")
	})

	t.Run("Weather Service Invalid ZipCode", func(t *testing.T) {
		mockAPI := mockServices.NewMockWeatherApiService(t)
		mockAPI.EXPECT().GetTemperaturesByZipCode(mock.Anything, "invalidzip").Return(nil, services.InvalidZipCode)
		mockZip := mockServices.NewMockZipCodeService(t)
		mockZip.EXPECT().IsValidCEP(mock.Anything, "invalidzip").Return(true)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"cep": "invalidzip"}`))
		require.NoError(t, err)

		h := &WeatherAPIHandler{
			ZipCodeService: mockZip,
			WeatherService: mockAPI,
		}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.Contains(t, w.Body.String(), services.InvalidZipCode.Error())
	})

	t.Run("ZipCode Not Found", func(t *testing.T) {
		mockAPI := mockServices.NewMockWeatherApiService(t)
		mockAPI.EXPECT().GetTemperaturesByZipCode(mock.Anything, "99999999").Return(nil, services.ZipCodeNotFound)
		mockZip := mockServices.NewMockZipCodeService(t)
		mockZip.EXPECT().IsValidCEP(mock.Anything, "99999999").Return(true)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"cep": "99999999"}`))
		require.NoError(t, err)

		h := &WeatherAPIHandler{
			ZipCodeService: mockZip,
			WeatherService: mockAPI,
		}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusNotFound, w.Code)
		require.Contains(t, w.Body.String(), services.ZipCodeNotFound.Error())
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		mockAPI := mockServices.NewMockWeatherApiService(t)
		mockAPI.EXPECT().GetTemperaturesByZipCode(mock.Anything, "12345678").Return(nil, errors.New("some internal error"))
		mockZip := mockServices.NewMockZipCodeService(t)
		mockZip.EXPECT().IsValidCEP(mock.Anything, "12345678").Return(true)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"cep": "12345678"}`))
		require.NoError(t, err)

		h := &WeatherAPIHandler{
			ZipCodeService: mockZip,
			WeatherService: mockAPI,
		}
		h.ServeHTTP(w, r)

		require.Equal(t, http.StatusInternalServerError, w.Code)
		require.Contains(t, w.Body.String(), http.StatusText(http.StatusInternalServerError))
	})
}
