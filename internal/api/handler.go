package api

import (
	"encoding/json"
	"go-exercise-vvila/internal/service"
	"net/http"
	"strings"
)

const pairs = "BTC/USD,BTC/CHF,BTC/EUR"

func GetLTPHandler(krakenService *service.KrakenService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tradedPairs := strings.Split(pairs, ",")
		ltp, err := krakenService.GetLastTradedPrices(tradedPairs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"ltp": ltp})
	}
}

func GetAveragePriceHandler(service *service.AverageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// from string to []string separated by commas
		averagePairs := strings.Split(pairs, ",")
		averagePrices, err := service.GetAveragePrice(averagePairs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"average_prices": averagePrices})
	}
}
