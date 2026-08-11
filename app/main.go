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

		if prompt == "exit" {
			os.Exit(0)
		}

		arguments := strings.Split(prompt, " ")

		if arguments[0] == "echo" {
			fmt.Println(strings.Join(arguments[1:], " "))
		} else {
			fmt.Printf("%s: command not found\n", prompt)
		}
	}
}
