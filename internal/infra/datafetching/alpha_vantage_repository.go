package datafetching

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/renatosaksanni/apf/internal/domain/datafetching"
)

type AlphaVantageRepository struct{}

func NewAlphaVantageRepository() *AlphaVantageRepository {
	return &AlphaVantageRepository{}
}

func (r *AlphaVantageRepository) FetchData(symbol string) ([]datafetching.Data, error) {
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=1min&apikey=%s", symbol, apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	var data []datafetching.Data
	for date, value := range result["Time Series (1min)"].(map[string]interface{}) {
		transactions := value.(map[string]interface{})["4. close"].(float64)
		data = append(data, datafetching.Data{Date: date, Transactions: transactions})
	}
	return data, nil
}
