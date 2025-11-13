package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Παίρνουμε όλα τα arguments εκτός από το όνομα του προγράμματος
	args := os.Args[1:]

	// Για κάθε argument...
	for _, arg := range args {
		// Μετατρέπουμε το string σε slice από runes (για να τυπώνουμε χαρακτήρα-χαρακτήρα)
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		// Μετά από κάθε argument, τυπώνουμε ένα newline
		z01.PrintRune('\n')
	}
}
