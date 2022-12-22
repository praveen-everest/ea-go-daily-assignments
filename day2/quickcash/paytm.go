package quickcash

type PaytmWallet struct {
	balance    float64
	identifier string
}

func (pw *PaytmWallet) WithDraw(amount float64) error {
	if pw.CanWithDraw(amount) {
		pw.balance -= amount
	} else {
		return NotEnoughFundsError
	}
	return nil
}

func (pw *PaytmWallet) CanWithDraw(amount float64) bool {
	return pw.balance >= amount
}

func (pw *PaytmWallet) GetIdentifier() string {
	return pw.identifier
}
