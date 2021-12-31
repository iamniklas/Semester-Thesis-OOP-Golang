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

func (b *Bank) Register(newAccount accounttypes.Account, accountType int) {
	newAccount.AccountIdentifier = fmt.Sprintf("%s%d%d", leadingSequence, blz, time.Now().Unix()/1000)
	newAccount.LastTimeLoggedIn = time.Now().Truncate(time.Second)

	switch accType := accountType; accType {
	case 0:
		basicAcc := *accounttypes.NewBasicAccount()
		basicAcc.Account = newAccount
		(b.loggedInAs) = basicAcc
		b.registeredAccounts = append(b.registeredAccounts, basicAcc)
		println(&b)
		break

	case 1:
		standardAccount := *accounttypes.NewBasicAccount()
		standardAccount.Account = newAccount
		(b.loggedInAs) = standardAccount
		b.registeredAccounts = append(b.registeredAccounts, standardAccount)
		break

	case 2:
		superPremiumAccount := accounttypes.NewBasicAccount()
		superPremiumAccount.Account = newAccount
		(b.loggedInAs) = superPremiumAccount
		b.registeredAccounts = append(b.registeredAccounts, superPremiumAccount)
		break
	}

	fmt.Printf("Welcome %s %s\n", newAccount.FirstName, newAccount.LastName)
	fmt.Printf("Your card identifier: %s\n", newAccount.AccountIdentifier)
	println("You can now use our services")
}

func (b *Bank) Login(cardId string, pin string) {
	for _, v := range b.registeredAccounts {
		if v.GetAccountData().AccountIdentifier == cardId && v.GetAccountData().Pin == pin {
			println("MATCH")
			(b.loggedInAs) = v
			//b.loggedInAs = v
			fmt.Printf("Welcome %s %s\n", (b.loggedInAs).GetAccountData().FirstName, (b.loggedInAs).GetAccountData().LastName)
			fmt.Printf("Welcome %s\n", (b.loggedInAs).GetAccountData().LastTimeLoggedIn)
			(b.loggedInAs).GetAccountData().LastTimeLoggedIn = time.Now().Truncate(time.Second)

			return
		}
	}
	println("FAIL")
}

func (b *Bank) Logout() {
	if !b.requireLogin() {
		return
	}
	b.loggedInAs = nil
}

func (b *Bank) Withdraw(amount float32) {
	if !b.requireLogin() {
		return
	}
	(b.loggedInAs).Withdraw(0.0)
}

func (b *Bank) Deposit(amount float32) {
	if !b.requireLogin() {
		return
	}
	(b.loggedInAs).Deposit(0.0)
}

func (b *Bank) Transfer(accId string, moneyToTransfer float32) {
	if !b.requireLogin() {
		return
	}
}

func (b *Bank) GetAccountInfo() {
	if !b.requireLogin() {
		return
	}
	println((b.loggedInAs).GetAccountInfo())
}

func (b *Bank) requireLogin() bool {
	println(len(b.registeredAccounts))
	if b.loggedInAs == nil {
		println("You have to be logged in to perform this action")
		println("Login to your account and try again")
		return false
	}
	return true
}
