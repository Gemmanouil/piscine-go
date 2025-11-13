package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get all arguments except the program name
	args := os.Args[1:]

	// Loop backwards: start from the last argument down to the first
	for i := len(args) - 1; i >= 0; i-- {
		// Convert the string argument into runes (characters)
		for _, r := range []rune(args[i]) {
			z01.PrintRune(r) // Print each character
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
