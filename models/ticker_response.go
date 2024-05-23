package models

type TickerResponse struct {
	USD struct {
		Last float64 `json:"last"`
	} `json:"USD"`
	CHF struct {
		Last float64 `json:"last"`
	} `json:"CHF"`
	EUR struct {
		Last float64 `json:"last"`
	} `json:"EUR"`
}
