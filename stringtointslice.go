package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) {
	// Print opening bracket
	z01.PrintRune('[')

	for i, r := range str {
		printInt(int(r))
		// Print space between numbers
		if i != len(str)-1 {
			z01.PrintRune(' ')
		}
	}

	// Print closing bracket and newline
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

// Helper function to print an integer using only z01.PrintRune
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Handle digits
	var digits []rune
	for n > 0 {
		d := rune('0' + n%10)
		digits = append([]rune{d}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
