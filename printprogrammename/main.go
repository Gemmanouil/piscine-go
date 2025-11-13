package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	runes := []rune(path)
	position := 0

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '/' {
			position = i
			break
		}
	}

	for i := position + 1; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
