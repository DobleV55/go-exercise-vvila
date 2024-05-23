package kraken

type TickerClient interface {
	GetTicker(pair string) (TickerResponse, error)
}
