package commands

import "strings"

func preprocess(command string) (cmd string, args []string) {
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
