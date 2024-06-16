package datafetching

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	timeSeries, ok := result["Time Series (1min)"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected data format")
	}

	var data []datafetching.Data
	for date, value := range timeSeries {
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected data format for value map")
		}
		closeValueStr, ok := valueMap["4. close"].(string)
		if !ok {
			return nil, fmt.Errorf("unexpected data format for close value")
		}
		closeValue, err := strconv.ParseFloat(closeValueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing close value: %v", err)
		}
		data = append(data, datafetching.Data{Date: date, Transactions: closeValue})
	}
	return data, nil
}
