package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/internal/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		if ok := scanner.Scan(); !ok {
			log.Fatalf("An error occurred: %v", scanner.Err())
		}
		command := scanner.Text()
		cmd, args := commands.Preprocess(command)
		commands.HandleCommand(cmd, args)
	}
}
