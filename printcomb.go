package piscine01

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; i++ {
			for k := j + 1; k <= '9'; i++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				if i != '7' || j != '8' || k != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
