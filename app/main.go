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
	tokenMap := map[byte]string{
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

	hasError := false

	for _, symbol := range input {
		if token, exists := tokenMap[symbol]; exists {
			fmt.Printf("%s %c null\n", token, symbol)
		} else {
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", symbol)
			hasError = true
		}
	}

	fmt.Println("EOF  null")

	// Exit with error code 65 if any errors were detected
	if hasError {
		os.Exit(65)
	}
}
