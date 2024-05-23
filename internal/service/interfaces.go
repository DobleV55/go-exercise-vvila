package service

type PriceServiceInterface interface {
	GetLastTradedPrices(pairs []string) ([]map[string]string, error)
}
