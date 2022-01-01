package main

import (
	"fmt"
)

var running = true

func main() {
	bank := NewBank()
	println(bank)

	fmt.Println("CMD BANK")
	fmt.Println("BUILT WITH GOLANG")
	fmt.Println("Type in 'help' if you need support")
	fmt.Println("---------------")

	// testCommands := []string{"register", "info", "quit"}
	// iteration := 0

	for {
		input := ""
		fmt.Scanln(&input)

		runCommand(bank, input)
		//runCommand(bank, testCommands[iteration])
		// iteration = iteration + 1

		if !running {
			println("GOODBYE")
			break
		}
	}
}
