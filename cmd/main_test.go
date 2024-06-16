package main

import (
	"os"
	"testing"

	"github.com/renatosaksanni/apf/config"
)

func TestMain(t *testing.T) {
	os.Setenv("ALPHA_VANTAGE_API_KEY", "test_key")
	err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
}
