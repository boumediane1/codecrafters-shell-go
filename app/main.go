package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
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
		case "type":
			if slices.Contains([]string{"exit", "echo", "type"}, arguments[1]) {
				fmt.Printf("%s is a shell builtin\n", arguments[1])
			} else {
				path, err := exec.LookPath(arguments[1])

				if err != nil {
					if errors.Is(err, exec.ErrNotFound) {
						fmt.Printf("%s: not found\n", arguments[1])
					}
					break
				}

				fmt.Printf("%s is %s\n", arguments[1], path)

			}
		default:
			_, err := exec.LookPath(arguments[0])
			if err != nil {
				if errors.Is(err, exec.ErrNotFound) {
					fmt.Printf("%s: command not found\n", arguments[0])
				}
				break
			}

			cmd := exec.Command(arguments[0], arguments[1:]...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				break
			}

			fmt.Print(string(out))
		}
	}
}
