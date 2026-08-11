package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	fmt.Printf("%s: command not found", input.Text())
}
