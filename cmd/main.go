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
	priceService := service.NewPriceService(krakenClient, redisClient)

	http.HandleFunc("/api/v1/ltp", api.GetLTPHandler(priceService))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
