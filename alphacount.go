package piscine

// AlphaCount μετράει πόσα λατινικά γράμματα υπάρχουν στη συμβολοσειρά s
func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		// Ελέγχουμε αν ο χαρακτήρας είναι μεταξύ 'A'-'Z' ή 'a'-'z'
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}
