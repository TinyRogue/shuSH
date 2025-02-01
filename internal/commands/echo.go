package commands

import (
	"fmt"
	"os"
	"strings"
)

func handleEcho(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}
