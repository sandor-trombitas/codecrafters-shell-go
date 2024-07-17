package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	command := strings.Split(strings.TrimRight(input, "\n"), " ")[0]

	fmt.Printf("%s: command not found\n", command)
}
