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
	cacheClient := cache.NewCacheClient("redis:6379")
	krakenClient := kraken.NewKrakenClient()
	blockchainClient := blockchain.NewBlockchainClient()

	krakenService := service.NewKrakenService(krakenClient, cacheClient)
	blockchainService := service.NewBlockchainService(blockchainClient, cacheClient)

	services := []service.PriceServiceInterface{krakenService, blockchainService}
	averageService := service.NewAverageService(services)

	http.HandleFunc("/api/v1/ltp", api.GetLTPHandler(krakenService))
	http.HandleFunc("/api/v1/average", api.GetAveragePriceHandler(averageService)) // It might be slower, but it's more "accurate" when prices fluctuate a lot
	log.Fatal(http.ListenAndServe(":8080", nil))
}
