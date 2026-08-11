package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		input.Scan()
		command := input.Text()
		fmt.Printf("%s: command not found\n", command)
	}
}
