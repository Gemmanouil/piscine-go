package piscine

// Fibonacci επιστρέφει τον αριθμό Fibonacci στη θέση index χρησιμοποιώντας αναδρομή
// Αν το index είναι αρνητικό, επιστρέφει -1
func Fibonacci(index int) int {
	if index < 0 {
		return -1 // αρνητική θέση δεν υπάρχει στην ακολουθία
	}
	if index == 0 {
		return 0 // η θέση 0 είναι πάντα 0
	}
	if index == 1 {
		return 1 // η θέση 1 είναι πάντα 1
	}

	// Για κάθε άλλη θέση, υπολογίζουμε το άθροισμα των δύο προηγούμενων
	return Fibonacci(index-1) + Fibonacci(index-2)
}
