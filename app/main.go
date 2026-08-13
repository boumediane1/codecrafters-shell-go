package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
				path := os.Getenv("PATH")
				paths := strings.Split(path, ":")

				found := false
				for _, path := range paths {
					err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
						if err != nil {
							return nil
						}
						if !info.IsDir() {
							substrings := strings.Split(path, "/")
							executable := substrings[len(substrings)-1]

							if arguments[1] == executable {
								if info.Mode()&0100 != 0 {
									found = true
									fmt.Printf("%s is %s\n", executable, path)
									return filepath.SkipAll
								}
							}
						}

						return nil
					})
					if err != nil {
						panic(err)
					}

					if found {
						break
					}
				}

				if !found {
					fmt.Printf("%s: not found\n", arguments[1])
				}
			}
		default:
			fmt.Printf("%s: command not found\n", arguments[0])
		}
	}
}
