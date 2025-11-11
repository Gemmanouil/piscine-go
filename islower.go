package piscine

// IsLower επιστρέφει true αν η συμβολοσειρά περιέχει μόνο μικρά λατινικά γράμματα
func IsLower(s string) bool {
	if s == "" {
		return false // κενή συμβολοσειρά δεν θεωρείται "μόνο μικρά"
	}

	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false // αν βρούμε χαρακτήρα εκτός a-z, επιστρέφουμε false
		}
	}

	return true // όλοι οι χαρακτήρες ήταν μικρά
}
