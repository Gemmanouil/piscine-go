package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Special case: if n is 0, just print 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Step 1: Extract digits
	var digits []int
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n = n / 10
	}

	// Step 2: Sort digits manually using Bubble Sort
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				// Swap
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	// Step 3: Print digits
	for _, digit := range digits {
		z01.PrintRune(rune(digit))
	}
}
