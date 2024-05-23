package service

import (
	"fmt"
	"go-exercise-vvila/internal/cache"
	"go-exercise-vvila/internal/kraken"
	"time"
)

type PriceProvider interface {
	GetLastTradedPrice(pair string) (float64, error)
}

type PriceService struct {
	krakenClient kraken.TickerClient
	cacheClient  cache.CacheClient
}

func NewPriceService(kc kraken.TickerClient, cc cache.CacheClient) *PriceService {
	return &PriceService{
		krakenClient: kc,
		cacheClient:  cc,
	}
}

func (ps *PriceService) GetLastTradedPrices(pairs []string) ([]map[string]string, error) {

	cacheKey := fmt.Sprintf("ltp:%d", time.Now().Unix()/60) // this way we cache the results for the current minute, different from storing the results for 60 seconds.
	if prices, found := ps.cacheClient.Get(cacheKey); found {
		return prices, nil
	}

	var results []map[string]string
	for _, pair := range pairs {
		ticker, err := ps.krakenClient.GetTicker(pair)
		if err != nil {
			return nil, err
		}

		for k, v := range ticker.Result {
			result := map[string]string{
				"pair":   k,
				"amount": v.LastTradeClosed[0],
			}
			results = append(results, result)
		}
	}

	ps.cacheClient.Set(cacheKey, results, 60*time.Second)

	return results, nil
}
