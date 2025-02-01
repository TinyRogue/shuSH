package commands

import "fmt"

func handleType(args []string) {
	if _, ok := BuiltIns[args[0]]; ok {
		fmt.Printf("%s is a shell builtin\n", args[0])
	} else {
		fmt.Printf("%s: not found\n", args[0])
	}
}
