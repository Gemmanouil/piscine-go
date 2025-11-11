package piscine

// Index επιστρέφει τη θέση του πρώτου χαρακτήρα του toFind μέσα στο s
// Αν δεν βρεθεί, επιστρέφει -1
func Index(s string, toFind string) int {
	if toFind == "" {
		return 0 // κενό string θεωρείται ότι ξεκινά από την αρχή
	}

	sRunes := []rune(s)
	toFindRunes := []rune(toFind)
	sLen := len(sRunes)
	toFindLen := len(toFindRunes)

	// Δεν γίνεται να βρεθεί αν το toFind είναι μεγαλύτερο από το s
	if toFindLen > sLen {
		return -1
	}

	// Ελέγχουμε κάθε πιθανή θέση μέσα στο s
	for i := 0; i <= sLen-toFindLen; i++ {
		match := true
		for j := 0; j < toFindLen; j++ {
			if sRunes[i+j] != toFindRunes[j] {
				match = false
				break
			}
		}
		if match {
			return i // βρέθηκε το toFind στη θέση i
		}
	}

	return -1 // δεν βρέθηκε
}
