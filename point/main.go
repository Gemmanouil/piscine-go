package main

import (
	"github.com/01-edu/z01"
)

// Define a struct named 'point' with two integer fields: x and y
type point struct {
	x int
	y int
}

// This function sets the values of x and y in the given point to 42 and 21
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Helper function to print an integer using z01.PrintRune
func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	points := &point{}
	setPoint(points)

	// Print "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.x)

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNbr(points.y)

	// Print newline
	z01.PrintRune('\n')
}
