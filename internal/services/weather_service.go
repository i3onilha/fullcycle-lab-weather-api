package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type WeatherService struct {
	weatherAPIKey string
	weatherAPIURL string
}

func NewWeatherService() *WeatherService {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		apiKey = "d6ba1be8b8f94d50b21210624261801" // Default for testing, should be set via env var
	}

	return &WeatherService{
		weatherAPIKey: apiKey,
		weatherAPIURL: "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
	}
}

// GetTemperatureByLocation fetches current temperature from WeatherAPI
func (s *WeatherService) GetTemperatureByLocation(city, state string) (float64, error) {
	// Format query as "city, state" for better location matching
	query := fmt.Sprintf("%s,%s,Brazil", city, state)
	query = strings.ReplaceAll(query, " ", "%20")

	url := fmt.Sprintf(s.weatherAPIURL, s.weatherAPIKey, query)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("weather API error: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response: %w", err)
	}

	var weatherResp struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}

	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return 0, fmt.Errorf("failed to parse weather response: %w", err)
	}

	return weatherResp.Current.TempC, nil
}
