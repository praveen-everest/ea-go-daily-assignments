package quickcash

type QuickCash struct {
	PrimaryAccount   Withdrawable
	SecondaryAccount Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	err := qc.PrimaryAccount.WithDraw(amount)
	if err == nil {
		return amount, qc.PrimaryAccount.GetIdentifier()
	}
	_ = qc.SecondaryAccount.WithDraw(amount)
	return amount, qc.SecondaryAccount.GetIdentifier()
}
