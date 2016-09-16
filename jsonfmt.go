package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printNewLine(scope int) {
	fmt.Printf("\n")
	for i := 0; i < scope; i++ {
		fmt.Printf("  ")
	}
}

func main() {
	fileName := os.Args[1]

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Couldn't read file:", err)
	}

	scope := 0
	squote := false
	dquote := false
	escaped := false
	for i := 0; i < len(dat); i++ {
		// handle escapes
		if !escaped && dat[i] == '\\' {
			escaped = true
			fmt.Printf("%c", dat[i])
			continue
		}
		if escaped {
			escaped = false
			fmt.Printf("%c", dat[i])
			continue
		}

		// handle quotes
		if !squote && dat[i] == '"' {
			dquote = !dquote
		}
		if !dquote && dat[i] == '\'' {
			squote = !squote
		}
		if squote || dquote {
			fmt.Printf("%c", dat[i])
			continue
		}

		// eliminate whitespace
		if dat[i] == '\n' || dat[i] == '\r' || dat[i] == '\t' || dat[i] == ' ' {
			continue
		}

		// closing brackets
		if dat[i] == '}' || dat[i] == ']' {
			scope--
			printNewLine(scope)
		}

		fmt.Printf("%c", dat[i])

		// opening brackets
		if dat[i] == '{' || dat[i] == '[' {
			scope++
			printNewLine(scope)
		}

		// new keys
		if dat[i] == ',' {
			printNewLine(scope)
		}

		// new values
		if dat[i] == ':' {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}
