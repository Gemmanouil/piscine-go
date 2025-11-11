package piscine

// Fibonacci επιστρέφει τον αριθμό Fibonacci στη θέση index χρησιμοποιώντας επανάληψη (όχι αναδρομή)
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

	// Ξεκινάμε με τους δύο πρώτους αριθμούς της ακολουθίας
	prev := 0 // F(0)
	curr := 1 // F(1)

	// Επαναλαμβάνουμε από τη θέση 2 μέχρι index
	for i := 2; i <= index; i++ {
		next := prev + curr // υπολογίζουμε τον επόμενο αριθμό
		prev = curr         // μετακινούμε τον προηγούμενο
		curr = next         // μετακινούμε τον τρέχοντα
	}

	return curr // επιστρέφουμε τον αριθμό στη θέση index
}
