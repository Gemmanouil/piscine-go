package piscine

import "math"

// IsPrime επιστρέφει true αν το nb είναι πρώτος αριθμός, αλλιώς false
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false // το 1 και οι αρνητικοί αριθμοί δεν είναι πρώτοι
	}
	if nb == 2 {
		return true // το 2 είναι ο πρώτος και μικρότερος άρτιος πρώτος αριθμός
	}
	if nb%2 == 0 {
		return false // όλοι οι άρτιοι αριθμοί εκτός του 2 δεν είναι πρώτοι
	}

	// Ελέγχουμε διαιρέτες από 3 μέχρι sqrt(nb), μόνο περιττούς
	for i := 3; i <= int(math.Sqrt(float64(nb))); i += 2 {
		if nb%i == 0 {
			return false // βρήκαμε διαιρέτη, άρα δεν είναι πρώτος
		}
	}

	return true // δεν βρέθηκε κανένας διαιρέτης, άρα είναι πρώτος
}
