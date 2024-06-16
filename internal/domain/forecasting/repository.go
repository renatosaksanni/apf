package forecasting

type ForecastRepository interface {
	Save(forecast []Forecast) error
	FindAll() ([]Forecast, error)
}
