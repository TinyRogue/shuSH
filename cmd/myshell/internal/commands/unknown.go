package commands

import "fmt"

func handleUnknown(command string) {
	fmt.Printf("%s: command not found\n", command)
}
