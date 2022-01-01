package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
	"time"
)

var leadingSequence = "DE69"
var blz = 71150000

type Bank struct {
	loggedInAs         accounttypes.IAccount
	registeredAccounts []accounttypes.IAccount
}

func NewBank() *Bank {
	bank := new(Bank)
	return bank
}

func (b *Bank) Register(newAccount accounttypes.Account, accountType int) int {
	newAccount.AccountIdentifier = fmt.Sprintf("%s%d%d", leadingSequence, blz, time.Now().Unix())
	newAccount.LastTimeLoggedIn = time.Now().Truncate(time.Second)

	switch accType := accountType; accType {
	case 0:
		basicAcc := *accounttypes.NewBasicAccount()
		basicAcc.Account = newAccount
		basicAcc.Account.AccountBalance = 10.0
		(b.loggedInAs) = &basicAcc
		b.registeredAccounts = append(b.registeredAccounts, &basicAcc)
		break

	case 1:
		standardAccount := *accounttypes.NewStandardAccount()
		standardAccount.Account = newAccount
		standardAccount.Account.AccountBalance = 50.0
		(b.loggedInAs) = &standardAccount
		b.registeredAccounts = append(b.registeredAccounts, &standardAccount)
		break

	case 2:
		superPremiumAccount := *accounttypes.NewSuperPremiumAccount()
		superPremiumAccount.Account = newAccount
		superPremiumAccount.Account.AccountBalance = 100.0
		(b.loggedInAs) = &superPremiumAccount
		b.registeredAccounts = append(b.registeredAccounts, &superPremiumAccount)
		break
	}

	fmt.Printf("Welcome %s %s\n", newAccount.FirstName, newAccount.LastName)
	fmt.Printf("Your card identifier: %s\n", newAccount.AccountIdentifier)
	println("You can now use our services")
	return 0
}

func (b *Bank) Login(cardId string, pin string) int {
	for _, v := range b.registeredAccounts {
		if v.GetAccountData().AccountIdentifier == cardId && v.GetAccountData().Pin == pin {
			b.loggedInAs = v
			fmt.Printf("Welcome %s %s\n", (b.loggedInAs).GetAccountData().FirstName, (b.loggedInAs).GetAccountData().LastName)
			fmt.Printf("Last time logged in: %s\n", (b.loggedInAs).GetAccountData().LastTimeLoggedIn)
			(b.loggedInAs).GetAccountData().LastTimeLoggedIn = time.Now().Truncate(time.Second)
			return 0
		}
	}
	return 1
}

func (b *Bank) Logout() int {
	if !b.requireLogin() {
		return 1
	}
	b.loggedInAs = nil
	return 0
}

func (b *Bank) Withdraw(amount float64) int {
	if !b.requireLogin() {
		return 1
	}
	return (b.loggedInAs).Withdraw(amount, false)
}

func (b *Bank) Deposit(amount float64) int {
	if !b.requireLogin() {
		return 1
	}
	return (b.loggedInAs).Deposit(amount, false)
}

func (b *Bank) Transfer(accId string, moneyToTransfer float32) int {
	if !b.requireLogin() {
		return 1
	}

	//TODO: Check if account exists
	b.loggedInAs.Withdraw(float64(moneyToTransfer), true)
	for _, v := range b.registeredAccounts {
		if v.GetAccountData().AccountIdentifier == accId {
			v.Deposit(float64(moneyToTransfer), true)
			return 0
		}
	}
	return 2
}

func (b *Bank) GetAccountInfo() int {
	if !b.requireLogin() {
		return 1
	}
	println((b.loggedInAs).GetAccountInfo())
	return 0
}

func (b *Bank) requireLogin() bool {
	if b.loggedInAs == nil {
		println("You have to be logged in to perform this action")
		println("Login to your account and try again")
		return false
	}
	return true
}
