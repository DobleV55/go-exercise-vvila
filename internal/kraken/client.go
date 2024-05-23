package kraken

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://api.kraken.com/0/public/Ticker?pair=%s"

type KrakenClient struct{}

func NewKrakenClient() *KrakenClient {
	return &KrakenClient{}
}

func (kc *KrakenClient) GetTicker(pair string) (KrakenTickerResponse, error) {
	url := fmt.Sprintf(baseUrl, pair)
	resp, err := http.Get(url)
	if err != nil {
		return KrakenTickerResponse{}, err
	}
	defer resp.Body.Close()

	var tickerResp KrakenTickerResponse
	err = json.NewDecoder(resp.Body).Decode(&tickerResp)
	if err != nil {
		return KrakenTickerResponse{}, err
	}
	return tickerResp, nil
}
