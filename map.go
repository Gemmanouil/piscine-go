package piscine

// Map applies the function f to each element of the slice a
// and returns a slice of bools with the results.
func Map(f func(int) bool, a []int) []bool {
	// Create a slice to store the results
	result := make([]bool, len(a))

	// Loop through each element in the input slice
	for i, value := range a {
		// Apply the function f to the value and store the result
		result[i] = f(value)
	}

	// Return the slice of results
	return result
}
