package piscine

// Abort returns the median of five integers without using sort package
func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	// Simple bubble sort
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// Median is the middle element (index 2)
	return nums[2]
}
