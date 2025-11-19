package piscine

// IsSorted checks if the slice a is sorted using the comparison function f
func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true // A slice with 0 or 1 element is always sorted
	}

	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])

		if comp > 0 {
			ascending = false // Not ascending
		}
		if comp < 0 {
			descending = false // Not descending
		}
	}

	return ascending || descending
}
