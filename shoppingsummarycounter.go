package piscine

// ShoppingSummaryCounter counts how many times each item appears in the string.
// It includes empty words (caused by consecutive spaces) and single-character words.
func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int) // map to store word counts
	word := ""                      // temporary buffer for building each word

	for _, r := range str {
		if r == ' ' {
			// When we hit a space, store the current word (even if empty)
			summary[word]++
			word = "" // reset buffer
		} else {
			// Add character to the current word
			word += string(r)
		}
	}

	// Add the last word after the loop
	summary[word]++

	return summary
}
