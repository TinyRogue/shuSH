package commands

import (
	"fmt"
	"os"
	"strings"
)

const (
	Exit = "exit"
	Echo = "echo"
)

func HandleCommand(command string, args []string) {
	switch command {
	case Exit:
		if len(args) > 0 && args[0] == "0" {
			os.Exit(0)
		}
	case Echo:
		fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}

func Preprocess(command string) (cmd string, args []string) {
	fields := strings.Fields(strings.TrimSpace(command))
	len := len(fields)
	if len > 0 {
		cmd = fields[0]
	}

	if len > 1 {
		args = fields[1:]
	}

	return
}
