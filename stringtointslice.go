package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) []int {
	var result []int
	for _, r := range str {
		result = append(result, int(r))
	}

	// Print the slice manually using z01.PrintRune
	z01.PrintRune('[')
	for i, v := range result {
		printInt(v)
		if i != len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')

	return result
}

// Helper function to print an integer using z01.PrintRune
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Collect digits
	var digits []rune
	for n > 0 {
		d := rune('0' + n%10)
		digits = append([]rune{d}, digits...)
		n /= 10
	}

	// Print digits
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
