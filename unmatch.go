package piscine

// Unmatch returns the element that does not have a pair.
// If all elements have pairs, it returns -1.
func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		// If count is odd, this element has no pair
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}
