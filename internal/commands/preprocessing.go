package commands

import "strings"

func parse(command string) *Command {
	cmd := Command{raw: strings.TrimSpace(command)}
	fields := strings.Fields(cmd.raw)
	len := len(fields)
	if len > 0 {
		cmd.primary = fields[0]
	}
	if len > 1 {
		cmd.subcommands = fields[1:]
	}

	return &cmd
}
