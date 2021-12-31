package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"math"
	"testing"
)

func TestRegisterAccounts(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewBasicAccount()
	standardAccount := accounttypes.NewStandardAccount()
	superPremiumAccount := accounttypes.NewSuperPremiumAccount()

	bank.Register(basicAccount.Account)

	bank.Register(standardAccount.Account)

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
	basicAccount.AccountBalance = math.Floor(basicAccount.AccountBalance*100) / 100

	(bank.loggedInAs) = basicAccount

	bank.GetAccountInfo()

	expectedOutput := "Account Owner: Niklas Englmeier \nCard Identifier: DE6971150000 \nAccount Balance: 123.330000"

	if expectedOutput != bank.loggedInAs.GetAccountInfo() {
		t.Errorf("Output invalid: Expected: %s, actual %s", expectedOutput, bank.loggedInAs.GetAccountData())
	}
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

func TestWithdrawAndDeposit(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 1.33

	bank.Register(basicAccount.Account)
}

func TestTransfer(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 1.33

	bank.Register(basicAccount.Account)
}
