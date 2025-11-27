package piscine

import "strings"

// ShoppingSummaryCounter counts how many times each item appears in the string
func ShoppingSummaryCounter(str string) map[string]int {
	// Split the string into words (items)
	words := strings.Fields(str)

	// Create a map to store counts
	summary := make(map[string]int)

	// Count occurrences of each word
	for _, w := range words {
		summary[w]++
	}

	return summary
}
