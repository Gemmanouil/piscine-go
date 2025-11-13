package piscine

// building the function
func AppendRange(min, max int) []int {
	// checks for any wrong inputs
	if min >= max {
		return nil
	}
	// declaring list
	slice := []int{}
	// declaring counter also while counter is less than max and more than min the loop goes on
	for count := min; count < max; count++ {
		slice = append(slice, count)
	}
	// retutns the list
	return slice
}
