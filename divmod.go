package piscine

func DivMod(a int, b int, div *int, mod *int) {
	a / b
	*div = a / b
	*mod = a % b
}
