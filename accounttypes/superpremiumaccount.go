package accounttypes

type SuperPremiumAccount struct {
	Account
}

func NewSuperPremiumAccount() *SuperPremiumAccount {
	account := new(SuperPremiumAccount)
	return account
}

func (account SuperPremiumAccount) GetAccountData() *Account {
	return &account.Account
}

func (account SuperPremiumAccount) GetAccountInfo() string {
	return account.String()
}

func (account *SuperPremiumAccount) Withdraw(amount float64) {
	account.AccountBalance = account.AccountBalance - amount
}

func (account *SuperPremiumAccount) Deposit(amount float64) {
	account.AccountBalance = account.AccountBalance + amount
}
