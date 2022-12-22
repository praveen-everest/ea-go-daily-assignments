package account

import (
	"fmt"
	"math"
)

// TDD Bank Account app

type Account struct {
	balance float64
}

type NotEnoughFundsError struct {
	amount float64
	needed float64
}

func (nefe NotEnoughFundsError) Error() string {
	return fmt.Sprintf("not enough funds in the account. requested for %f but need %f", nefe.amount, nefe.needed)
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	if acc.balance >= amount {
		acc.balance -= amount
	} else {
		return NotEnoughFundsError{amount, math.Abs(acc.balance - amount)}
	}
	return nil
}
