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

func (basicAccount *BasicAccount) Withdraw(amount float64, transferToOtherAccount bool) int {
	amount = amount * 0.95
	if transferToOtherAccount && amount > basicAccount.AccountBalance {
		return 1
	}
	if amount <= 0.0 {
		return 2
	}
	basicAccount.AccountBalance = basicAccount.AccountBalance - amount
	basicAccount.Account.AddTransaction(transferToOtherAccount, -amount)
	return 0
}

func (basicAccount *BasicAccount) Deposit(amount float64, transferToOtherAccount bool) int {
	amount = amount * 0.95
	if amount > 0.0 {
		basicAccount.AccountBalance = basicAccount.AccountBalance + amount
		basicAccount.Account.AddTransaction(transferToOtherAccount, amount)
		return 0
	} else {
		return 1
	}
}
