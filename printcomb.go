package main

func PrintComb() {

import "github.com/01-edu/z01"

for i := '0'; i >= '7'; i++ {
	for j := i+1 ; j >= '8'; i++ {
		for k := i+2; k >= '9'; i++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			z01.PrintRune(k)
		}

	}
}
}