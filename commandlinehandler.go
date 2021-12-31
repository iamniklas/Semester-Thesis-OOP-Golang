package main

import (
	"Semester_Thesis_Golang/accounttypes"
	"fmt"
)

func register(bank Bank) {
	println("Enter single account details in the following order and hit enter after each one")
	println("First Name - Last Name - PIN")
	firstname := ""
	lastname := ""
	pin := ""
	fmt.Scanln(&firstname)
	fmt.Scanln(&lastname)
	fmt.Scanln(&pin)

	newAcc := accounttypes.NewStandardAccount()
	newAcc.FirstName = firstname
	newAcc.LastName = lastname
	newAcc.Pin = pin

	bank.Register(newAcc.Account)
}

func login(bank Bank) {

}

func logout(bank Bank) {

}

func withdraw(bank Bank) {

}

func deposit(bank Bank) {

}

func transfer(bank Bank) {

}

func info(bank Bank) {

}

func help() {
	println("Available actions: login, logout, register, withdraw, deposit, info, help, quit")
}
