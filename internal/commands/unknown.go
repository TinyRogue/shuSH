package commands

import (
	"fmt"
)

func (command *Command) handleUnknown() {
	fmt.Printf("%s: command not found\n", command.primary)
}
