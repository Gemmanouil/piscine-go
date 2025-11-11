package piscine

func IsPrintable(s string) bool {
	firstPrintableChar := ' '
	lastPrintableChar := '~'
	for _, r := range s {
		if r < firstPrintableChar || r > lastPrintableChar {
			return false
		}
	}
	return true
}
