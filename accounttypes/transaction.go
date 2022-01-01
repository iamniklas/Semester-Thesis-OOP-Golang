package accounttypes

type Transaction struct {
	tranferType int //0: ATM without receiving/sending account ; 1: Transaction to/from different account
	amount      float64
}

func NewTransaction(transferType int, amount float64) *Transaction {
	transaction := new(Transaction)
	transaction.tranferType = transferType
	transaction.amount = amount
	return transaction
}
