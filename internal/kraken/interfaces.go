package kraken

type KrakenClientInterface interface {
	GetTicker(pair string) (KrakenTickerResponse, error)
}
