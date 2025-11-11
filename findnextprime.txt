package piscine

// FindNextPrime επιστρέφει τον πρώτο πρώτο αριθμό που είναι >= nb
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2 // το 2 είναι ο πρώτος και μικρότερος πρώτος αριθμός
	}

	for {
		if IsPrime(nb) {
			return nb // βρήκαμε τον επόμενο πρώτο αριθμό
		}
		nb++
	}
}

// IsPrime ελέγχει αν ένας αριθμός είναι πρώτος
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	i := 3
	for i*i <= nb {
		if nb%i == 0 {
			return false
		}
		i += 2
	}
	return true
}
