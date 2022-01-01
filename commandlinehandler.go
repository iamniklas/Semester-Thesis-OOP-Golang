package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
	"strconv"
)

func register(bank *Bank) {
	println("Enter single account details in the following order and hit enter after each one")
	println("First Name - Last Name - PIN")
	firstname := ""
	lastname := ""
	pin := ""
	fmt.Scanln(&firstname)
	fmt.Scanln(&lastname)
	fmt.Scanln(&pin)

	newAcc := accounttypes.NewAccount()
	newAcc.FirstName = firstname
	newAcc.LastName = lastname
	newAcc.Pin = pin

	bank.Register(*newAcc, 0)
}

func login(bank *Bank) {

}

func logout(bank *Bank) {

}

func withdraw(bank *Bank) {

}

func deposit(bank *Bank) {
	input := ""
	fmt.Scanln(&input)
	floatNum, _ := strconv.ParseFloat(input, 64)
	bank.Deposit(floatNum)
}

func transfer(bank *Bank) {

}

func info(bank *Bank) {
	println(bank)
	bank.GetAccountInfo()
}

func help() {
	println("Available actions: login, logout, register, withdraw, deposit, info, help, quit")
}
