package piscine

// ForEach applies the function f to each element of the slice a
func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value) // Call the function f with the current value
	}
}
