package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"fullcycle-lab-weather-api/internal/models"
)

type CEPService struct {
	viaCEPURL string
}

func NewCEPService() *CEPService {
	return &CEPService{
		viaCEPURL: "https://viacep.com.br/ws/%s/json/",
	}
}

// ValidateCEP validates if the CEP has the correct format (8 digits)
func (s *CEPService) ValidateCEP(cep string) bool {
	// Remove any non-digit characters
	re := regexp.MustCompile(`\D`)
	cleanCEP := re.ReplaceAllString(cep, "")

	// Check if it has exactly 8 digits
	return len(cleanCEP) == 8
}

// GetLocationByCEP fetches location data from viaCEP API
func (s *CEPService) GetLocationByCEP(cep string) (*models.ViaCEPResponse, error) {
	// Clean CEP (remove any non-digit characters)
	re := regexp.MustCompile(`\D`)
	cleanCEP := re.ReplaceAllString(cep, "")

	if len(cleanCEP) != 8 {
		return nil, fmt.Errorf("invalid zipcode")
	}

	url := fmt.Sprintf(s.viaCEPURL, cleanCEP)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CEP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid zipcode")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var cepResponse models.ViaCEPResponse
	if err := json.Unmarshal(body, &cepResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check if CEP was not found
	if cepResponse.Erro || cepResponse.Localidade == "" {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return &cepResponse, nil
}
