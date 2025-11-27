package piscine

import "sort"

// Abort returns the median of five integers
func Abort(a, b, c, d, e int) int {
	// Put all numbers in a slice
	nums := []int{a, b, c, d, e}

	// Sort the slice
	sort.Ints(nums)

	// Return the middle element (index 2)
	return nums[2]
}
