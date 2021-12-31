package accounttypes

type BasicAccount struct {
	Account
}

func NewBasicAccount() *BasicAccount {
	return &BasicAccount{}
}

func (basicAccount BasicAccount) GetAccountData() *Account {
	return &basicAccount.Account
}

func (basicAccount BasicAccount) GetAccountInfo() string {
	return basicAccount.Account.String()
}

func (basicAccount BasicAccount) Withdraw(amount float64) {
	panic("not implemented") // TODO: Implement
}

func (basicAccount BasicAccount) Deposit(amount float64) {
	panic("not implemented") // TODO: Implement
}
