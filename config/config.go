package config

import (
	"fmt"
	"os"
)

type Config struct {
	AlphaVantageAPIKey string
}

var AppConfig Config

func LoadConfig() error {
	AppConfig.AlphaVantageAPIKey = os.Getenv("ALPHA_VANTAGE_API_KEY")
	if AppConfig.AlphaVantageAPIKey == "" {
		return fmt.Errorf("missing ALPHA_VANTAGE_API_KEY")
	}
	return nil
}
