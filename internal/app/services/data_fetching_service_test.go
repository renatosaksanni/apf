package services

import (
	"os"
	"testing"
)

func TestFetchAndSaveData(t *testing.T) {
	os.Setenv("ALPHA_VANTAGE_API_KEY", "test_key")
	service := NewDataFetchingService()
	err := service.FetchAndSaveData("AAPL")
	if err != nil {
		t.Fatalf("Failed to fetch and save data: %v", err)
	}

	if _, err := os.Stat("../../data/real_time_data.csv"); os.IsNotExist(err) {
		t.Fatalf("Data file not created")
	}
}
