package piscine

// IsNumeric επιστρέφει true αν η συμβολοσειρά περιέχει μόνο αριθμούς (0–9)
func IsNumeric(s string) bool {
	if s == "" {
		return false // κενή συμβολοσειρά δεν θεωρείται αριθμητική
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false // αν βρούμε μη αριθμητικό χαρακτήρα, επιστρέφουμε false
		}
	}

	return true // όλοι οι χαρακτήρες ήταν αριθμοί
}
