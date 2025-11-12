package piscine

func ToLower(s string) string {
	// We create a variable to hold the result string
	result := " "
	// We loop through each character (rune) in the input string
	for _, char := range s {
		// Check if the character is a uppercase letter (between 'A' and 'Z')
		if char >= 'A' && char <= 'Z' {
			// Convert it to uppercase by adding 32 from its ASCII value
			lowerChar := char + 32
			// Add the uppercase character to the result string
			result += string(lowerChar)
		} else {
			// If it's not a lowercase letter, just add it as it is
			result += string(char)
		}
	}
	return result
}
