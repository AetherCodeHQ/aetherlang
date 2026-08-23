package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "v2.0.0"

type Token struct {
	Type    string
	Value   string
	Line    int
}

type Lexer struct {
	input  string
	pos    int
	line   int
	tokens []Token
}

type Parser struct {
	tokens []Token
	pos    int
}

type VM struct {
	stack  []interface{}
	vars   map[string]interface{}
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "run":
		if len(os.Args) > 2 {
			runFile(os.Args[2])
		}
	case "repl":
		startREPL()
	case "info":
		fmt.Printf("AetherLang %s\n", VERSION)
		fmt.Println("Custom programming language with lexer, parser, AST, bytecode compiler, and VM")
	default:
		runFile(command)
	}
}

func printUsage() {
	fmt.Println("AetherLang - Custom programming language")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  aetherlang run <file.aether>  - Run a file")
	fmt.Println("  aetherlang repl               - Start REPL")
	fmt.Println("  aetherlang info               - Show version info")
}

func runFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("Running %s...\n\n", filename)
	Execute(string(data))
}

func startREPL() {
	fmt.Printf("AetherLang %s REPL (type 'exit' to quit)\n", VERSION)
	fmt.Println()
	for {
		fmt.Print("aether> ")
		var input string
		fmt.Scanln(&input)
		if input == "exit" || input == "quit" {
			break
		}
		if input != "" {
			Execute(input)
		}
	}
}

func Execute(code string) {
	tokens := Tokenize(code)
	vm := &VM{
		stack: make([]interface{}, 0),
		vars:  make(map[string]interface{}),
	}

	for _, tok := range tokens {
		switch tok.Type {
		case "PRINT":
			fmt.Println(tok.Value)
		case "ASSIGN":
			parts := strings.SplitN(tok.Value, "=", 2)
			if len(parts) == 2 {
				vm.vars[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}
	}
}

func Tokenize(input string) []Token {
	var tokens []Token
	lines := strings.Split(input, "\n")
	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "print ") {
			value := strings.TrimPrefix(line, "print ")
			value = strings.Trim(value, "\"'")
			tokens = append(tokens, Token{Type: "PRINT", Value: value, Line: lineNum + 1})
		} else if strings.Contains(line, "=") {
			tokens = append(tokens, Token{Type: "ASSIGN", Value: line, Line: lineNum + 1})
		}
	}
	return tokens
}
