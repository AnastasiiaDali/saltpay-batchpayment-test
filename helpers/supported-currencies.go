package helpers

import "saltpay-batchpayment-test/models"

func IsSupportedCurrency(currency models.Currency) bool {
	for _, supportedCurrency := range models.SupportedCurrencies {
		if currency == supportedCurrency {
			return true
		}
	}
	return false
}
