package piscine

func Capitalize(s string) string {
	// Convert the string into a slice of runes to handle Unicode characters properly
	runes := []rune(s)

	// Flag to track if we're at the beginning of a word
	newWord := true

	// Loop through each character in the string
	for i, char := range runes {
		// Check if the character is a letter or digit (alphanumeric)
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if newWord {
				// If it's the start of a word and the character is a lowercase letter, capitalize it
				if char >= 'a' && char <= 'z' {
					runes[i] = char - 32 // Convert to uppercase
				}
				newWord = false // We're now inside a word
			} else {
				// If we're inside a word and the character is uppercase, lowercase it
				if char >= 'A' && char <= 'Z' {
					runes[i] = char + 32 // Convert to lowercase
				}
			}
		} else {
			// If the character is not alphanumeric, the next character starts a new word
			newWord = true
		}
	}

	// Convert the rune slice back to a string and return it
	return string(runes)
}
