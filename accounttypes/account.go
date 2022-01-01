package accounttypes

import (
	"container/list"
	"fmt"
	"time"
)

type IAccount interface {
	GetAccountData() *Account
	GetAccountInfo() string
	Withdraw(amount float64)
	Deposit(amount float64)
}

type Account struct {
	FirstName             string
	LastName              string
	AccountIdentifier     string
	Pin                   string
	LastTimeLoggedIn      time.Time
	AccountBalance        float64
	TransactionHistory    list.List
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
