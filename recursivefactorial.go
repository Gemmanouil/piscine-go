package piscine

// RecursiveFactorial υπολογίζει το παραγοντικό ενός αριθμού nb
// Αν το nb είναι αρνητικό ή αν υπάρξει overflow, επιστρέφει 0
func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0 // αρνητικοί αριθμοί δεν έχουν παραγοντικό
	}
	if nb == 0 {
		return 1 // το 0! είναι 1
	}

	result := nb * RecursiveFactorial(nb-1) // καλούμε ξανά τη συνάρτηση με nb-1

	if result < 0 {
		return 0 // αν γίνει αρνητικό σημαίνει overflow
	}

	return result
}
