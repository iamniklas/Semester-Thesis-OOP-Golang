package accounttypes

type StandardAccount struct {
	Account
}

func NewStandardAccount() *StandardAccount {
	return &StandardAccount{}
}

func (standardAccount StandardAccount) GetAccountData() *Account {
	return &standardAccount.Account
}

func (standardAccount StandardAccount) GetAccountInfo() string {
	return standardAccount.Account.String()
}

func (standardAccount *StandardAccount) Withdraw(amount float64, transferToOtherAccount bool) {
	standardAccount.AccountBalance = standardAccount.AccountBalance - amount
	standardAccount.Account.AddTransaction(transferToOtherAccount, -amount)
}

func (standardAccount *StandardAccount) Deposit(amount float64, transferToOtherAccount bool) {
	standardAccount.AccountBalance = standardAccount.AccountBalance + amount
	standardAccount.Account.AddTransaction(transferToOtherAccount, amount)
}
