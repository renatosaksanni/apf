package services

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/renatosaksanni/apf/internal/domain/datafetching"
	datafetchingInfra "github.com/renatosaksanni/apf/internal/infra/datafetching"
)

type DataFetchingService struct {
	repository datafetching.DataRepository
}

func NewDataFetchingService() *DataFetchingService {
	repository := datafetchingInfra.NewAlphaVantageRepository()
	return &DataFetchingService{repository: repository}
}

func (s *DataFetchingService) FetchAndSaveData(symbol string) error {
	data, err := s.repository.FetchData(symbol)
	if err != nil {
		return err
	}

	file, err := os.Create("data/real_time_data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, d := range data {
		if err := writer.Write([]string{d.Date, fmt.Sprintf("%f", d.Transactions)}); err != nil {
			return err
		}
	}
	return nil
}
