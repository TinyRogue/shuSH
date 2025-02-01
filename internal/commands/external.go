package commands

import (
	"fmt"
	"os/exec"
)

func (command *Command) handleExternal() {
	program := command.primary
	if _, ok := hasProgram(program); !ok {
		command.handleUnknown()
		return
	}

	c := exec.Command(program, command.subcommands...)
	output, err := c.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(output))
}
