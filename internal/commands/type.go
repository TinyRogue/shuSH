package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var BuiltIns = map[string]struct{}{Exit: {}, Echo: {}, Type: {}}

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
	pathEnv := os.Getenv("PATH")
	dirs := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, dir := range dirs {
		path := filepath.Join(dir, command.subcommands[0])
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("%s is %s\n", command.subcommands[0], path)
			return
		}
	}
	fmt.Printf("%s: not found\n", fn)
}
