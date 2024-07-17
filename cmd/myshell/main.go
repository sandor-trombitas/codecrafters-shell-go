package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command := strings.TrimRight(input, "\n")

		if strings.HasPrefix(command, "exit") {
			var exitCode int

			_, err := fmt.Sscanf(command, "exit %d", &exitCode)
			if err == nil {
				os.Exit(exitCode)
			}

			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
