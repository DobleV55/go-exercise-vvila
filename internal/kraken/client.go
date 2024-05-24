package kraken

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://api.kraken.com/0/public/Ticker?pair=%s"

type KrakenClient struct {
	BaseURL string
}

func NewKrakenClient() *KrakenClient {
	return &KrakenClient{BaseURL: BaseURL}
}

func (kc *KrakenClient) GetTicker(pair string) (KrakenTickerResponse, error) {
	url := fmt.Sprintf(kc.BaseURL, pair)
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
