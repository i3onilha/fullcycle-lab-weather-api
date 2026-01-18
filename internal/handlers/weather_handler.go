package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"fullcycle-lab-weather-api/internal/models"
	"fullcycle-lab-weather-api/internal/services"
	"fullcycle-lab-weather-api/internal/utils"
)

type WeatherHandler struct {
	cepService     *services.CEPService
	weatherService *services.WeatherService
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{
		cepService:     services.NewCEPService(),
		weatherService: services.NewWeatherService(),
	}
}

// GetWeather handles GET /weather/:cep
func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	// Get CEP from URL path using mux vars
	vars := mux.Vars(r)
	cep := vars["cep"]
	
	// Validate CEP format
	if !h.cepService.ValidateCEP(cep) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "invalid zipcode",
		})
		return
	}

	// Get location from viaCEP
	location, err := h.cepService.GetLocationByCEP(cep)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Message: "invalid zipcode",
			})
			return
		}
		if err.Error() == "can not find zipcode" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Message: "can not find zipcode",
			})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get temperature from WeatherAPI
	tempC, err := h.weatherService.GetTemperatureByLocation(location.Localidade, location.UF)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert temperatures
	tempF := utils.CelsiusToFahrenheit(tempC)
	tempK := utils.CelsiusToKelvin(tempC)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	})
}
