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

func (standardAccount *StandardAccount) Withdraw(amount float64, transferToOtherAccount bool) int {
	if standardAccount.AccountBalance < 0.0 && !transferToOtherAccount {
		amount = amount * 0.9
	}
	if standardAccount.AccountBalance < 0.0 && transferToOtherAccount {
		amount = amount * 1.1
	}
	if amount > 0.0 {
		standardAccount.AccountBalance = standardAccount.AccountBalance - amount
		standardAccount.Account.AddTransaction(transferToOtherAccount, -amount)
		return 0
	} else {
		return 1
	}
}

func (standardAccount *StandardAccount) Deposit(amount float64, transferToOtherAccount bool) int {
	if amount > 0.0 {
		standardAccount.AccountBalance = standardAccount.AccountBalance + amount
		standardAccount.Account.AddTransaction(transferToOtherAccount, amount)
		return 0
	} else {
		return 1
	}
}
