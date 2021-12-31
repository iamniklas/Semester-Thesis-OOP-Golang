package main

import (
	"Semester_Thesis_Golang/accounttypes"
)

func register(bank *Bank) {
	println("Enter single account details in the following order and hit enter after each one")
	println("First Name - Last Name - PIN")
	firstname := "Niklas"
	lastname := "Englmeier"
	pin := "1234"
	// fmt.Scanln(&firstname)
	// fmt.Scanln(&lastname)
	// fmt.Scanln(&pin)

	newAcc := accounttypes.NewAccount()
	newAcc.FirstName = firstname
	newAcc.LastName = lastname
	newAcc.Pin = pin

	bank.Register(*newAcc, 0)
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

func info(bank *Bank) {
	println(bank)
	bank.GetAccountInfo()
}

func help() {
	println("Available actions: login, logout, register, withdraw, deposit, info, help, quit")
}
