package kraken

type KrakenTickerResponse struct {
	Result map[string]struct {
		LastTradeClosed []string `json:"c"`
	} `json:"result"`
}
