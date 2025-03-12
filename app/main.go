package main

import (
	"fmt"
	"os"
)

var validCommands = map[string]struct{}{
	"tokenize": {},
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command, filename := os.Args[1], os.Args[2]

	if _, exists := validCommands[command]; !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		panic("Scanner not implemented")
	} else {
		fmt.Println("EOF  null")
	}
}
