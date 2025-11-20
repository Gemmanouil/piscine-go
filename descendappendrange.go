package piscine

func DescendAppendRange(max, min int) []int {
	// declaring list
	slice := []int{}
	if max <= min {
		return slice
	}
	// declaring counter also while counter is less than max and more than min the loop goes on
	for count := max; count > min; count-- {
		slice = append(slice, count)
	}
	// retutns the list
	return slice
}
