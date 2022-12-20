package helpers

import (
	"math"
	"saltpay-batchpayment-test/models"
)

func DeductFees(payments []models.Payment) {
	for i := range payments {
		payment := &payments[i]
		if payment.Currency == models.GBP {
			payment.Amount = payment.Amount - payment.Amount/300
		} else {
			payment.Amount *= 0.995
		}
		payment.Amount = math.Ceil(payment.Amount*100) / 100
	}
}
