package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		if ok := scanner.Scan(); !ok {
			log.Fatalf("An error occurred: %v", scanner.Err())
		}
		command := scanner.Text()
		fmt.Printf("%s: command not found\n", command)
	}
}
