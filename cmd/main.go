package main

import (
	"go-exercise-vvila/internal/blockchain"
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
	blockchainClient := blockchain.NewBlockchainClient()

	krakenService := service.NewPriceService(krakenClient, redisClient)
	blockchainService := service.NewBlockchainPriceService(blockchainClient)

	services := []service.PriceServiceInterface{krakenService, blockchainService}
	averageService := service.NewAverageService(services)

	http.HandleFunc("/api/v1/ltp", api.GetLTPHandler(krakenService))
	http.HandleFunc("/api/v1/average", api.GetAveragePriceHandler(averageService)) // It might be slower, but it's more "accurate" when prices fluctuate a lot
	log.Fatal(http.ListenAndServe(":8080", nil))
}
