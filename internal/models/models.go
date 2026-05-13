package models

import (
	"encoding/json"
	"strings"
)

// ViaCEPErro decodes viaCEP's "erro" field, which is usually a JSON boolean but
// may appear as the string "true" in some responses or proxies.
type ViaCEPErro bool

func (e *ViaCEPErro) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch t := v.(type) {
	case bool:
		*e = ViaCEPErro(t)
	case string:
		*e = ViaCEPErro(strings.EqualFold(strings.TrimSpace(t), "true"))
	case float64:
		*e = ViaCEPErro(t != 0)
	default:
		*e = false
	}
	return nil
}

// WeatherResponse represents the response with temperatures in different scales
type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// ViaCEPResponse represents the response from viaCEP API
type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
	Erro        ViaCEPErro `json:"erro"`
}

// WeatherAPIResponse represents the response from WeatherAPI
type WeatherAPIResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}
