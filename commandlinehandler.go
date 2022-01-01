package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
	"strconv"
)

func register(bank *Bank) {
	println("Enter single account details in the following order and hit enter after each one")
	println("First Name - Last Name - PIN - AccountType (as number: 0 BasicAccount - 1 StandardAccount - 2 SuperPremiumAccount")
	firstname := ""
	lastname := ""
	pin := ""
	accountType := ""
	fmt.Scanln(&firstname)
	fmt.Scanln(&lastname)
	fmt.Scanln(&pin)
	fmt.Scanln(&accountType)

	accountTypeInt, _ := strconv.Atoi(accountType)

	newAcc := accounttypes.NewAccount()
	newAcc.FirstName = firstname
	newAcc.LastName = lastname
	newAcc.Pin = pin

	result := bank.Register(*newAcc, accountTypeInt)

	if result == 0 {
		fmt.Printf("Welcome %s %s\n", bank.loggedInAs.GetAccountData().FirstName, bank.loggedInAs.GetAccountData().LastName)
		fmt.Printf("Your card identifier: %s\n", bank.loggedInAs.GetAccountData().AccountIdentifier)
		println("You can now use our services")
	} else {
		fmt.Printf("Registration failed, Error-Code: %d\n", result)
	}
}

func login(bank *Bank) {
	println("Enter your account identifier")
	accountId := ""
	fmt.Scanln(&accountId)
	println("Enter your PIN")
	pin := ""
	fmt.Scanln(&pin)
	result := bank.Login(accountId, pin)

	if result == 0 {
		fmt.Printf("Welcome %s %s\n", (bank.loggedInAs).GetAccountData().FirstName, (bank.loggedInAs).GetAccountData().LastName)
		fmt.Printf("Last time logged in: %s\n", (bank.loggedInAs).GetAccountData().LastTimeLoggedIn)
	} else {
		fmt.Printf("Login failed, Error-Code: %d\n", result)
	}
}

func logout(bank *Bank) {
	result := bank.Logout()
	if result == 0 {
		println("Logout successful")
	} else {
		fmt.Printf("Logout failed, Error-Code: %d\n", result)
	}
}

func withdraw(bank *Bank) {
	println("Enter the amount to withdraw")
	input := ""
	fmt.Scanln(&input)
	floatNum, _ := strconv.ParseFloat(input, 64)
	result := bank.Withdraw(floatNum)

	if result == 0 {
		println("Withdraw successful")
	} else {
		fmt.Printf("Withdraw failed, Error-Code: %d\n", result)
	}
}

func deposit(bank *Bank) {
	println("Enter the amount to deposit")
	input := ""
	fmt.Scanln(&input)
	floatNum, _ := strconv.ParseFloat(input, 64)
	result := bank.Deposit(floatNum)

	if result == 0 {
		println("Deposit successful")
	} else {
		fmt.Printf("Deposit failed, Error-Code: %d\n", result)
	}
}

func transfer(bank *Bank) {
	targetAccId := "Enter the target account's identifier"
	fmt.Scanln(&targetAccId)
	println("Enter the amount of money you want to transfer")
	amountToTransfer := ""
	fmt.Scanln(&amountToTransfer)
	floatNum, _ := strconv.ParseFloat(amountToTransfer, 64)

	result := bank.Transfer(targetAccId, float32(floatNum))

	if result == 0 {
		println("Transfer successful")
	} else {
		fmt.Printf("Transfer failed, Error-Code: %d\n", result)
	}
}

func info(bank *Bank) {
	println(bank)
	result := bank.GetAccountInfo()

	if result != 0 {
		fmt.Printf("Failed to get Account Info, Error-Code: %d\n", result)
	}
}

func history(bank *Bank) {
	println(bank.loggedInAs.GetAccountData().GetLastFiveTransaction())
}

func help() {
	println("Available actions: login, logout, register, withdraw, deposit, info, help, quit")
}
