package piscine

func DivMod(a int, b int, div *int, mod *int, sum *int) {
	*sum = (a / b)
	*div = a / b
	*mod = a % b
}
