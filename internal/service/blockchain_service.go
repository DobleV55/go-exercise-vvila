package service

import (
	"fmt"
	"go-exercise-vvila/internal/blockchain"
)

type BlockchainPriceService struct {
	client *blockchain.BlockchainClient
}

func NewBlockchainPriceService(client *blockchain.BlockchainClient) *BlockchainPriceService {
	return &BlockchainPriceService{
		client: client,
	}
}

func (bps *BlockchainPriceService) GetLastTradedPrices(pairs []string) ([]map[string]string, error) {
	ticker, err := bps.client.GetTicker()
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

	return results, nil
}
