package piscine

// Rot14 applies a ROT14 cipher to the input string
func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		// Check for lowercase letters
		if r >= 'a' && r <= 'z' {
			// Shift by 14 and wrap around using modulo
			r = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			// Shift uppercase letters
			r = 'A' + (r-'A'+14)%26
		}
		// Append the transformed rune to result
		result = append(result, r)
	}

	return string(result)
}
