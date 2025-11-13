package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	// Get all arguments except the program name
	args := os.Args[1:]

	// Sort arguments in ASCII order
	sort.Strings(args)

	// Print each argument character by character
	for _, arg := range args {
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
