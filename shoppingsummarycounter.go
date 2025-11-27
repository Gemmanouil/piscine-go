package piscine

// ShoppingSummaryCounter takes a string of items separated by spaces,
// counts how many times each item appears, and returns a map with the summary.
func ShoppingSummaryCounter(str string) map[string]int {
	// Δημιουργούμε ένα map που θα κρατάει το όνομα του προϊόντος (string)
	// και τον αριθμό εμφανίσεων του (int).
	summary := make(map[string]int)

	// Χρησιμοποιούμε ένα προσωρινό slice από runes για να χτίζουμε κάθε λέξη.
	wordRunes := []rune{}

	// Διατρέχουμε χαρακτήρα-χαρακτήρα το string.
	for _, r := range str {
		if r == ' ' {
			// Αν βρούμε κενό, σημαίνει ότι τελείωσε μια λέξη.
			// Προσθέτουμε τη λέξη στο map ΜΟΝΟ αν δεν είναι κενή.
			if len(wordRunes) > 0 {
				word := string(wordRunes)
				summary[word]++
				// Καθαρίζουμε το προσωρινό slice για την επόμενη λέξη.
				wordRunes = []rune{}
			}
		} else {
			// Αν δεν είναι κενό, προσθέτουμε τον χαρακτήρα στη λέξη.
			wordRunes = append(wordRunes, r)
		}
	}

	// Μετά το loop, μπορεί να έχει μείνει μια τελευταία λέξη
	// (αν το string δεν τελείωνε με κενό).
	if len(wordRunes) > 0 {
		word := string(wordRunes)
		summary[word]++
	}

	// Επιστρέφουμε το map με τα αποτελέσματα.
	return summary
}
