package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		if ok := scanner.Scan(); !ok {
			log.Fatalf("An error occurred: %v", scanner.Err())
		}
		command := scanner.Text()
		commands.HandleCommand(command)
	}
}
