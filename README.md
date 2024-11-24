# Microservices Weather API

This project implements a system to retrieve temperature data by CEP (Brazilian postal code). It consists of two microservices: **Service A** and **Service B**. The system is containerized using Docker and includes OpenTelemetry for distributed tracing with Zipkin.

## Overview

- **Service B**: Retrieves the city for a given CEP using the viaCEP API, fetches temperature data using the WeatherAPI, and returns the data in Celsius, Fahrenheit, and Kelvin.

## Requirements

- A valid WeatherAPI key ([get it here](https://www.weatherapi.com/))

## API Responses

### Service B

- **Successful Response**:
  - HTTP Status: `200`
  - Response:

    ```json
    {
      "city": "São Paulo",
      "temp_C": 28.5,
      "temp_F": 83.3,
      "temp_K": 301.65
    }
    ```

- **Invalid CEP (format)**:
  - HTTP Status: `422`
  - Response: `invalid zipcode`

- **Non-existent CEP**:
  - HTTP Status: `404`
  - Response: `can not find zipcode`

## Notes

- Ensure the WeatherAPI key is correctly set in the `docker-compose.yml`.
- The application is designed to work with valid Brazilian postal codes (8 digits).
