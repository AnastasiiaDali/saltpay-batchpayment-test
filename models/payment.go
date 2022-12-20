package models

type Currency string

const (
	GBP Currency = "GBP"
	CZK Currency = "CZK"
	EUR Currency = "EUR"
)

var SupportedCurrencies = []Currency{GBP, CZK, EUR}

type Payment struct {
	MerchantID string
	Currency   Currency
	Amount     float64
}
