package commands

import (
	"fmt"
	"os"
)

const (
	Exit = "exit 0"
)

func HandleCommand(command string) {
	switch command {
	case Exit:
		os.Exit(0)
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
