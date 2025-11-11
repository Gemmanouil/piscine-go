package piscine

// RecursivePower υπολογίζει το nb^power χρησιμοποιώντας αναδρομή
// Αν το power είναι αρνητικό, επιστρέφει 0 σύμφωνα με τις οδηγίες
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0 // αρνητικές δυνάμεις δεν υποστηρίζονται
	}
	if power == 0 {
		return 1 // κάθε αριθμός στην δύναμη 0 είναι 1
	}

	// Αναδρομική κλήση: nb * nb^(power-1)
	return nb * RecursivePower(nb, power-1)
}
