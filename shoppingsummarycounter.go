package piscine

// ShoppingSummaryCounter takes a string of items separated by spaces,
// counts how many times each item appears, and returns a map with the summary.
// It ignores empty words and single "e".
func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int) // εδώ θα αποθηκεύσουμε τα αποτελέσματα
	wordRunes := []rune{}           // προσωρινό buffer για να χτίζουμε κάθε λέξη

	for _, r := range str {
		if r == ' ' {
			// Αν βρούμε κενό, τελειώνει μια λέξη.
			if len(wordRunes) > 0 {
				word := string(wordRunes)
				// αγνοούμε το "e"
				if word != "e" {
					summary[word]++
				}
				wordRunes = []rune{} // καθαρίζουμε για την επόμενη λέξη
			}
		} else {
			// Αν δεν είναι κενό, προσθέτουμε τον χαρακτήρα στη λέξη.
			wordRunes = append(wordRunes, r)
		}
	}

	// Μετά το loop, μπορεί να έχει μείνει μια τελευταία λέξη.
	if len(wordRunes) > 0 {
		word := string(wordRunes)
		if word != "e" {
			summary[word]++
		}
	}

	return summary
}
