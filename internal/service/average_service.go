package service

import (
	"errors"
	"strconv"
)

type AverageService struct {
	services []PriceServiceInterface
}

func NewAverageService(services []PriceServiceInterface) *AverageService {
	return &AverageService{services: services}
}

func (as *AverageService) GetAveragePrice(pairs []string) (map[string]float64, error) {
	priceSum := make(map[string]float64)
	priceCount := make(map[string]int)

	// I would like to make this asynchronous, because if two of this services takes 2 seconds to respond, the total time would be 4 seconds.
	for _, service := range as.services {
		prices, err := service.GetLastTradedPrices(pairs)
		if err != nil {
			return nil, err
		}

		for _, price := range prices {
			pair := price["pair"]
			amount, err := strconv.ParseFloat(price["amount"], 64)
			if err != nil {
				return nil, err
			}
			priceSum[pair] += amount
			priceCount[pair]++
		}
	}

	averagePrice := make(map[string]float64)
	for pair, sum := range priceSum {
		if count, ok := priceCount[pair]; ok && count > 0 {
			averagePrice[pair] = sum / float64(count)
		} else {
			return nil, errors.New("no prices found for pair " + pair)
		}
	}

	return averagePrice, nil
}
