package main

import (
	"log"
	"net/http"
	"os"

	"servicea/internal/handlers"
	"servicea/internal/services"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Could not load .env file")
	}

	zipCodeService := services.NewZipCodeService()
	weatherService := services.NewWeatherAPIService(
		os.Getenv("WEATHER_API_SERIVCE_URL"),
	)

	weatherAPIHandler := &handlers.WeatherAPIHandler{
		ZipCodeService: zipCodeService,
		WeatherService: weatherService,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.Handle("/", weatherAPIHandler)
	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
