package api

import (
	"encoding/json"
	"go-exercise-vvila/internal/service"
	"net/http"
)

const pairs = "BTC/USD,BTC/CHF,BTC/EUR"

func GetLTPHandler(service *service.PriceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ltp, err := service.GetLastTradedPrices([]string{pairs})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"ltp": ltp})
	}
}
