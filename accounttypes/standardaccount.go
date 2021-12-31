package accounttypes

type StandardAccount struct {
	Account
}

func NewStandardAccount() *StandardAccount {
	account := new(StandardAccount)
	return account
}

func (acc *StandardAccount) PerformTransaction() {
	acc.FirstName = "C"
}
