package api

import (
	"net/http"

	"github.com/renatosaksanni/apf/internal/app/services"
)

func ForecastHandler(w http.ResponseWriter, r *http.Request) {
	service := services.NewForecastingService()
	err := service.GenerateForecast("data/transaction_data.csv", "prophet", 30)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
