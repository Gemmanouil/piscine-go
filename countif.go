package piscine

// CountIf counts how many elements in the slice tab return true when passed to function f
func CountIf(f func(string) bool, tab []string) int {
	count := 0 // Start with zero matches

	for _, value := range tab {
		if f(value) {
			count++ // Increment count if f(value) is true
		}
	}

	return count // Return the total number of matches
}
