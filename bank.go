package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
)

var leadingSequence = "DE69"
var blz = 71150000

type Bank struct {
	loggedInAs         accounttypes.IAccount
	registeredAccounts []accounttypes.Account
}

func NewBank() *Bank {
	bank := new(Bank)
	return bank
}

func (b *Bank) Register(newAccount accounttypes.Account) {
	newAccount.AccountIdentifier = fmt.Sprintf("%s%d%s", leadingSequence, blz, "d")

	b.registeredAccounts = append(b.registeredAccounts, newAccount)
}

func (b Bank) Login(cardId string, pin string) {
	for _, v := range b.registeredAccounts {
		if v.AccountIdentifier == cardId && v.Pin == pin {
			println("MATCH")
			return
		}
	}
	println("FAIL")
}

func (b Bank) Logout() {
	if !b.requireLogin() {
		return
	}
	b.loggedInAs = nil
}

func (b Bank) Withdraw(amount float32) {

}

func (b Bank) Transfer(accId string, moneyToTransfer float32) {

}

func (b Bank) GetAccountInfo() {
	if !b.requireLogin() {
		return
	}
	println((b.loggedInAs).String())
}

func (b Bank) requireLogin() bool {
	if b.loggedInAs == nil {
		println("You have to be logged in to perform this action")
		println("Login to your account and try again")
		return false
	}
	return true
}
