package helpers

import (
	"saltpay-batchpayment-test/models"
	"strconv"
	"strings"
)

type AvailableFunds map[models.Currency]float64

func ParseAvailableFunds(s string) AvailableFunds {
	availableFunds := make(AvailableFunds)
	for _, item := range strings.Split(s, ",") {
		parts := strings.Split(item, ":")
		currency := models.Currency(parts[0])
		if !IsSupportedCurrency(currency) {
			continue
		}
		amount, _ := strconv.ParseFloat(parts[1], 64)
		availableFunds[currency] = amount
	}
	return availableFunds
}
