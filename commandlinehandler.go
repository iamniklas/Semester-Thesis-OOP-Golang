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

	bank.Register(*newAcc, accountTypeInt)
}

func login(bank *Bank) {
	println("Enter your account identifier")
	accountId := ""
	fmt.Scanln(&accountId)
	println("Enter your PIN")
	pin := ""
	fmt.Scanln(&pin)
	bank.Login(accountId, pin)
}

func logout(bank *Bank) {
	bank.Logout()
}

func withdraw(bank *Bank) {
	input := ""
	fmt.Scanln(&input)
	floatNum, _ := strconv.ParseFloat(input, 64)
	bank.Withdraw(floatNum)
}

func deposit(bank *Bank) {
	input := ""
	fmt.Scanln(&input)
	floatNum, _ := strconv.ParseFloat(input, 64)
	bank.Deposit(floatNum)
}

func transfer(bank *Bank) {
	targetAccId := "Enter the target account's identifier"
	fmt.Scanln(&targetAccId)
	println("Enter the amount of money you want to transfer")
	amountToTransfer := ""
	fmt.Scanln(&amountToTransfer)
	floatNum, _ := strconv.ParseFloat(amountToTransfer, 64)

	bank.Transfer(targetAccId, float32(floatNum))
}

func info(bank *Bank) {
	println(bank)
	bank.GetAccountInfo()
}

func help() {
	println("Available actions: login, logout, register, withdraw, deposit, info, help, quit")
}
