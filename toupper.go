package piscine

// ToUpper takes a string and returns a new string with all letters in uppercase
func ToUpper(s string) string {
	// We create a variable to hold the result string
	var result string

	// We loop through each character (rune) in the input string
	for _, char := range s {
		// Check if the character is a lowercase letter (between 'a' and 'z')
		if char >= 'a' && char <= 'z' {
			// Convert it to uppercase by subtracting 32 from its ASCII value
			upperChar := char - 32
			// Add the uppercase character to the result string
			result += string(upperChar)
		} else {
			// If it's not a lowercase letter, just add it as it is
			result += string(char)
		}
	}

	// Return the final result string
	return result
}
