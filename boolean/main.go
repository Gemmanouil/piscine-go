package main

import (
	"os" // Used to access command-line arguments

	"github.com/01-edu/z01" // Used to print characters one by one
)

// printStr prints a string character by character using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r) // Print each rune (character)
	}
	z01.PrintRune('\n') // Print newline at the end
}

// isEven returns true if the number is even, false otherwise
func isEven(nbr int) bool {
	return nbr%2 == 0 // Even numbers have no remainder when divided by 2
}

func main() {
	// Get the number of arguments passed to the program (excluding the program name)
	lengthOfArg := len(os.Args) - 1

	// Define the messages to print
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	// Check if the number of arguments is even and print the appropriate message
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
