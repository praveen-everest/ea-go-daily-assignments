package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromSavingsAccount(t *testing.T) {
	sa := &SavingsAccount{600, "Savings Account #123"}
	cca := &CreditCardAccount{1000, "Credit Card Account #ABC"}

	qc := QuickCash{
		sa,
		cca,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, "Savings Account #123", accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {
	sa := &SavingsAccount{1000, "Savings Account"}
	cca := &CreditCardAccount{2000, "Credit Card Account"}

	qc := QuickCash{
		sa,
		cca,
	}

	amt, accType := qc.getCash(1500)
	assert.Equal(t, float64(1500), amt)
	assert.Equal(t, "Credit Card Account", accType)
}
