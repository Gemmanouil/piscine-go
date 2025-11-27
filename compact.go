package piscine

// Compact removes zero-value elements from the slice and returns the count of non-zero elements
func Compact(ptr *[]string) int {
	// Dereference the pointer to get the slice
	slice := *ptr

	// Create a new slice to store non-empty strings
	newSlice := []string{}

	// Collect only non-empty strings
	for _, v := range slice {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}

	// Update the original slice via the pointer
	*ptr = newSlice

	// Return the number of non-empty elements
	return len(newSlice)
}
