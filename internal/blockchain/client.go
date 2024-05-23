package blockchain

import (
	"encoding/json"
	"fmt"
	"go-exercise-vvila/models"
	"net/http"
)

const baseUrl = "https://blockchain.info/ticker"

type BlockchainClient struct{}

func NewBlockchainClient() *BlockchainClient {
	return &BlockchainClient{}
}

func (bc *BlockchainClient) GetTicker() (models.TickerResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s?cors=true", baseUrl))
	if err != nil {
		return models.TickerResponse{}, err
	}
	defer resp.Body.Close()

	var tickerResponse models.TickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&tickerResponse); err != nil {
		return tickerResponse, err
	}

	return tickerResponse, nil
}
