package quickcash

import (
	"errors"
)

type CreditCardAccount struct {
	limit      float64
	identifier string
}

var NotEnoughFundsError = errors.New("not enough funds in the account")

func (cca *CreditCardAccount) WithDraw(amount float64) error {
	if cca.CanWithDraw(amount) {
		cca.limit -= amount
	} else {
		return NotEnoughFundsError
	}
	return nil
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return cca.limit >= amount
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return cca.identifier
}
