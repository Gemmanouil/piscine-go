package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= "98"; i++ {
		for j := i + 1; j <= "99"; j++ {

			z01.PrintRune(rune(i/10 + '0'))
			z01.PrintRune(rune(i%10 + '0'))
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j/10 + '0'))

			if i != '9' || j != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
