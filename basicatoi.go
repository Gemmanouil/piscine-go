package piscine

func BasicAtoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		result *= 10 + int(char-'0')

	}
	return result
}
