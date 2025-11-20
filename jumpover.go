package piscine

func JumpOver(str string) string {
	// If string is empty or shorter than 3 chars → return newline
	if str == "" || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// Always append newline
	result += "\n"
	return result
}
