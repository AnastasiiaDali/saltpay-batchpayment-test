package helpers

import (
	"saltpay-batchpayment-test/models"
	"strconv"
	"strings"
)

func ParsePayments(allPayments string) []models.Payment {
	var payments []models.Payment
	for _, payment := range strings.Split(allPayments, "\n") {
		if payment == "" {
			continue
		}

		paymentPart := strings.Split(payment, ":")

		merchantID := paymentPart[0]
		currency := models.Currency(paymentPart[1])
		if !IsSupportedCurrency(currency) {
			continue
		}
		amount, err := strconv.ParseFloat(paymentPart[2], 64)
		if err != nil {
			continue
		}

		payments = append(payments, models.Payment{MerchantID: merchantID, Currency: currency, Amount: amount})
	}
	return payments
}
