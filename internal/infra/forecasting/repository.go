package forecasting

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/renatosaksanni/apf/internal/domain/forecasting"
)

type FileForecastRepository struct {
	filePath string
}

func NewFileForecastRepository(filePath string) *FileForecastRepository {
	return &FileForecastRepository{filePath: filePath}
}

func (r *FileForecastRepository) Save(forecast []forecasting.Forecast) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, f := range forecast {
		if err := writer.Write([]string{f.Date, fmt.Sprintf("%f", f.Value)}); err != nil {
			return err
		}
	}
	return nil
}

func (r *FileForecastRepository) FindAll() ([]forecasting.Forecast, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var forecasts []forecasting.Forecast
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		forecasts = append(forecasts, forecasting.Forecast{Date: record[0], Value: value})
	}
	return forecasts, nil
}
