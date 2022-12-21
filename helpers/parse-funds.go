package helpers

import (
	"saltpay-batchpayment-test/models"
	"strconv"
	"strings"
)

type AvailableFunds map[models.Currency]float64

func ParseAvailableFunds(funds string) AvailableFunds {
	availableFunds := make(AvailableFunds)
	for _, fund := range strings.Split(funds, ",") {
		parts := strings.Split(fund, ":")
		currency := models.Currency(parts[0])
		if !IsSupportedCurrency(currency) {
			continue
		}
		amount, _ := strconv.ParseFloat(parts[1], 64)
		availableFunds[currency] = amount
	}
	return availableFunds
}
