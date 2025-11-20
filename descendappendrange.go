package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return nil
	}
	// declaring list
	slice := []int{}
	// declaring counter also while counter is less than max and more than min the loop goes on
	for count := max; count > min; count-- {
		slice = append(slice, count)
	}
	// retutns the list
	return slice
}
