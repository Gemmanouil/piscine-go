package piscine

// building the function
func MakeRange(min, max int) []int {
	// checks for any wrong inputs
	if min >= max {
		return nil
	}
	// declaring list
	varsize := max - min
	slice := make([]int, varsize)
	// declaring counter also while counter is less than max and more than min the loop goes on
	for count := min; count < varsize; count++ {
		slice[count] = min + count
	}
	// retutns the list
	return slice
}
