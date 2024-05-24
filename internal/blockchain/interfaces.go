package blockchain

import "go-exercise-vvila/models"

type BlockchainTickerClient interface {
	GetTicker(pair string) (models.TickerResponse, error)
}
