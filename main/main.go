package main

import (
	"fmt"
	batch_payment "saltpay-batchpayment-test/batch-payment"
)

func main() {
	result := batch_payment.BatchPayments("GBP:100,EUR:200,CZK:1000", "743:EUR:5.76\n932:GBP:32.10\n909:CZK:223.26\n23:CZK:890.22\n902:GBP:58.23\n89:EUR:104.25\n663:EUR:97.43\n902:EUR:20.01")
	fmt.Println(result)
}
