package accounttypes

type SuperPremiumAccount struct {
	Account
}

func NewSuperPremiumAccount() *SuperPremiumAccount {
	return &SuperPremiumAccount{}
}

func (superPremiumAccount SuperPremiumAccount) GetAccountData() *Account {
	return &superPremiumAccount.Account
}

func (superPremiumAccount SuperPremiumAccount) GetAccountInfo() string {
	return superPremiumAccount.String()
}

func (superPremiumAccount *SuperPremiumAccount) Withdraw(amount float64, transferToOtherAccount bool) int {
	if amount <= 0.0 {
		superPremiumAccount.AccountBalance = superPremiumAccount.AccountBalance - amount
		superPremiumAccount.Account.AddTransaction(transferToOtherAccount, -amount)
		return 0
	} else {
		return 1
	}
}

func (superPremiumAccount *SuperPremiumAccount) Deposit(amount float64, transferToOtherAccount bool) int {
	if amount <= 0.0 {
		superPremiumAccount.AccountBalance = superPremiumAccount.AccountBalance + amount
		superPremiumAccount.Account.AddTransaction(transferToOtherAccount, amount)
		return 0
	} else {
		return 1
	}
}
