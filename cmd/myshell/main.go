package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		log.Fatalf("An error occurred: %v", scanner.Err())
	}
	command := scanner.Text()
	fmt.Printf("%s: command not found\n", command)
}
