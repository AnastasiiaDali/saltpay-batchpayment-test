# saltpay-batchpayment-test


SaltPay merchants in different countries are paid out everyday in different currencies. Your objective is to write a part of the payment engine that deducts some fees, and batches the payments by currency, subject to availability of funds.

The function you need to write (change according to programming language):

func BatchPayments(funds string, payments string) {
// your code here
}
Inputs
Available Funds
Delimited string, looks like this:  `GBP:100,EUR:200,CZK:1000`

Only these three currencies currently supported.

Payments to be made
Delimited string, looks like this:


`743:EUR:5.76
932:GBP:32.10
909:CZK:223.26
23:CZK:890.22
902:GBP:58.23
89:EUR:104.25
663:EUR:97.43
902:EUR:20.01`
Fields: Merchant ID (internal to SaltPay), currency and amount.

## Business Rules
1. Before paying out, deduct processing fees for each payment as a % of amount being processed: 1/3rd of a percent for GBP, 1/2 a percent for the other currencies
2. For each currency, create a batch of payments that can be paid based on the funds available. 
3. Batching should be done in a way that processes the most number of payments that can be paid using the available funds. Include the smallest payments first. You don’t have to optimise for maximum utilisation of funds.  e.g. if the EUR payments after fees are 1, 1, 2, 2, 3, 4, 5, 6, 7 and the funds available are EUR 10, then the batch of EUR payments is 1, 1, 2, 2, 3 .
4. When paying merchants, amounts should be always be rounded up  e.g. 43.545 -> 43.55, and 43.512 -> 43.52, 43.5102 -> 43.52

## Output
Batches for the payments and funds shown above (alphabetical order of currencies, decreasing order of amounts):


CZK
909:CZK:222.15
EUR
743:EUR:5.74
902:EUR:19.91
663:EUR:96.95
GBP
932:GBP:32.00
902:GBP:58.04
If there are no batches for a given currency, please do not include the currency at all.

## Guidelines
1. You will be asked to extend/modify the code you write here, so think about structure and maintainability
2. If you encounter bad input, log it out but don't stop processing the rest of the input
3. All amounts will be positive numbers
4. Feel free to use your own editor/IDE and then paste the output into HackerRank
5. You can submit the code and run the tests as many times as you like. There are some test cases you can’t see the test data for, so consider edge cases when you see these fail

