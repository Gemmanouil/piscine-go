package piscine

// IsUpper επιστρέφει true αν η συμβολοσειρά περιέχει μόνο κεφαλαία λατινικά γράμματα
func IsUpper(s string) bool {
	if s == "" {
		return false // κενή συμβολοσειρά δεν θεωρείται "μόνο κεφαλαία"
	}

	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false // αν βρούμε χαρακτήρα εκτός A-Z, επιστρέφουμε false
		}
	}

	return true // όλοι οι χαρακτήρες ήταν κεφαλαία
}
