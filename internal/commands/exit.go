package commands

import "os"

func (command *Command) handleExit() {
	if len(command.subcommands) > 0 && command.subcommands[0] == "0" {
		os.Exit(0)
	} else {
		command.handleUnknown()
	}
}
