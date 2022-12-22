package quickcash

type SavingsAccount struct {
	balance    float64
	identifier string
}

func (sa *SavingsAccount) WithDraw(amount float64) error {
	if sa.CanWithDraw(amount) {
		sa.balance -= amount
	} else {
		return NotEnoughFundsError
	}
	return nil
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return sa.balance >= amount
}

func (sa *SavingsAccount) GetIdentifier() string {
	return sa.identifier
}
