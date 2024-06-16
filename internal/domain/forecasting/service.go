package forecasting

type ForecastService interface {
	GenerateForecast(dataPath string, model string, periods int) error
	VisualizeForecast(model string) error
}
