package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"math"
	"testing"
)

func TestRegisterAccounts(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewBasicAccount()
	basicAccount.FirstName = "Basic Account"
	standardAccount := accounttypes.NewStandardAccount()
	standardAccount.FirstName = "Standard Account"
	superPremiumAccount := accounttypes.NewSuperPremiumAccount()
	superPremiumAccount.FirstName = "Super Premium Account"

	bank.Register(basicAccount.Account, 0)
	bank.Register(standardAccount.Account, 1)
	bank.Register(superPremiumAccount.Account, 2)

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
	basicAccount.Pin = "0000"
	basicAccount.AccountBalance = math.Floor(basicAccount.AccountBalance*100) / 100

	(bank.loggedInAs) = basicAccount

	bank.GetAccountInfo()

	expectedOutput := "Account Owner: Niklas Englmeier \nCard Identifier: DE6971150000 \nAccount Balance: 123.330000 \nPIN: 0000"

	if expectedOutput != (bank.loggedInAs).GetAccountInfo() {
		t.Errorf("Output invalid: Expected: %s, actual %s", expectedOutput, (bank.loggedInAs).GetAccountData())
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

	bank.registeredAccounts = append(bank.registeredAccounts, basicAccount)

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

	bank.Register(basicAccount.Account, 0)

	if len(bank.registeredAccounts) != 1 {
		t.Errorf("Excepted 1, actual %d", len(bank.registeredAccounts))
	}

	bank.GetAccountInfo()
}

func TestWithdrawAndDeposit(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 20.0

	bank.Register(*basicAccount, 2)

	bank.Deposit(50.00)
	println("-----")
	println(bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 70.00 {
		t.Errorf("Deposit Error! Got %f, excpected %f", bank.loggedInAs.GetAccountData().AccountBalance, 50.00)
	}

	bank.Withdraw(25.00)
	println("-----")
	println(bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 45.00 {
		t.Errorf("Deposit Error! Got %f, excpected %f", bank.loggedInAs.GetAccountData().AccountBalance, 25.00)
	}
}

func TestTransfer(t *testing.T) {
	bank := NewBank()

	basicAccount := accounttypes.NewSuperPremiumAccount()
	basicAccount.FirstName = "Niklas"
	basicAccount.LastName = "Englmeier"
	basicAccount.AccountIdentifier = "DE6971150000"
	basicAccount.Pin = "1234"
	basicAccount.AccountBalance = 1.33

	bank.Register(basicAccount.Account, 0)
}
