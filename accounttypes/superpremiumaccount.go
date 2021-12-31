package accounttypes

type SuperPremiumAccount struct {
	Account
}

func NewSuperPremiumAccount() *SuperPremiumAccount {
	account := new(SuperPremiumAccount)
	return account
}

func (acc *SuperPremiumAccount) PerformTransaction() {
	acc.FirstName = "D"
}
