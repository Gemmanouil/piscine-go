package piscine

func IsPrintable(s string) bool {
	check := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= ' ' || char <= '~' { // elegxei apo acii  table ta printable chars alliws vgazei false
			check += i
			return true
		}

	}
	return false
}
