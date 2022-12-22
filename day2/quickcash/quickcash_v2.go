package quickcash

type QuickCashV2 struct {
	Accounts []Withdrawable
}

func (qc *QuickCashV2) LinkAccount(account Withdrawable) {
	qc.Accounts = append(qc.Accounts, account)
}

func (qc *QuickCashV2) GetCash(amount float64) (float64, string) {
	for _, account := range qc.Accounts {
		err := account.WithDraw(amount)
		if err == nil {
			return amount, account.GetIdentifier()
		}
	}
	return 0, ""
}
