package piscine

import "github.com/01-edu/z01"

var sum int

func PrintComb() {
	sum := 0
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {

			z01.PrintRune(i)
			z01.PrintRune(j)
			if j == '9' {
				if i <= j {
					sum += 1
					i++

				}
			}

			if i != '9' || j != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
