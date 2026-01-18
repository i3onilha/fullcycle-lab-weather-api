package utils

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "zero celsius",
			celsius:  0,
			expected: 32,
		},
		{
			name:     "28.5 celsius",
			celsius:  28.5,
			expected: 83.3, // 28.5 * 1.8 + 32 = 83.3
		},
		{
			name:     "100 celsius",
			celsius:  100,
			expected: 212, // 100 * 1.8 + 32 = 212
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tt.celsius)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.celsius, result, tt.expected)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "zero celsius",
			celsius:  0,
			expected: 273,
		},
		{
			name:     "28.5 celsius",
			celsius:  28.5,
			expected: 301.5,
		},
		{
			name:     "100 celsius",
			celsius:  100,
			expected: 373,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToKelvin(tt.celsius)
			if result != tt.expected {
				t.Errorf("CelsiusToKelvin(%v) = %v, want %v", tt.celsius, result, tt.expected)
			}
		})
	}
}
