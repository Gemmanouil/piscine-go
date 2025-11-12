package piscine

// ToLower takes a string and returns a new string with all letters in lowercase
func ToLower(s string) string {
	// Create a variable to hold the result string
	var result string

	// Loop through each character (rune) in the input string
	for _, char := range s {
		// Check if the character is an uppercase letter (between 'A' and 'Z')
		if char >= 'A' && char <= 'Z' {
			// Convert it to lowercase by adding 32 to its ASCII value
			lowerChar := char + 32
			// Add the lowercase character to the result string
			result += string(lowerChar)
		} else {
			// If it's not an uppercase letter, just add it as it is
			result += string(char)
		}
	}

	// Return the final result string
	return result
}
