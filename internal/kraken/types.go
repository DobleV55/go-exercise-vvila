package kraken

type TickerResponse struct {
	Result map[string]struct {
		LastTradeClosed []string `json:"c"`
	} `json:"result"`
}
