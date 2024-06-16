package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("ALPHA_VANTAGE_API_KEY", "test_key")
	err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if AppConfig.AlphaVantageAPIKey != "test_key" {
		t.Fatalf("Expected 'test_key', got %s", AppConfig.AlphaVantageAPIKey)
	}
}
