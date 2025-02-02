package commands

import (
	"fmt"
)

func (command *Command) handleType() {
	if len(command.subcommands) < 1 {
		fmt.Printf(": not found\n")
		return
	}

	fn := command.subcommands[0]
	if _, ok := BuiltIns[fn]; ok {
		fmt.Printf("%s is a shell builtin\n", fn)
		return
	}
	if path, ok := hasProgram(fn); ok {
		fmt.Printf("%s is %s\n", fn, *path)
		return
	}
	fmt.Printf("%s: not found\n", fn)
}
