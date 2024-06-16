package services

import (
	"os"
	"testing"
)

func TestGenerateForecast(t *testing.T) {
	service := NewForecastingService()
	err := service.GenerateForecast("../../data/transaction_data.csv", "prophet", 30)
	if err != nil {
		t.Fatalf("Failed to generate forecast: %v", err)
	}

	if _, err := os.Stat("../../data/prophet_forecast.csv"); os.IsNotExist(err) {
		t.Fatalf("Forecast file not created")
	}
}
