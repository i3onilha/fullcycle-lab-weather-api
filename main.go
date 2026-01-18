package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fullcycle-lab-weather-api/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	weatherHandler := handlers.NewWeatherHandler()

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Weather endpoint
	router.HandleFunc("/weather/{cep}", weatherHandler.GetWeather).Methods("GET")

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
