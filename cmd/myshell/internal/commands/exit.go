package commands

import "os"

func handleExit(args []string) {
	if len(args) > 0 && args[0] == "0" {
		os.Exit(0)
	}
}
