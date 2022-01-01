package accounttypes

import (
	"fmt"
	"time"
)

type IAccount interface {
	GetAccountData() *Account
	GetAccountInfo() string
	Withdraw(amount float64, transferToOtherAccount bool) int
	Deposit(amount float64, transferToOtherAccount bool) int
}

type Account struct {
	FirstName             string
	LastName              string
	AccountIdentifier     string
	Pin                   string
	LastTimeLoggedIn      time.Time
	AccountBalance        float64
	TransactionHistory    []Transaction
	TransactionMultiplier float64
}

func NewAccount() *Account {
	account := new(Account)
	return account
}

func (acc Account) String() string {
	return fmt.Sprintf("Account Owner: %s %s \nCard Identifier: %s \nAccount Balance: %f \nPIN: %s",
		acc.FirstName, acc.LastName, acc.AccountIdentifier, acc.AccountBalance, acc.Pin)
}

func (acc *Account) AddTransaction(transferToOtherAccount bool, amount float64) {
	transferType := 0
	if transferToOtherAccount {
		transferType = 1
	}
	acc.TransactionHistory = append(acc.TransactionHistory, *NewTransaction(transferType, amount))
}

func (acc *Account) GetLastFiveTransaction() string {
	output := ""

	reversedTransactionSlice := acc.TransactionHistory
	for i, j := 0, len(reversedTransactionSlice)-1; i < j; i, j = i+1, j-1 {
		reversedTransactionSlice[i], reversedTransactionSlice[j] = reversedTransactionSlice[j], reversedTransactionSlice[i]
	}

	for i := 0; i < min(len(reversedTransactionSlice), 5); i = i + 1 {
		output += fmt.Sprintf("%d: %f / %s\n", i+1, reversedTransactionSlice[i].amount, accountTypeIntToName(reversedTransactionSlice[i].tranferType))
	}

	return output
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func accountTypeIntToName(accType int) string {
	if accType == 0 {
		return "ATM"
	} else if accType == 1 {
		return "Account Transfer"
	}
	return ""
}
