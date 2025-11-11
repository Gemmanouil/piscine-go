package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 12 {
		return 0 // elegxos an to n einai mikrotero tou 0 h megalytero tou 12
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}

	return result
}
