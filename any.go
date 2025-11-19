package piscine

// Any checks if at least one element in the slice returns true when passed to function f
func Any(f func(string) bool, a []string) bool {
	for _, value := range a {
		if f(value) {
			return true // If any element returns true, return true immediately
		}
	}
	return false // If none return true, return false
}
