package commands

import (
	"fmt"
	"os"
	"strings"
)

func (command *Command) handleEcho() {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(command.subcommands, " "))
}
