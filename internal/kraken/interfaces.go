package kraken

type KrakenTickerClient interface {
	GetTicker(pair string) (KrakenTickerResponse, error)
}
