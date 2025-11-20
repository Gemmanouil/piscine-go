package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) {
	z01.PrintRune('[')
	z01.PrintRune(']')
	z01.PrintRune('i')
	z01.PrintRune('n')
	z01.PrintRune('t')
	z01.PrintRune('{')

	for i, r := range str {
		printInt(int(r))
		if i != len(str)-1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('}')
	z01.PrintRune('\n')
}

//func printInt(n int) {
//if n == 0 {
//		z01.PrintRune('0')
//		return
//	}
//
//	var digits []rune
//	for n > 0 {
//		d := rune('0' + n%10)
//		digits = append([]rune{d}, digits...)
//		n /= 10
//	}
//
//	for _, d := range digits {
//		z01.PrintRune(d)
//	}
//
