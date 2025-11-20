package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) []int {
	var result []int

	for _, r := range str {
		result = append(result, int(r))
	}

	// Print the slice using z01.PrintRune
	z01.PrintRune('[')
	for i, v := range result {
		// Convert int to string digits
		s := []rune{}
		if v == 0 {
			s = append(s, '0')
		} else {
			n := v
			var digits []rune
			for n > 0 {
				digits = append([]rune{rune('0' + n%10)}, digits...)
				n /= 10
			}
			s = append(s, digits...)
		}

		// Print digits
		for _, d := range s {
			z01.PrintRune(d)
		}

		// Print space between numbers
		if i != len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')

	return result
}
