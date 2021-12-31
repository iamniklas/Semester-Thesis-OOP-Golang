package main

import (
	"fmt"
)

var running = true

func main() {
	bank := NewBank()

	fmt.Println("CMD BANK")
	fmt.Println("BUILT WITH GOLANG")
	fmt.Println("Type in 'help' if you need support")
	fmt.Println("---------------")

	for {
		input := ""
		fmt.Scanln(&input)

		runCommand(*bank, input)

		if !running {
			println("GOODBYE")
			break
		}
	}
}
