package commands

import (
	"fmt"
	"os"
)

func (command *Command) handleChangeDirectory() {
	path := ""
	if len(command.subcommands) > 0 {
		path = command.subcommands[0]
	}
	if err := os.Chdir(path); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}
