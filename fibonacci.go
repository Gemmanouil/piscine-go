package piscine

// RecursiveFibonacci επιστρέφει τον αριθμό Fibonacci στη θέση index
// Αν το index είναι αρνητικό, επιστρέφει -1
func RecursiveFibonacci(index int) int {
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
	return RecursiveFibonacci(index-1) + RecursiveFibonacci(index-2)
}
