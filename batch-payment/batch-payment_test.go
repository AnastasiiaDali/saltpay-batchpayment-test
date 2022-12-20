package batch_payment_test

import (
	batch_payment "saltpay-batchpayment-test/batch-payment"
	"testing"
)

func TestBatchPayments(t *testing.T) {
	t.Run("should return correct batches", func(t *testing.T) {
		testCases := []struct {
			description     string
			fund            string
			payments        string
			expectedBatches string
		}{
			{
				description:     "should create batches correctly",
				fund:            "GBP:100,EUR:200,CZK:1000",
				payments:        "743:EUR:5.76\n932:GBP:32.10\n909:CZK:223.26\n23:CZK:890.22\n902:GBP:58.23\n89:EUR:104.25\n663:EUR:97.43\n902:EUR:20.01",
				expectedBatches: "CZK\n909:CZK:222.15\nEUR\n743:EUR:5.74\n902:EUR:19.91\n663:EUR:96.95\nGBP\n932:GBP:32.00\n902:GBP:58.04",
			},
			{
				description:     "parse amount correctly with different number of decimal",
				fund:            "EUR:1000000",
				payments:        "321:EUR:123.2\n543:EUR:123.20\n909:EUR:123\n902:EUR:123.0",
				expectedBatches: "EUR\n909:EUR:122.39\n902:EUR:122.39\n321:EUR:122.59\n543:EUR:122.59",
			},
			{
				description:     "should ignore not supported currency",
				fund:            "GBP:100,EUR:1000000,CZK:1000,HRK:1000",
				payments:        "321:HRK:123.2\n543:EUR:123.20",
				expectedBatches: "EUR\n543:EUR:122.59",
			},
			{
				description:     "parse amount correctly with empty decimal",
				fund:            "EUR:1000000",
				payments:        "543:EUR:123.20\n123:EUR:",
				expectedBatches: "EUR\n543:EUR:122.59",
			},
			{
				description:     "use random letters in mid",
				fund:            "GBP:100,EUR:1000000,CZK:1000,HRK:1000",
				payments:        "wiu3t532:EUR:123.20",
				expectedBatches: "EUR\nwiu3t532:EUR:122.59",
			},
			{
				description:     "empty input for payments",
				fund:            "GBP:100,EUR:1000000,CZK:1000,HRK:1000",
				payments:        "",
				expectedBatches: "",
			},
			{
				description:     "exact amount of available funding",
				fund:            "EUR:123.20",
				payments:        "ert:EUR:123.20",
				expectedBatches: "EUR\nert:EUR:122.59",
			},
			{
				description:     "pay maximum number of payments if not enough funds",
				fund:            "EUR:10",
				payments:        "4:EUR:2\n5:EUR:3\n6:EUR:4\n7:EUR:5\n1:EUR:1\n2:EUR:1\n3:EUR:2\n8:EUR:6\n9:EUR:7",
				expectedBatches: "EUR\n1:EUR:1.00\n2:EUR:1.00\n4:EUR:1.99\n3:EUR:1.99\n5:EUR:2.99",
			},
		}

		for _, tc := range testCases {
			output := batch_payment.BatchPayments(tc.fund, tc.payments)
			if output != tc.expectedBatches {
				t.Errorf("for %s got \n%s\n, want \n%s\n", tc.description, output, tc.expectedBatches)
			}
		}
	})
}
