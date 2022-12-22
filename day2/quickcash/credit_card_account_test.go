package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardAccount_GetIdentifier(t *testing.T) {
	cca := CreditCardAccount{5000, "Credit Card Account #ABC"}

	actual := cca.GetIdentifier()
	assert.Equal(t, "Credit Card Account #ABC", actual)
}

func TestCreditCardAccount_CanWithDraw(t *testing.T) {
	cca := CreditCardAccount{50000, "Credit Card Account #DEF"}

	for _, tc := range []struct {
		amount float64
		want   bool
	}{
		{10000, true},
		{60000, false},
		{50000, true},
	} {
		actual := cca.CanWithDraw(tc.amount)
		assert.Equal(t, tc.want, actual)
	}
}

func TestCreditCardAccount_WithDraw(t *testing.T) {
	cca := CreditCardAccount{100000, "Credit Card Account #XYZ"}

	for _, tc := range []struct {
		amount    float64
		ack       error
		remaining float64
	}{
		{10000, nil, 90000},
		{50000, nil, 40000},
		{40000, nil, 0},
		{10000, NotEnoughFundsError, 0},
	} {
		err := cca.WithDraw(tc.amount)
		assert.Equal(t, tc.ack, err)
		actual := cca.limit
		assert.Equal(t, tc.remaining, actual)
	}
}
