package quickcash

import (
	"fmt"
	"math"
)

type CreditCardAccount struct {
	limit      float64
	identifier string
}

type NotEnoughFundsError struct {
	requested  float64
	lacking    float64
	identifier string
}

func (nefe *NotEnoughFundsError) Error() string {
	return fmt.Sprintf("not enough funds in %s. requested %f but needed %f more", nefe.identifier, nefe.requested, nefe.lacking)
}

func (cca *CreditCardAccount) WithDraw(amount float64) error {
	if cca.CanWithDraw(amount) {
		cca.limit -= amount
	} else {
		return &NotEnoughFundsError{amount, math.Abs(cca.limit - amount), cca.GetIdentifier()}
	}
	return nil
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return cca.limit >= amount
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return cca.identifier
}
