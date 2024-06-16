package services

import (
	"fmt"
	"os/exec"

	"github.com/renatosaksanni/apf/internal/domain/forecasting"
	forecastingInfra "github.com/renatosaksanni/apf/internal/infra/forecasting"
)

type ForecastingService struct {
	repository forecasting.ForecastRepository
}

func NewForecastingService() *ForecastingService {
	repository := forecastingInfra.NewFileForecastRepository("data/forecast.csv")
	return &ForecastingService{repository: repository}
}

func (s *ForecastingService) GenerateForecast(dataPath, model string, periods int) error {
	var cmd *exec.Cmd

	if model == "prophet" {
		cmd = exec.Command("python3", "models/prophet_model.py", dataPath, fmt.Sprintf("%d", periods))
	} else if model == "garch" {
		cmd = exec.Command("python3", "models/garch_model.py", dataPath, fmt.Sprintf("%d", periods))
	} else {
		return fmt.Errorf("unknown model specified: use 'prophet' or 'garch'")
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	// Optionally, load and save forecast results into the repository
	// Assuming the forecast results are saved in "data/forecast.csv"
	forecasts, err := s.repository.FindAll()
	if err != nil {
		return err
	}

	return s.repository.Save(forecasts)
}

func (s *ForecastingService) VisualizeForecast(model string) error {
	cmd := exec.Command("python3", "models/visualize_forecast.py", model)
	return cmd.Run()
}
