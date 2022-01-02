# Semester-Thesis-OOP-Golang

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

## 1. Description

This project is part of my semester thesis in the subject "Concepts of Programming Languages" (KP) at the Rosenheim Technical University of Applied Sciences.
The purpose of the project is to compare object-oriented programming principles in Kotlin and Golang.

The central element of the implementation is a banking system that manages different types of accounts.
There are 3 types of accounts: Basic, Standard and SuperPremium. 
Each type has a different variant of the deposit and withdraw function, whereby e.g. bonuses are awarded randomly or a fee is deducted if the account balance is lower than 0.
You can log in to the bank, register and logout. 
As mentioned, money can be withdrawn or deposited from a registered account. (Equivalent to the TransactionType ATM)
In addition, money can be transferred from one account to another registered account.
All transactions are stored in a transaction history and the last 5 can be called up in text form.

## 2. Usage

The project is written in Go. I've used the "Visual Studio Code" IDE, so the best is to use the same program for testing.

You can run the program by the command ``go run .``.

Run the unit tests by the command ``go test -v``.

## 3. Knowledge Base

Everything I know about Go comes from course content from "Concepts of Programming Languages" (KP) @ Rosenheim Technical University of Applied Sciences taught by Sebastian Macke.
