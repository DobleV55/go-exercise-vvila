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

func (kc *KrakenClient) GetTicker(pair string) (TickerResponse, error) {
	url := fmt.Sprintf(baseUrl, pair)
	resp, err := http.Get(url)
	if err != nil {
		return TickerResponse{}, err
	}
	defer resp.Body.Close()

	var tickerResp TickerResponse
	err = json.NewDecoder(resp.Body).Decode(&tickerResp)
	if err != nil {
		return TickerResponse{}, err
	}
	return tickerResp, nil
}
