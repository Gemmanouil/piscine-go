package piscine

import "github.com/01-edu/z01"

// helper to print a two-digit number
func printTwoDigit(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

// DescendComb prints all combinations of two different two-digit numbers in descending order
func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printTwoDigit(i)
			z01.PrintRune(' ')
			printTwoDigit(j)
			if !(i == 1 && j == 0) { // last pair is "01 00"
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
