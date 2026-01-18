package utils

// CelsiusToFahrenheit converts Celsius to Fahrenheit
// Formula: F = C * 1.8 + 32
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

// CelsiusToKelvin converts Celsius to Kelvin
// Formula: K = C + 273
func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
