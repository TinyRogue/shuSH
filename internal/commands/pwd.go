package commands

import (
	"fmt"
	"os"
)

func (command *Command) handlePrintWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not print current working directory.\nCause: %v", err)
		return
	}
	fmt.Println(dir)
}
