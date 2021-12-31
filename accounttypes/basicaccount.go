package accounttypes

type BasicAccount struct {
	Account
}

func NewBasicAccount() *BasicAccount {
	account := new(BasicAccount)
	return account
}

func (basicAccount *BasicAccount) PerformTransaction() {
	basicAccount.FirstName = "B"
}

func (basicAccount BasicAccount) String() string {
	return ""
}
