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

	if command == "tokenize" {
		ScanTokens(fileContents)
	}
}

func ScanTokens(input []byte) {
	for _, symbol := range input {
		if token, exists := tokenMap[symbol]; exists {
			fmt.Printf("%s %c null\n", token, symbol)
		}
	}
	fmt.Println("EOF  null")
}

var tokenMap = map[byte]string{
	'(': "LEFT_PAREN",
	')': "RIGHT_PAREN",
	'{': "LEFT_BRACE",
	'}': "RIGHT_BRACE",
	',': "COMMA",
	'.': "DOT",
	'-': "MINUS",
	'+': "PLUS",
	';': "SEMICOLON",
	'*': "STAR",
}
