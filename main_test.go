package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
	"math"
	"testing"
	"time"
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

	bank.Register(*basicAccount, 2)

	bank.Deposit(50.00)
	println("-----")
	println(bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 150.00 {
		t.Errorf("Deposit Error! Got %f, excpected %f", bank.loggedInAs.GetAccountData().AccountBalance, 150.00)
	}

	bank.Withdraw(25.00)
	println("-----")
	println(bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 125.00 {
		t.Errorf("Deposit Error! Got %f, excpected %f", bank.loggedInAs.GetAccountData().AccountBalance, 125.00)
	}
}

func TestTransfer(t *testing.T) {
	bank := NewBank()

	acc1Id := ""
	_ = acc1Id
	acc1Pin := "0847"
	_ = acc1Pin
	acc2Id := ""
	_ = acc2Id
	acc2Pin := "9870"
	_ = acc2Pin

	account1 := accounttypes.NewAccount()
	account1.FirstName = "Dominik"
	account1.LastName = "Meister"
	account1.Pin = "0847"

	account2 := accounttypes.NewAccount()
	account2.FirstName = "Lucas"
	account2.LastName = "Wirth"
	account2.Pin = "9870"

	println("ACCOUNT 1 REGISTRATION")
	bank.Register(*account1, 1)
	acc1Id = bank.loggedInAs.GetAccountData().AccountIdentifier
	println("---")
	bank.GetAccountInfo()
	bank.Logout()

	time.Sleep(2 * time.Second)

	println("---")
	println("ACCOUNT 2 REGISTRATION")
	bank.Register(*account2, 1)
	acc2Id = bank.loggedInAs.GetAccountData().AccountIdentifier
	println("---")
	bank.GetAccountInfo()

	println("---")
	println("TRANSFER")
	bank.Transfer(acc1Id, 50.0)

	println("---")
	println("Account 2 Balance")
	//Excpected: 0.0
	fmt.Printf("Sending account now has %f\n", bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 0.0 {
		t.Errorf("Transfer Send Error: Expected %f, actual %f", 0.0, bank.loggedInAs.GetAccountData().AccountBalance)
	}
	println("Logout")
	bank.Logout()
	bank.Login(acc1Id, acc1Pin)
	println("---")
	println("Account 1 Balance")
	//Excpected: 100.0
	fmt.Printf("Sending account now has %f\n", bank.loggedInAs.GetAccountData().AccountBalance)
	if bank.loggedInAs.GetAccountData().AccountBalance != 100.0 {
		t.Errorf("Transfer Receive Error: Expected %f, actual %f", 100.0, bank.loggedInAs.GetAccountData().AccountBalance)
	}
}
