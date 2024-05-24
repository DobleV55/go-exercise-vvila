package service

import (
	"fmt"
	"go-exercise-vvila/internal/cache"
	"go-exercise-vvila/internal/kraken"
	"time"
)

type KrakenService struct {
	krakenClient kraken.KrakenClientInterface
	cacheClient  cache.CacheClient
}

func NewKrakenService(kc kraken.KrakenClientInterface, cc cache.CacheClient) *KrakenService {
	return &KrakenService{
		krakenClient: kc,
		cacheClient:  cc,
	}
}

func (kps *KrakenService) GetLastTradedPrices(pairs []string) ([]map[string]string, error) {

	cacheKey := fmt.Sprintf("kps:%d", time.Now().Unix()/60) // this way we cache the results for the current minute, different from storing the results for 60 seconds.
	if prices, found := kps.cacheClient.Get(cacheKey); found {
		return prices, nil
	}

	var results []map[string]string
	for _, pair := range pairs {
		ticker, err := kps.krakenClient.GetTicker(pair)
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

	kps.cacheClient.Set(cacheKey, results, 60*time.Second) // the ttl is just for cleaning up the cache

	return results, nil
}
