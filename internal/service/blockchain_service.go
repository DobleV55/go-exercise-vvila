package service

import (
	"fmt"
	"go-exercise-vvila/internal/blockchain"
	"go-exercise-vvila/internal/cache"
	"time"
)

type BlockchainService struct {
	blockchainClient *blockchain.BlockchainClient
	cacheClient      cache.CacheClient
}

func NewBlockchainService(bc *blockchain.BlockchainClient, cc cache.CacheClient) *BlockchainService {
	return &BlockchainService{
		blockchainClient: bc,
		cacheClient:      cc,
	}
}

func (bps *BlockchainService) GetLastTradedPrices(pairs []string) ([]map[string]string, error) {
	cacheKey := fmt.Sprintf("bps:%d", time.Now().Unix()/60) // this way we cache the results for the current minute, different from storing the results for 60 seconds.
	if prices, found := bps.cacheClient.Get(cacheKey); found {
		return prices, nil
	}

	ticker, err := bps.blockchainClient.GetTicker()
	if err != nil {
		return nil, err
	}

	var results []map[string]string
	for _, pair := range pairs {
		var price float64
		switch pair {
		case "BTC/USD":
			price = ticker.USD.Last
		case "BTC/CHF":
			price = ticker.CHF.Last
		case "BTC/EUR":
			price = ticker.EUR.Last
		default:
			continue
		}
		result := map[string]string{
			"pair":   pair,
			"amount": fmt.Sprintf("%.2f", price),
		}
		results = append(results, result)
	}

	bps.cacheClient.Set(cacheKey, results, 60*time.Second)

	return results, nil
}
