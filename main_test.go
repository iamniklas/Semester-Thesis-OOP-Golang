package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"testing"
)

func TestRegisterAccounts(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewBasicAccount()
	standardAccount := accounttypes.NewStandardAccount()
	superPremiumAccount := accounttypes.NewSuperPremiumAccount()

	basicAccount.PerformTransaction()
	bank.Register(basicAccount.Account)

	standardAccount.PerformTransaction()
	bank.Register(standardAccount.Account)

	superPremiumAccount.PerformTransaction()
	bank.Register(superPremiumAccount.Account)

	t.Log(basicAccount.FirstName)
	t.Log(standardAccount.FirstName)
	t.Log(superPremiumAccount.FirstName)
}

func TestAccountInfo(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.AccountBalance = 123.33

	(bank.loggedInAs) = basicAccount

	(bank.loggedInAs).PerformTransaction()
	(bank.loggedInAs).PerformTransaction()

	b := (bank.loggedInAs).(*accounttypes.SuperPremiumAccount).Account
	println(b.FirstName)

	bank.GetAccountInfo()
}

func TestSlice(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.AccountBalance = 1.33

	(bank.loggedInAs) = basicAccount

	bank.registeredAccounts = append(bank.registeredAccounts, basicAccount.Account)

	println(len(bank.registeredAccounts))
	println(bank.registeredAccounts[0].FirstName)
}

func TestLogin(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 1.33

	(bank.loggedInAs) = basicAccount

	bank.registeredAccounts = append(bank.registeredAccounts, basicAccount.Account)

	bank.Login("DE6971150000", "1234")
	bank.Login("DE6971150000", "1235")
}

func TestRegister(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 1.33

	//(bank.loggedInAs) = basicAccount

	bank.Register(basicAccount.Account)

	if len(bank.registeredAccounts) != 1 {
		t.Errorf("Excepted 1, actual %d", len(bank.registeredAccounts))
	}
}
