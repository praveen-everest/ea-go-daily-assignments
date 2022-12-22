package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	_ = acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestUnsuccessfulWithdraw(t *testing.T) {
	acc := Account{100}

	err := acc.Withdraw(200)

	assert.Equal(t, NotEnoughFundsError{200, 100}, err)
}
