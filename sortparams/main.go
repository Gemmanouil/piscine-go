package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get all arguments except the program name
	args := os.Args[1:]

	// Bubble sort implementation to sort arguments in ASCII order
	n := len(args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if args[j] > args[j+1] {
				// Swap if out of order
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print each argument character by character
	for _, arg := range args {
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
