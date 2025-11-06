package piscine

import "github.com/01-edu/z01"

func IsNegative(cyka int) {
	if cyka < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
