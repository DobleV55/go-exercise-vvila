package main

import (
	"log"
	"net/http"

	"go-exercise-vvila/internal/api"
	"go-exercise-vvila/internal/cache"
	"go-exercise-vvila/internal/kraken"
	"go-exercise-vvila/internal/service"
)

func main() {
	redisClient := cache.NewRedisClient("localhost:6379")
	krakenClient := kraken.NewKrakenClient()
	krakenService := service.NewPriceService(krakenClient, redisClient)

	services := []service.PriceServiceInterface{krakenService}
	averageService := service.NewAverageService(services)

	http.HandleFunc("/api/v1/ltp", api.GetLTPHandler(krakenService))
	http.HandleFunc("/api/v1/average", api.GetAveragePriceHandler(averageService)) // It might be slower, but it's more "accurate" when prices fluctuate a lot.s
	log.Fatal(http.ListenAndServe(":8080", nil))
}
