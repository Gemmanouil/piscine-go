package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	wordRunes := []rune{}

	for _, r := range str {
		if r == ' ' {
			if len(wordRunes) > 0 {
				word := string(wordRunes)
				summary[word]++
				wordRunes = []rune{} // reset
			}
		} else {
			wordRunes = append(wordRunes, r)
		}
	}

	// Add last word if exists
	if len(wordRunes) > 0 {
		word := string(wordRunes)
		summary[word]++
	}

	return summary
}
