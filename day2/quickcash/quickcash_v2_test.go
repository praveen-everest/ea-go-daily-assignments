package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkNewAccount(t *testing.T) {
	sa := SavingsAccount{}
	cca := CreditCardAccount{}

	qc := QuickCashV2{[]Withdrawable{&sa, &cca}}

	paytm := PaytmWallet{}
	qc.LinkAccount(&paytm)

	assert.Equal(t, &paytm, qc.Accounts[2])
}

func TestWithDrawFromValidAccount(t *testing.T) {
	sa := SavingsAccount{100, "Savings"}
	cca := CreditCardAccount{1000, "Credit Card"}
	pw := PaytmWallet{2000, "Paytm"}

	qc := QuickCashV2{[]Withdrawable{&sa, &cca, &pw}}

	for _, tc := range []struct {
		amount float64
		acc    string
	}{
		{800, "Credit Card"},
		{50, "Savings"},
		{1400, "Paytm"},
		{500, "Paytm"},
	} {
		amt, acc := qc.GetCash(tc.amount)

		assert.Equal(t, tc.amount, amt)
		assert.Equal(t, tc.acc, acc)
	}
}
