package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			// Only add non-empty words
			if word != "" {
				summary[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	// Add the last word if exists
	if word != "" {
		summary[word]++
	}

	return summary
}
