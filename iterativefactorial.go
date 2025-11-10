func IterativeFactorial(n int) int {
	if n < 0 || n > 12 {
		return 0 // elegxos an to n einai mikrotero tou 0 h megalytero tou 12
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
