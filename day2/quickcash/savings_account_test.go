package quickcash

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestGetIdentifier(t *testing.T) {
	sa := SavingsAccount{100, "Savings Account #123"}

	actual := sa.GetIdentifier()
	assert.Equal(t, "Savings Account #123", actual)
}

func TestSavingsAccount_CanWithDraw(t *testing.T) {
	sa := SavingsAccount{100, "Savings Account #234"}

	for _, tc := range []struct {
		amount float64
		want   bool
	}{
		{90, true},
		{110, false},
		{100, true},
	} {
		actual := sa.CanWithDraw(tc.amount)
		assert.Equal(t, tc.want, actual)
	}
}

func TestSavingsAccount_WithDraw(t *testing.T) {
	sa := SavingsAccount{1000, "Savings Account #999"}

	for _, tc := range []struct {
		amount    float64
		ack       error
		remaining float64
	}{
		{100, nil, 900},
		{200, nil, 700},
		{600, nil, 100},
		{250, NotEnoughFundsError, 100},
	} {
		err := sa.WithDraw(tc.amount)
		assert.Equal(t, tc.ack, err)
		actual := sa.balance
		assert.Equal(t, tc.remaining, actual)
	}
}
