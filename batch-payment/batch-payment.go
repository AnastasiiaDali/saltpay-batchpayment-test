package batch_payment

import (
	"fmt"
	"sort"
	"strings"

	"saltpay-batchpayment-test/helpers"
	"saltpay-batchpayment-test/models"
)

func BatchPayments(funds string, payments string) string {
	availableFunds := helpers.ParseAvailableFunds(funds)

	paymentsToMake := helpers.ParsePayments(payments)

	helpers.DeductFees(paymentsToMake)

	sort.Slice(paymentsToMake, func(i, j int) bool {
		if paymentsToMake[i].Currency != paymentsToMake[j].Currency {
			return paymentsToMake[i].Currency < paymentsToMake[j].Currency
		}
		return paymentsToMake[i].Amount < paymentsToMake[j].Amount
	})

	batches := make(map[string][]models.Payment)
	for _, payment := range paymentsToMake {
		currency := payment.Currency
		amount := payment.Amount
		if availableFunds[currency] < amount {
			continue
		}
		availableFunds[currency] -= amount
		batches[string(currency)] = append(batches[string(currency)], payment)
	}

	var output string
	for currency, payments := range batches {
		output += currency + "\n"
		for _, payment := range payments {
			output += fmt.Sprintf("%s:%s:%.2f\n", payment.MerchantID, payment.Currency, payment.Amount)
		}
	}
	return strings.TrimSuffix(output, "\n")
}
