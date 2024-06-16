package cli

import (
	"flag"
	"fmt"

	"github.com/renatosaksanni/apf/internal/app/services"
)

func Run(args []string) error {
	var (
		fetch    = flag.Bool("fetch", false, "Fetch real-time data using Alpha Vantage")
		symbol   = flag.String("symbol", "AAPL", "Symbol for fetching real-time data")
		model    = flag.String("model", "prophet", "Choose forecasting model: prophet or garch")
		dataPath = flag.String("data", "data/transaction_data.csv", "Path to the transaction data file")
		periods  = flag.Int("periods", 30, "Number of periods to forecast")
	)
	flag.Parse()

	if *fetch {
		dataService := services.NewDataFetchingService()
		err := dataService.FetchAndSaveData(*symbol)
		if err != nil {
			return err
		}
		*dataPath = "data/real_time_data.csv"
	}

	forecastService := services.NewForecastingService()
	err := forecastService.GenerateForecast(*dataPath, *model, *periods)
	if err != nil {
		return err
	}

	fmt.Println("Forecasting complete. Results saved.")
	return nil
}
