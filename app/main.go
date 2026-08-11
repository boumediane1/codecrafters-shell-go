package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		input.Scan()
		prompt := input.Text()

		arguments := strings.Split(prompt, " ")

		switch arguments[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(arguments[1:], " "))
		default:
			fmt.Printf("%s: command not found\n", prompt)
		}
	}
}
