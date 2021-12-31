package accounttypes

type StandardAccount struct {
	Account
}

func NewStandardAccount() *StandardAccount {
	account := new(StandardAccount)
	return account
}

func (standardAccount StandardAccount) GetAccountData() Account {
	return standardAccount.Account
}

func (standardAccount StandardAccount) GetAccountInfo() string {
	return standardAccount.Account.String()
}

func (standardAccount StandardAccount) Withdraw(amount float64) {
	panic("not implemented") // TODO: Implement
}

func (standardAccount StandardAccount) Deposit(amount float64) {
	panic("not implemented") // TODO: Implement
}
