package main

import (
	"fmt"
	"os"
)

var validCommands = map[string]struct{}{
	"tokenize": {},
}

type Token struct {
	Type    string
	Lexeme  string
	Literal interface{}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command, filename := os.Args[1], os.Args[2]
	if _, ok := validCommands[command]; !ok {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "tokenize":
		tokens, err := scanTokens(contents)
		for _, t := range tokens {
			fmt.Printf("%s %s %v\n", t.Type, t.Lexeme, t.Literal)
		}
		fmt.Println("EOF  null")
		if err != nil {
			os.Exit(65)
		}
	}
}

func scanTokens(input []byte) ([]Token, error) {
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

	var tokens []Token
	var hasError bool

	for _, symbol := range input {
		if tokenType, ok := tokenMap[symbol]; ok {
			tokens = append(tokens, Token{Type: tokenType, Lexeme: string(symbol), Literal: nil})
		} else {
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", symbol)
			hasError = true
		}
	}

	if hasError {
		return tokens, fmt.Errorf("one or more scanning errors occurred")
	}

	return tokens, nil
}
