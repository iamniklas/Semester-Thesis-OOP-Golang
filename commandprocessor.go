package main

func runCommand(bank *Bank, input string) {
	switch command := input; command {
	case "login":
		login(bank)
		break
	case "logout":
		logout(bank)
		break
	case "register":
		register(bank)
		break
	case "withdraw":
		withdraw(bank)
		break
	case "deposit":
		deposit(bank)
		break
	case "transfer":
		transfer(bank)
		break
	case "info":
		info(bank)
		break
	case "quit":
		running = false
	case "help":
		help()
	default:
		println("Whoops, command has not been found. Try again or enter help if you need support")
	}
}
